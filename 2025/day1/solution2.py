# I planned to do only the first solution but pipeline didn't agree.

def sign(n):
    return 1 if n >= 0 else -1

with open('input.txt', 'r') as input:
    turns = [int(i) for i in input.read().strip().replace("L", "").replace("R", "-").split("\n")]
    position = 50
    counter = 0
    for turn in turns:
        for _ in range(abs(turn)):
            position = (position + sign(turn)) % 100
            if position == 0:
                counter += 1
    print(counter)
