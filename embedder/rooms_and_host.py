import pandas as pd
import psycopg2 as pg
import sys

conn = pg.connect("postgres://go_backend:go_backend@host.docker.internal:5432/go_db?sslmode=disable")
cursor = conn.cursor()

df = pd.read_csv('data/listings.csv')

import sys

def table_has_data(table_name, cursor):
    cursor.execute(f"SELECT EXISTS (SELECT 1 FROM {table_name} LIMIT 1);")
    return cursor.fetchone()[0] 

if table_has_data("rooms", cursor) and table_has_data("hosts", cursor):
    print("Tables contain data. Skipping...")
    sys.exit(0)
else:
    print("One or both tables are empty or do not exist. Running migration...")


host_table_query = """CREATE TABLE IF NOT EXISTS hosts (
    id BIGINT PRIMARY KEY,
    host_url VARCHAR(200) NOT NULL,
    host_name VARCHAR(50) NOT NULL,
    host_since DATE,
    host_location VARCHAR(50),
    host_about TEXT,
    host_thumbnail_url VARCHAR(200),
    host_picture_url VARCHAR(200)
);
"""

cursor.execute(host_table_query)
conn.commit()


host_columns = [
    'host_id',
    'host_url',
    'host_name',
    'host_since',
    'host_location',
    'host_about',
    'host_thumbnail_url',
    'host_picture_url',
]


host_table = df[host_columns].copy()


host_table.drop_duplicates(subset="host_id", keep='first', inplace=True)


import re
def remove_html_tags(text):
    clean_text = re.sub(r'<.*?>', '', text)
    return clean_text


conn.rollback()
cursor.execute("BEGIN;")


insert_query = """INSERT INTO hosts (
    id, host_url, host_name, host_since, host_location, 
    host_about, host_thumbnail_url, host_picture_url
)
VALUES (%s, %s, %s, %s, %s, %s, %s, %s);
"""

for _, row in host_table.iterrows():
    vals = list(row.values)
    for i, val in enumerate(vals):
        if i == 5 and not pd.isna(val):
            vals[i] = remove_html_tags(val)
    cursor.execute(insert_query, vals)


conn.commit()

rooms_columns = [
    'id'    ,
    'listing_url',
    'name',
    'description',
    'neighborhood_overview',
    'picture_url',
    'price'    ,
    'bedrooms',
    'beds',
    'room_type',
    'property_type',
    'neighbourhood',
    'host_id'
]

rooms_df = df[rooms_columns].copy()
rooms_df['price'] = rooms_df['price'].replace({'\$': '', ',' : ''}, regex=True).astype(float)

rooms_table_query = """CREATE TABLE IF NOT EXISTS rooms (
    id BIGINT PRIMARY KEY,
    listing_url VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    neighborhood_overview TEXT,
    picture_url VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2),
    bedrooms INT,
    beds INT,
    room_type VARCHAR(50) NOT NULL,
    property_type VARCHAR(50) NOT NULL,
    neighbourhood VARCHAR(100)
);
"""

cursor.execute(rooms_table_query)
conn.commit()


cursor.execute("BEGIN;")


insert_query = """
    INSERT INTO rooms (
        id, listing_url, name, description, neighborhood_overview, 
        picture_url, price, bedrooms, beds, room_type, property_type, neighbourhood, host_id
    )
    VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
"""
for _, row in rooms_df.iterrows():
    vals = row.values
    try: 
        for i, val in enumerate(vals):
            if (i == 3 or i == 4) and not pd.isna(val):
                vals[i] = remove_html_tags(val)
            if (i == 7 or i == 8) and not pd.isna(val):
                vals[i] = int(val)
            elif pd.isna(val):
                vals[i] = None
        cursor.execute(insert_query, vals)
    except Exception as e:
        print(e)
        print(vals)
        break

conn.commit()