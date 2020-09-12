import re


fname = "bench.log"
writename = "filtered.log"

with open(fname) as f:
    with open(writename, "w") as ff:
        for line in f.readlines():
            m = re.search(r'user_agent":".*bot.*', line)
            if m:
                ff.write(line)

        