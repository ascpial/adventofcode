import math

with open('input.txt', 'r', encoding='utf-8') as file:
    text = file.read()

categories = text.split("\n\n")

raw_seeds = [int(i) for i in categories[0].split(": ")[1].split(" ")]

def intersect(a1, a2):
    if a1.start > a2.end or a2.start > a1.end:
        return None
    return slice(max(a1.start, a2.start), min(a1.end, a2.end))  

class Map:
    def __init__(self, data: list, name: str):
        self.name = name
        self.data = data
        # on va trier tout ça
        self.data.sort(key=lambda map: map[1])

    def __getitem__(self, index):
        if isinstance(index, int):
            print("hello...")
            for destination, source, length in self.data:
                if source <= index < source + length:
                    return destination - source + index
            return index
        elif isinstance(index, slice):
            print("hello!")
            start = index.start
            end = index.stop
            intervals = []
            for destination, source, length in self.data:
                if (source >= start or start <= source + length) and source <= end:
                    intervals.append(slice(start, source - 1))
                    intervals.append(slice(self[max(source, start)], self[min(source + length - 1, end)]))
                    start = source + length
            if start <= end:
                intervals.append(slice(start, end))
            return intervals
                    
            
    def __repr__(self):
        return f"<Map {self.name}>"

start = slice(0, 16)

maps = [(7, 2, 3), (2, 7, 3), (21, 12, 2), (24, 15, 6)]

test = Map(maps, "ozaijd")
for interval in test[start]:
    print(interval)



# maps = []
# for category in categories[1:]:
#     lines = category.strip().split("\n")
#     name = lines[0].strip()[:-5]
#     data = [[int(i) for i in line.split(" ")] for line in lines[1:]]
#     maps.append(Map(data, name))

# def get_location(seed: int) -> int:
#     source = seed
#     for map in maps:
#         source = map[source]
#     return source

# min_value = math.inf

# ranges = []
# for i in range(len(raw_seeds) // 2):
#     ranges.append(slice(raw_seeds[2 * i], raw_seeds[2 * i] + raw_seeds[2 * i + 1] - 1))


# print(min_value)
