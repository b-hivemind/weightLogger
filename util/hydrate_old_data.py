import csv
import sys
from datetime import datetime

import mysql.connector

USAGE = """
        Usage: python hydrate_old_data.py /path/to/dump user_id
        """
TABLE_NAME = "main"

if len(sys.argv) != 3:
    print(USAGE)
    sys.exit(1)

filename = sys.argv[1]
user_id = sys.argv[2]

csv_rows = []

try:
    with open(filename, newline='') as csvfile:
        dump_reader = csv.reader(csvfile)
        for row in dump_reader:
            csv_rows.append(row)
except FileNotFoundError() as e:
    print(e)
    sys.exit(-1)

print(f"OK | Read {len(csv_rows)} successfully")

def sanitize_date_stamp(date_stamp):
    date_stamp = date_stamp.strip("'")
    raw_time = datetime.strptime(date_stamp, "%Y-%m-%d")
    return "%.0f" % raw_time.timestamp()

final_rows = [(sanitize_date_stamp(csvrow[0]), user_id, csvrow[1]) for csvrow in csv_rows]

print("OK | Successfully sanitized data")

try:
    db = mysql.connector.connect(
    host="172.18.0.3",
    user="archimedes",
    password="",
    database="weight_data"
    )
except mysql.connector.errors.DatabaseError as e:
    print(e)
    sys.exit(1)

print("OK | Successfully established connection to db")

cursor = db.cursor()
query = f"INSERT INTO {TABLE_NAME} (timestamp, uid, weight) VALUES (%s, %s, %s)"
cursor.executemany(query, final_rows)
db.commit()
print(cursor.rowcount)