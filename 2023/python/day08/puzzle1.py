with open("input.txt", 'r', encoding='utf-8') as file:
    data = file.readlines()

directions = [1 if i=="R" else 0 for i in data[0].strip()]

print(directions)

nodes = {}
for line in data[2:]:
    node = line.split(" = ")[0]
    children = line.split(" = ")[1]
    left = children[1:-2].split(", ")[0]
    right = children[1:-2].split(", ")[1]
    nodes[node] = (left, right)

START = "AAA"
TARGET = "ZZZ"

instruction_index = 0
position = START
steps = 0

while position != TARGET:
    steps += 1
    position = nodes[position][directions[instruction_index]]
    instruction_index += 1
    if instruction_index >= len(directions):
        instruction_index = 0

print(steps)
