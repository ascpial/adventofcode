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

second_part = False
links = {}
manuals = []
for line in input.split("\n"):
    if line == "":
        second_part = True
        continue
    if not second_part:
        first, second = line.split("|")
        first = int(first)
        second = int(second)
        if first in links:
            links[first].append(second)
        else:
            links[first] = [second]
    else:
        manuals.append([int(i) for i in line.split(",")])

# print(links, manuals)

cache = {}
def print_after(start):
    if start not in cache:
        visited = {}
        followers = []
        def aux(edge):
            followers.append(edge)
            if edge in links:
                for p in links[edge]:
                    if p == start:
                        return None
                    if p not in visited:
                        visited[p] = True
                        aux(p)
        visited[start] = True
        aux(start)
        cache[start] = followers
    return cache[start]

correct = []
for manual in manuals:
    print(manual)
    # for i in range(len(manual)-1, -1, -1):
    #     followers = print_after(manual[i])
    #     if followers is None:
    #         break
    #     else:
    #         for j in range(0, i):
    #             if manual[j] in followers:
    #                 print(f"break {manual} {i} {j}")
    #                 break
    #         else:
    #             continue
    #         break
    # else:
    #     print(f"found correct {manual}")
    #     correct.append(manual)
    for i in range(len(manual)):
        for j in range(i):
            if j in links.get(i, []):
                break
        else:
            continue
        break
    else:
        print(f"found correct {manual}")
        correct.append(manual)

# print(correct)

somme = 0
for manual in correct:
    somme += manual[len(manual) // 2]
print(somme)
