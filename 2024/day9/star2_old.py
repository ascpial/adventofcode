input = "2333133121414131402"

with open('input.txt', 'r') as file:
    input = file.read().strip()

output = sum([int(i) for i in input])
disk = [None] * output

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

def empty_space(size, max):
    """Returns the first empty space found"""
    i = 0
    n = 0
    while i < max:
        if n == size:
            return i - n
        if disk[i] is None:
            i+=1
            n+=1
        else:
            i+=1
            n=0
    return None

print(input)
cur_x = len(disk)
for i in range(len(input) // 2, -1, -1):
    if i % 100 == 0:
        print(f"{i} / {len(input)//2}")
    file_length = int(input[i*2])
    cur_x = cur_x - file_length
    empty_i = empty_space(file_length, cur_x)
    # print(i, i*2, file_length)
    # print(''.join(map(lambda i: str(i) if i is not None else ".", disk)))
    # print(" "*cur_x+"^")
    if empty_i is not None and empty_i < cur_x:
        # print(" "*empty_i + "^")
        for m in range(file_length):
            # print(l)
            disk[cur_x+m] = None
            disk[empty_i+m] = i
    if i > 0:
        cur_x -= int(input[i*2-1])

# print(test)

sum = 0
for i, j in enumerate(disk):
    if j is not None:
        sum += i * j

print(sum)

