data = [(53,250),(91,1330),(67,1081),(68,1025)]

result = 1

for time, distance in data:
  max = 0
  ways=[]
  for c in range(1,time):
    if c*(time-c)>distance:
      ways.append(c)
  result = len(ways)*result

print(result)
