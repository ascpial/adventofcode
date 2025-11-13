input = """47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47"""

with open('input.txt', 'r') as file:
    input = file.read().strip()

is_checking_manuals = False
before = {}
after = {}
manuals = []
for line in input.split("\n"):
    if line == "":
        is_checking_manuals = True
        continue
    if not is_checking_manuals:
        b, a = line.split("|")
        b, a = int(b), int(a)
        if a in before:
            before[a].add(b)
        else:
            before[a] = {b}
        if b in after:
            after[b].add(a)
        else:
            after[b] = {a}
    else:
        manuals.append([int(i) for i in line.split(",")])

"""before/after usage
If x should appear before y, then x is in before[y]
If x should appear after x, then x is in after[y]
If x is neither in before[x] or before[y], then there is no restriction on their respective placements.
"""

correct_manuals = []

for manual in manuals:
    for i in range(len(manual)):
        for j in range(i): # the last one is i-1
            if manual[j] in after.get(manual[i], []): # page of indice j should be printed AFTER page of indice i
                # print(manual, "nope", i, j)
                break
        else: # we didn't break, we can keep going fine
            continue
        break # we broke out of the loop, so the manual is not valid
    else: # we never broke out of a loop so the manual is fine
        correct_manuals.append(manual)

# print(correct_manuals)

sum = 0
for manual in correct_manuals:
    sum += manual[len(manual)//2]
print(sum)
