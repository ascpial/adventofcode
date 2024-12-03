import re

p = re.compile(r"mul\(([0-9]{1,3}),([0-9]{1,3})\)")

input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
with open("input.txt", 'r', encoding='utf8') as file:
    input = file.read()

inputs = p.findall(input)

sum = 0
for a, b in inputs:
    sum += int(a) * int(b)

print(sum)
