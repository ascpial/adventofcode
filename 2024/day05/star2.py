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

incorrect_manuals = []

def correct_order(a, b):
    """Returns True if a should appear before b, False if a should appear after b, and None if there is no restriction on their placement"""
    if a in after.get(b, []):
        return False
    if a in before.get(b, []):
        return True
    return None

def is_this_order_fine(manual):
    for i in range(len(manual)):
        for j in range(i): # the last one is i-1
            if correct_order(manual[j], manual[i]) is False:
            # if manual[j] in after.get(manual[i], []): # page of indice j should be printed AFTER page of indice i
                # print(manual, "nope", i, j)
                break
        else: # we didn't break, we can keep going fine
            continue
        break # we broke out of the loop, so the manual is not valid
    else: # we never broke out of a loop so the manual is fine
        return True
    return False

for manual in manuals:
    if not is_this_order_fine(manual):
        incorrect_manuals.append(manual)

print(incorrect_manuals)

# On essaie d'abord avec un basique tri à bulle ig
corrected_manuals = []
for manual in incorrect_manuals:
    processing_manual = manual.copy()
    for i in range(len(processing_manual)):
        for j in range(len(processing_manual)-1):
            c = correct_order(processing_manual[j], processing_manual[j+1])
            if c is False: # wrong order!!!
                processing_manual [j], processing_manual[j+1] = processing_manual[j+1], processing_manual[j]
    if not is_this_order_fine(processing_manual):
        raise ValueError(f"The manual {manual} has been converted to {processing_manual} but the order is not valid")
    corrected_manuals.append(processing_manual)
    # print(processing_manual)

sum = 0
for manual in corrected_manuals:
    sum += manual[len(manual)//2]
print(sum)
