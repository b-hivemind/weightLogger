import csv 
import json
import mysql.connector

from datetime import datetime

DB_NAME = 'weight_data'
TABLE_NAME = 'main'

def connect_to_db(db_name):
    """ Tries to connect to a database, returns tuple (database obj, connection) or None """
    try:
        db = mysql.connector.connect(
        host='localhost',
        user='root',
        password='Not@GoodPassword1234',
        database=f'{db_name}'
        )
    except mysql.connector.Error as e:
        print(f"Something went wrong: {e}")
        return None
    return (db, db.cursor())

with open('update_weight_data.csv', mode='r') as file:
    csvFile = csv.DictReader(file)
    db, cursor = connect_to_db(DB_NAME)
    count = 0
    for line in csvFile:
        dateObj = datetime.strptime(line['Date'], '%m/%d/%Y')
        correctedDate = dateObj.strftime("%Y-%m-%d")
        query = "INSERT INTO {} VALUES ('{}', '{}')".format(TABLE_NAME, correctedDate, line["Weight(lbs)"])
        cursor.execute(query)
        count += 1
    print(f"Successfully imported {count} rows")
    db.commit()
    
        
        