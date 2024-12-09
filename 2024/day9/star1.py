input = "2333133121414131402"

with open('input.txt', 'r') as file:
    input = file.read().strip()

output = sum([int(i) for i in input])
test = [None] * output

empty = False
n = 0
block_i = 0
for i in map(int, input):
    if not empty:
        for j in range(i):
            test[j+n] = block_i
        block_i += 1
    n += i
    empty = not empty

a = 0
b = len(test) - 1
while a < b:
    if test[a] is not None:
        a += 1
    elif test[b] is None:
        b -= 1
    else:
        test[a], test[b] = test[b], None
        a += 1
        b -= 1

print(test)

sum = 0
for i, j in enumerate(test):
    if j is not None:
        sum += i * j

print(sum)

