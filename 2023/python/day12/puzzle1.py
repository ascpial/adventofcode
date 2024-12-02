from functools import lru_cache

with open("input_check.txt", 'r', encoding='utf-8') as file:
    data = file.readlines()

PART_2 = True

# Format: ("string", states)
records = []
for line in data:
    string, groups = line.strip().split(" ")
    groups = tuple([int(i) for i in groups.split(",")])
    records.append((string, groups))

if PART_2:
    for i, (record, groups) in enumerate(records):
        new_record = "?".join([record]*5)
        new_groups = groups * 5
        records[i] = (new_record, new_groups)

def verify(combination: str, excepted_length, groups) -> bool:
    """Check if the given record start matchs groups"""
    current = 0
    current_group = 0
    current_group_index = None
    while current < len(combination):
        # print(combination[current], end=" ")
        if current_group is not None and current_group_index is not None:
            if current_group_index < groups[current_group] and combination[current] != "#":
                return False
            elif current_group_index == groups[current_group] and combination[current] == "#":
                return False
            else:
                current_group_index += 1
                if current_group_index > groups[current_group]:
                    # print(f"ending group {current_group}")
                    current_group += 1
                    if current_group >= len(groups):
                        current_group = None
                        current_group_index = None
                    else:
                        current_group_index = None
        elif current_group is not None:
            if combination[current] == "#":
                # print(f"started group {current_group}")
                current_group_index = 1
        elif current_group is None:
            if combination[current] != ".":
                return False
        current += 1
        # print("")
    # print(combination)
    if len(combination) == excepted_length:
        if current_group is None:
            return True
        elif current_group != len(groups) - 1:
            return False
        elif current_group_index is None:
            return False
        elif current_group_index >= groups[current_group]:
            return True
        return False
    return True

@lru_cache
def test(combination: str, source: str, groups) -> int:
    if len(combination) == len(source):
        if verify(combination, len(source), groups):
            # print(combination)
            return 1
        else:
            return 0

    # We check that this combination is possible
    if not verify(combination, len(source), groups):
        return 0
    next = source[len(combination)]
    combinations = 0
    if next != "?":
        combinations += test(combination + next, source, groups)
    else:
        combinations += test(combination + "#", source, groups)
        combinations += test(combination + ".", source, groups)
    return combinations


# print(test("", "?###????????", (3,2,1)))

sum = 0
for record, groups in records:
    sum += test("", record, groups)
print(sum)
