import re


fname = "bench.log"
fname = "filtered.log"
writename = "no-bottole.log"

with open(fname) as f:
    with open(writename, "w") as ff:
        for line in f.readlines():
            m = re.search(r'user_agent":".*bottle.*', line)
            if not m:
                ff.write(line)

        