with open("input_check.txt", 'r', encoding='utf-8') as file:
    data = [line.strip() for line in file.readlines()]

def show(data):
    return
    for line in data:
        print(line)

show(data)

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

for column in sorted(columns_to_add, reverse=True):
    for line in range(len(data)):
        data[line] = data[line][:column] + "." + data[line][column:]

for line in sorted(lines_to_add, reverse=True):
    data.insert(line, "." * len(data[0]))

show(data)

galaxies = []
# x = 0
for y in range(len(data)):
    for x in range(len(data[0])):
        if data[y][x] == "#":
            # x += 1
            # galaxies.append(x)
            galaxies.append((x, y))
print(galaxies)

distances_sum = 0
for i, galaxy1 in enumerate(galaxies):
    for galaxy2 in galaxies[i+1:]:
        distance = max(galaxy1[0], galaxy2[0]) - min(galaxy1[0], galaxy2[0]) + max(galaxy1[1], galaxy2[1]) - min(galaxy1[1], galaxy2[1])
        distances_sum += distance

print(distances_sum)
