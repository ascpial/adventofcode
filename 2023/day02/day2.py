with open("input.txt", 'r', encoding='utf-8') as file:
    text = file.readlines()

# Parse the input in the format:
# (id, ((red, green, blue), ...))

COLORS = ("red", "green", "blue")

partys = []

for line in text:
    line = line.strip()
    game, sets_list = line.split(": ")
    game_id = int(game[5:])
    raw_sets = sets_list.split("; ")
    sets = []
    for set in raw_sets:
        raw_colors = set.split(", ")
        colors = [0] * len(COLORS)

        for balls in raw_colors:
            for i, color in enumerate(COLORS):
                if balls.endswith(color):
                    colors[i] += int(balls.split(" ")[0])
        
        sets.append(colors)

    partys.append([game_id, sets])

# First part
if False:
    valids = []
    for id, sets in partys:
        valid = True
        for set in sets:
            if set[0] <= 12 and set[1] <= 13 and set[2] <= 14:
                continue
            else:
                valid = False
                break

        if valid:
            valids.append(id)
        

    print(sum(valids))

# Second part
if True:
    powers = []
    for _, sets in partys:
        colors = [0, 0, 0]
        for set in sets:
            for color, number in enumerate(set):
                if number > colors[color]:
                    colors[color] = number
        power = colors[0] * colors[1] * colors[2]
        powers.append(power)
    print(sum(powers))
