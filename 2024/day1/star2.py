# input = """3   4
# 4   3
# 2   5
# 1   3
# 3   9
# 3   3"""

with open("input.txt", 'r', encoding='utf8') as file:
    input = file.read()

liste1 = []
liste2 = []
for line in input.split("\n"):
    if len(line) > 1:
        nombre1, nombre2 = line.split()
        liste1.append(int(nombre1))
        liste2.append(int(nombre2))

sim = 0
for i in range(len(liste1)):
    sim += liste1[i] * liste2.count(liste1[i])
print(sim)
