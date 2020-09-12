import re
import json

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
    with open(late, "w") as ff:
        for ob in ll2:
            st = json.dumps(ob) + "\n"
            ff.write(st)
    with open("latency.csv", "w") as ff2:
        for ob in ll2:
            new_ob = {}
            if not "latency" in ob: continue
            ff2.write(ob["method"] + "," + ob["uri"] + "," + ob["user_agent"] + "," + str(ob["latency"]) + "\n")



        