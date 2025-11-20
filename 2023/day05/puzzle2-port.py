with open('/home/ascpial/Projets/AdventOfCode/python/day5/input.txt', 'r') as file:
    text = file.read()

raw_categories = text.strip().split("\n\n")

raw_seeds = [int(i) for i in raw_categories[0].split(":")[1].split(" ") if i != ""]

seeds = [(raw_seeds[n], raw_seeds[n+1]) for n in range(0, len(raw_seeds), 2)]

categories = []

for category in raw_categories:
    category_data = []
    for line in category.split("\n")[1:]:
        if line == "":
            continue
        values = line.split(" ")
        destination = int(values[0])
        source = int(values[1])
        length = int(values[2])
        category_data.append((destination, length, source))
    category_data.sort(key=lambda interval: interval[2])
    categories.append(category_data)

categories.reverse()

def value_in(value: int, source: list[tuple[int, int]]) -> bool:
    for start, length in source:
        if start <= value < start + length:
            return True
    return False

def reverse(value: int, categories: list[list[tuple[int, int, int]]]) -> int:
    state = value
    for category in categories:
        for destination, length, source in category:
            if destination <= state < destination + length:
                state = source + state - destination
                break
    return state

print(reverse(1, categories))
print(value_in(reverse(3218294198, categories), seeds))

i = 1
while not value_in(reverse(i, categories), seeds):
    i += 1
print(i)
