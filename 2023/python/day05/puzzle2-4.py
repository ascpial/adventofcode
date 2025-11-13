with open('input_check.txt', 'r', encoding='utf-8') as file:
    input = file.read()

categories = input.split("\n\n")

raw_seeds = [int(i) for i in categories[0].split(": ")[1].split(" ")]
seeds = []

for i in range(0, len(raw_seeds), 2):
    seeds.append(
            (raw_seeds[i], raw_seeds[i] + raw_seeds[i+1] - 1),
            )

print(seeds)
maps = []

for category in categories[1:]:
    category = category.strip()
    lines = category.split("\n")
    name = lines[0][:-5]
    intervals = []
    for line in lines[1:]:
        target, source, length = (int(i) for i in line.split(" "))
        intervals.append((source, length, target))
    intervals.sort(key=lambda interval: interval[0])
    maps.append((name, intervals))

def deepen(maps_collection, interval):

    start, end = interval

    results = []

    maps = maps_collection[0][1]
    map_id = -1
    print(" "*(6-len(maps_collection)), maps_collection, len(maps_collection))
    while start <= end and map_id < len(maps) - 1:
        map_id += 1
        print(" "*(7-len(maps_collection)),  map_id)
        map = maps[map_id]
        map_start = map[0]
        map_end  = map[1] + map_start - 1
        map_target = map[2]

        if start < map_start:
            print("less")
            if len(maps_collection) == 1:
                results.append((start, map_start - 1))
            else:
                results.extend(deepen(maps_collection[1:], (start, map_start - 1)))
            start = map_start - 1
        if not (start > map_end or map_start > end):
            image = (
                    max(map_start, start) - map_start + map_target,
                    min(map_end, end) - map_start + map_target,
                    )
            if len(maps_collection) == 1:
                print("end")
                results.append(image)
            else:
                results.extend(deepen(maps_collection[1:], image))

            start = map_end + 1
    
    if start <= end:
        print("the very end")
        results.append((start, end))

    print("results:", results)
    return results

results = []

for start, end in seeds:
    results.extend(deepen(maps, (start, end)))

print(results)
