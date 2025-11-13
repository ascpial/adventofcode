input = "2333133121414131402"

with open('input.txt', 'r') as file:
    input = file.read().strip()

output = sum([int(i) for i in input])
disk: list[None|int] = [None] * output

empty = False
n = 0
block_i = 0
for i in map(int, input):
    if not empty:
        for j in range(i):
            disk[j+n] = block_i
        block_i += 1
    n += i
    empty = not empty

cache = {}

def free(size, max):
    if size not in cache:
        cache[size] = 0
    n = 0
    while cache[size] <= max:
        if n==size:
            return cache[size] - n
        if disk[cache[size]] is None:
            n += 1
            cache[size] += 1
        else:
            n=0
            cache[size] += 1
    return None

# print(input)

x = len(disk)
for i in range(len(input) // 2, 0, -1):
    # if i%100==0:
    #     print(f"{i} / {len(input)//2}")
    size = int(input[2*i])
    # print(size)
    x -= size
    free_i = free(size, x)
    # print(''.join(map(lambda i: str(i) if i is not None else '.', disk)))
    # print(' '*x+'^')
    if free_i is not None:
        # print(' '*free_i+'^')
        for m in range(size):
            disk[x+m] = None
            disk[free_i+m] = i
        # cache = {}
    x -= int(input[2*i-1])

# print(test)

sum = 0
for i, j in enumerate(disk):
    if j is not None:
        sum += i * j

print(sum)

