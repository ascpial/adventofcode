import re

p = re.compile(r"(?:mul\(([0-9]{1,3}),([0-9]{1,3})\))|(?:(do)\(\))|(?:(don't)\(\))")

# input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
# input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
with open("input.txt", 'r', encoding='utf8') as file:
    input = file.read()

inputs = p.findall(input)
print(inputs)

sum = 0
enabled = True
for a, b, do, dont in inputs:
    if do != '':
        enabled = True
    elif dont != '':
        enabled = False
    elif enabled:
        sum += int(a) * int(b)

print(sum)
