import psycopg2 as pg
import pandas as pd
from config import *

with pg.connect(
    database='vue-go-gcp-20191111-db',
    user='maasa',
    password=password,
    host='localhost',
    port=5432,
) as conn:
    with conn.cursor() as cur:
        # events = pd.read_sql('SELECT * FROM events', con=conn)
        # print(events)
        cur.execute(" \
        DROP TABLE IF EXISTS events; \
        CREATE TABLE events( \
        ID    TEXT    NOT NULL, \
        Title   TEXT       NOT NULL, \
        Description    TEXT, \
        PRIMARY KEY (id) \
        ); \
        ");
        # IDをNumに変えて、serialの一意なIDを後で追加しよっと。

        cur.execute(" \
        INSERT INTO events (ID, Title, Description) VALUES \
        ('1', 'Introduction to Golang', 'Come join us for a chance to learn how golang works and get to eventually try it out'), \
        ('2','222','22222'), \
        ('3','333','33333'), \
        ('4','444','44444'); \
        ");
        events = pd.read_sql('SELECT * FROM events', con=conn)
        print(events)
