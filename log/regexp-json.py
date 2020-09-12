import re
import json

fname = "bench.log"

start = '{"time":'

counter = 0
with open(fname) as f:
    for line in f.readlines():
        counter += 1
        if counter < 10: continue
        idx = line.index(start)
        json_line = line[idx:]
        jso = json.loads(json_line)
        # print(jso)

        if "status" in jso and jso["status"] != 200:
            print(jso)
        # break


        