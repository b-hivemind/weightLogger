import requests
import json
import logging as log

API_BASEPATH = 'http://127.0.0.1'
API_PORT = ':10000'

def api_fetch_weight():
    res = requests.get("http://127.0.0.1:10000/")
    if not res.status_code == 200:
        log.error(f"{res.status_code} {res.content}")
        return (res.status_code, res.content)
    else:
        return json.loads(res.content)

def api_log_weight(weight, force):
    if not weight.replace('.', '').isdecimal():
        log.error(f"Non-numeric weight detected {weight}")
        return False
    data = json.dumps({'weight': weight, 'force': force})
    res = requests.post("{}{}/new".format(API_BASEPATH, API_PORT), data=data)
    if res.status_code != 200:
        log.error(f"{res.status_code} {res.content}")
        return (res.status_code, res.content)
    return json.loads(res.content)

def api_fetch_all_weight():
    res = requests.get("http://127.0.0.1:10000/all")
    if not res.status_code == 200:
        log.error(f"{res.status_code} {res.content}")
        return (res.status_code, res.content)
    else:
        return json.loads(res.content)