input = """0 1 10 99 999"""

stones = [int(i) for i in input.strip().split()]

def blink(liste):
    output = []
    for i in liste:
        if i == 0:
            output.append(1)
        elif len(str(i)) % 2 == 0:
            output.append(int(str(i)[:len(str(i))//2]))
            output.append(int(str(i)[len(str(i))//2:]))
        else:
            output.append(i*2024)
    return output

def blink_multiple_times(liste, times):
    output = liste
    for _ in range(times):
        output = blink(output)
    return output

print(blink(stones))
stones = [125, 17]
print(blink_multiple_times(stones, 2))
print(blink_multiple_times(stones, 3))
print(len(blink_multiple_times(stones, 25)))

stones = [872027, 227, 18, 9760, 0, 4, 67716, 9245696]
print(len(blink_multiple_times(stones, 50)))
