import mysql.connector

def is_ordered(answer_list, values):
    idx = []
    for v in values:
        idx.append( answer_list.index(v) )
    
    id2 = sorted(idx)
    if not idx == id2:
        print("wrong ordered!: idx: ", idx)
        return  False
    return True

cnx = mysql.connector.connect(user='isucon', password='isucon',
                              host='127.0.0.1',
                              database='isuumo')
cursor = cnx.cursor()
query = ("SELECT DISTINCT features "
         "FROM isuumo.estate")
# query = ("SELECT DISTINCT features "
#          "FROM isuumo.chair")

cursor.execute(query)


features = []
for (feature,) in cursor:
  print(feature)
  for val in feature.split(","):
    #   features.add(val)
    if val not in features:
        features.append(val)
  if not is_ordered(features, feature.split(",")):
      print("WRONG! : ", features, feature)


# with open("chair-value-text2.txt", "w") as f:
#     for l in features:
#         f.write(l + "\n")




cnx.close()
