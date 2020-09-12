import re
import json
from urllib.parse import urlparse

fname = "bench2-after-key.log"
late = "latency.dic.log"

start = '{"time":'
counter = 0
ll = []


with open(fname) as f:
    
    for line in f.readlines():
        counter += 1
        if counter < 10: continue
        idx = line.index(start)
        json_line = line[idx:]
        jso = json.loads(json_line)
        ll.append(jso)
    
    ll2 = sorted(ll, key = lambda x: -x["latency"] if "latency" in x else 0)
    with open("latency.csv", "w") as ff2:
        for ob in ll2:
            new_ob = {}
            if not "latency" in ob: continue
            uridecoded = urlparse(ob["uri"])
            urinew = uridecoded.geturl()
            ff2.write(ob["method"] + "," + urinew + "," + ob["user_agent"] + "," + str(ob["latency"]) + "\n")



        