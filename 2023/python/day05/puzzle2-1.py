import math

with open('input.txt', 'r', encoding='utf-8') as file:
    text = file.read()

categories = text.split("\n\n")

raw_seeds = [int(i) for i in categories[0].split(": ")[1].split(" ")]

maps = []

def intersect(a1, a2):
    if a1.start > a2.end or a2.start > a1.end:
        return None
    return slice(max(a1.start, a2.start), min(a1.end, a2.end))  

class Map:
    def __init__(self, data: list, name: str):
        self.name = name
        self.data = data

    def __getitem__(self, index):
        if isinstance(index, int):
            for destination, source, length in self.data:
                if source <= index < source + length:
                    return destination - source + index
            return index
        if isinstance(index, slice):
            intervals = []
            for destination, source, length in self.data:
                interval = slice(source, source + length - 1)
                pass

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

ranges = []
for i in range(len(raw_seeds) // 2):
    ranges.append(slice(raw_seeds[2 * i], raw_seeds[2 * i] + raw_seeds[2 * i + 1] - 1))


print(min_value)
