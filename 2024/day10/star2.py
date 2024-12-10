input = """89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"""

with open('input.txt', 'r') as file:
    input = file.read()

map = [[int(c) for c in i.strip()] for i in input.strip().split("\n")]

def get(x, y):
    if 0 <= x < len(map[0]) and 0 <= y < len(map):
        return map[y][x]

def score(x, y):
    i = get(x, y)
    if i == 9:
        return 1
    sum = 0
    for dx, dy in ((1, 0), (0, 1), (-1, 0), (0, -1)):
        if get(x+dx, y+dy) == i+1:
            sum += score(x+dx, y+dy)
    return sum

somme = 0

for y in range(len(map)):
    for x in range(len(map[0])):
        if get(x, y) == 0:
            test=score(x, y)
            # print(len(test), test)
            somme += test

print(somme)
