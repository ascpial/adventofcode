import functools

@functools.cache
def number_of_stones(entry, n, limit):
    if n == limit:
        return 1
    elif entry == 0:
        return number_of_stones(1, n+1, limit)
    if len(str(entry)) % 2 == 0:
        return number_of_stones(int(str(entry)[:len(str(entry))//2]), n+1, limit) + number_of_stones(int(str(entry)[len(str(entry))//2:]), n+1, limit)
    return number_of_stones(entry * 2024, n+1, limit)

input = [872027, 227, 18, 9760, 0, 4, 67716, 9245696]
somme = 0

for i in input:
    somme += number_of_stones(i, 0, 75)

print(somme)

import time
time.sleep(10)
