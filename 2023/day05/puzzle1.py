import math

with open('input.txt', 'r', encoding='utf-8') as file:
    text = file.read()

categories = text.split("\n\n")

raw_seeds = [int(i) for i in categories[0].split(": ")[1].split(" ")]

maps = []

class Map:
    def __init__(self, data: list, name: str):
        self.name = name
        self.data = data

    def __getitem__(self, index):
        for destination, source, length in self.data:
            if source <= index < source + length:
                return destination - source + index
        return index

    def __repr__(self):
        return f"<Map {self.name}>"

for category in categories[1:]:
    lines = category.strip().split("\n")
    name = lines[0].strip()[:-5]
    data = [[int(i) for i in line.split(" ")] for line in lines[1:]]
    maps.append(Map(data, name))

def get_location(seed: int) -> int:
    source = seed
    for map in maps:
        source = map[source]
    return source

min_value = math.inf

seeds = []

for i in range(len(raw_seeds) // 2):
    seeds.append((raw_seeds[2 * i], raw_seeds[2 * i + 1]))


for start, length in seeds:
    print(start, length)
    for seed in range(start, start+length):
        location = get_location(seed)
        if location < min_value:
            min_value = location

print(min_value)
