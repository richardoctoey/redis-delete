Difference using delete by pattern vs flushdb.

Instead of put everything in one redis database, you can: define segments of your data, and put it on multiple database.
Therefore to delete the thoose data segments, you can use only "FLUSHDB" instead of "SCAN" and "DEL".

![alt text](./image.png)

On this example, it will:
1. insert 100k data.
2. DeleteFlushDB (FLUSHDB, on DB0) & DeleteByPattern (SCAN and DEL, on DB1)


This is the time difference between FLUSHDB (33 ms) and SCAN-DEL (15 seconds).
