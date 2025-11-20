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

STARTS = [node for node in nodes.keys() if node.endswith("A")]

repeat = []

for start in STARTS:
    instruction_index = 1
    node = nodes[start][directions[0]]
    steps = 1

    while instruction_index != 0 or node != start:
        steps += 1
        node = nodes[node][directions[instruction_index]]
        instruction_index += 1
        if instruction_index >= len(directions):
            instruction_index = 0
    
    repeat.append(steps)

print(repeat)
