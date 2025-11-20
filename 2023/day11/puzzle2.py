with open("input.txt", 'r', encoding='utf-8') as file:
    data = [line.strip() for line in file.readlines()]

columns_to_add = []
lines_to_add = []

for column in range(len(data[0])):
    for line in range(len(data)):
        if data[line][column] == "#":
            break
    else:
        columns_to_add.append(column)

for line in range(len(data)):
    for column in range(len(data[0])):
        if data[line][column] == "#":
            break
    else:
        lines_to_add.append(line)

print(columns_to_add)
print(lines_to_add)

galaxies = set()
STEP = 1000000
real_y = 0
for y in range(len(data)):
    real_x = 0
    for x in range(len(data[0])):
        if data[y][x] == "#":
            galaxies.add((real_x, real_y))
        if x in columns_to_add:
            real_x += STEP
        else:
            real_x += 1
    if y in lines_to_add:
        real_y += STEP
    else:
        real_y += 1
galaxies = list(galaxies)
print(galaxies)

distances_sum = 0
for i, galaxy1 in enumerate(galaxies):
    for galaxy2 in galaxies[i+1:]:
        distance = max(galaxy1[0], galaxy2[0]) - min(galaxy1[0], galaxy2[0]) + max(galaxy1[1], galaxy2[1]) - min(galaxy1[1], galaxy2[1])
        distances_sum += distance

print(distances_sum)
