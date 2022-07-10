import statistics
import copy
import os
import datetime
import logging as log

import discord
import plotly.express as px
import pandas as pd

from api import api_fetch_weight, api_log_weight, api_fetch_all_weight

# Development Notes
#
# Currently run as its own bot, or, the only function of Atlas, but eventually
# it would be better to create this as a separate library for atlas
# Current code should be written with that mindset to reduce future toil

class WeightLogger:
    __TOKEN = None
    __client = None

    def __init__(self, token):
        self.__TOKEN = token
        self.__client = discord.Client()
        self.on_ready = self.__client.event(self.on_ready)
        self.on_message = self.__client.event(self.on_message)
        self.__weight_data = None

    def start(self):
        log.basicConfig(filename="archimedes_discord.log", level=log.INFO)
        self.__client.run(self.__TOKEN, destructor=self.graceful_death)
    
    async def graceful_death(self, loop):
        log.warn('Received signal - shutting down')
        loop.stop()
        
    def _get_weight_data(self):
        log.info("Fetching weight")
        res = api_fetch_weight()
        if isinstance(res, tuple):
            log.error(res)
        return res

    def _format_weight(self, weight_data):
        log.info("Formatting weight")
        formatted_data = copy.deepcopy(weight_data)
        for coll in formatted_data:
            for entry in coll:
                try:
                    entry['weight'] = float(entry['weight'])
                except ValueError:
                    log.error(f"Could not convert {entry} to float")
        return formatted_data

    def _make_chart(self, weight_data):
        if not os.path.exists("charts"):
            log.warn("Charts dir not found, creating...")
            os.mkdir("charts")
        df = pd.DataFrame.from_dict(weight_data)
        fig = px.line(df, x='date', y='weight')
        fig.write_image("charts/chart.jpeg")
        log.info("Successfully saved image: charts/chart.jpeg")
        return discord.File("charts/chart.jpeg")
    
    def _get_stats(self):
        self.__weight_data = self._format_weight(self._get_weight_data())
        try:
            last_2_days = self.__weight_data[0]
            current_weight = last_2_days[0]['weight']
            last_7_days = self.__weight_data[1]
            last_30_days = self.__weight_data[2]
        except IndexError as e:
            log.error("Index eror {}".format(e))
            return
        stats = {}
        stats['last_logged_weight_date'] = last_2_days[0]['date']
        stats['last_logged_weight'] = current_weight
        stats['last_weight_delta'] = current_weight - last_2_days[1]['weight']
        stats['seven_day_average'] = statistics.mean([entry['weight'] for entry in last_7_days])
        stats['seven_day_av_delta'] = current_weight - stats['seven_day_average']
        stats['thirty_day_average'] = statistics.mean([entry['weight'] for entry in last_30_days])
        stats['thirty_day_av_delta'] = current_weight - stats['thirty_day_average']
        return stats
    
    async def on_ready(self):
        log.info("Live")
        
    async def on_message(self, message):
        content = message.content.lower()
        log.info(f"Received message: {content}")
        # Parse the message if it starts with !
        if content.startswith('!'):
            # Strip the ! for convenience
            content = content.lstrip('!').split()
            # Format the command into command and arguments
            command = content[0]
            args = content[1:]

            # Begin command table
            if command == "log": await self.write_weight(message.author, message.channel, args)
            elif command == "stats": await self.stats(message.author, message.channel)
            elif command == "chart": await self.chart(message.channel, args)
            elif command == "ping": await message.channel.send("Pong")

    async def chart(self, channel, args):
        period = 0
        try:
            period = int(args[0])
        except IndexError:
            log.warn("No period provided, continuing with all")
        except ValueError:
            log.error(f"Received invalid period type {period}")
            await channel.send("Please provide a numerical period")
            return
        
        if period == 0:
            weight_data = api_fetch_all_weight()
            if isinstance(weight_data, tuple):
                log.error(weight_data)
                return
            weight_data = [weight_data]
            weight_data  = self._format_weight(weight_data)
            await channel.send(file=self._make_chart(weight_data[0]))
            return

        self.__weight_data = self._format_weight(self._get_weight_data())
        if period == 7:
            await channel.send(file=self._make_chart(self.__weight_data[1]))
        elif period == 30:
            await channel.send(file=self._make_chart(self.__weight_data[2]))
        else:
            log.error(f"Received invalid period value {period}")
            await channel.send("Invalid period, currently supported periods: [0, 7, 30]")
    
    async def stats(self, user, channel):
        stats = self._get_stats()
        embed = discord.Embed(title="Weight data for {} (updated {})\n".format(user.nick if user.nick else user.name, stats['last_logged_weight_date']), color=discord.Color.blue())
        
        embed.set_author(name="Archimedes", icon_url="http://3.bp.blogspot.com/-NxzImif4-DI/UDoI8hVSp-I/AAAAAAAAAAM/ZPPNiAcqBDQ/s1600/archimedes.jpg")

        outputstr = f"{stats['last_logged_weight']:9.2f} lbs\t"
        outputstr += "🔴 +" if stats['last_weight_delta'] > 0 else "🟢 "
        outputstr += f"{stats['last_weight_delta']:9.2f}\n".lstrip()
        embed.add_field(name="Last weight", value=outputstr)
        
        outputstr = f"{stats['seven_day_average']:9.2f} lbs\t"
        outputstr += "🔴 +" if stats['seven_day_av_delta'] > 0 else "🟢 "    
        outputstr += f"{stats['seven_day_av_delta']:9.2f}\n".lstrip()
        embed.add_field(name="7 Day average", value=outputstr)

        outputstr = f"{stats['thirty_day_average']:9.2f} lbs\t"
        outputstr += "🔴 +" if stats['thirty_day_av_delta'] > 0 else "🟢 "
        outputstr += f"{stats['thirty_day_av_delta']:9.2f}\n".lstrip()
        embed.add_field(name="30 Day average", value=outputstr)
        log.info("Embed created successfully, sending stats")
        await channel.send(embed=embed)
        log.info("Calling chart")
        await self.chart(channel, ["30"])


    async def write_weight(self, user, channel, args):
        embed = discord.Embed(color=discord.Color.red())
        embed.set_author(name="Archimedes", icon_url="http://3.bp.blogspot.com/-NxzImif4-DI/UDoI8hVSp-I/AAAAAAAAAAM/ZPPNiAcqBDQ/s1600/archimedes.jpg")
        if len(args) == 0:
            log.error("No weight received")
            embed.title = "No weight specified"
            embed.description = "Usage: !log <weight>"
            await channel.send(embed=embed)
            return
        force = False
        if len(args) > 1 and args[1].lower() == "-f":
            log.info("Force flag detected")
            force = True
        res = api_log_weight(args[0], force)
        if not res:
            log.error("Incorrect weight format detected")
            embed.title = "Incorrect weight format"
            embed.description = "Please provide a numerical value for weight"
            await channel.send(embed=embed)
            return
        elif isinstance(res, tuple):
            # Error condition
            if res[0] == 300:
                log.warn(f"Repeat weight log attempt for date {datetime.date.today()}")
                embed.color = discord.Color.gold()
                embed.title = "Weight has already been logged for {}".format(datetime.date.today())
                embed.description = "Use !log <weight> -f to overwrite"
                await channel.send(embed=embed)
                return
            else:
                log.error(res)
        else:
            log.info(f"Successfully logged weight {res.get('weight')} at {datetime.date.today()}")
            embed.color = discord.Color.green()
            embed.title = "Success!"
            embed.description = ("{} lbs successfully logged for {}".format(res.get('weight'), res.get('date')))
            await channel.send(embed=embed)
            await self.stats(user, channel)

# Add your token here
token = ""
w = WeightLogger(token)
w.start()

            
