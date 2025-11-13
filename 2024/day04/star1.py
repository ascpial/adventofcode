input = """MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"""

with open("input.txt", 'r', encoding='utf8') as file:
    input = file.read()

matrix = [line.strip() for line in input.split("\n")]
if matrix[-1] == "":
    del matrix[-1]

"""
Modes:
    - 1 : normal
    - 2 : backward
    - 3 : up
    - 4 : down
    - 5 : diagonal up right
"""

def maxy(mode):
    if mode in [1, 2]:
        return len(matrix[0]) - 1
    elif mode in [3, 4]:
        return len(matrix) - 1
    elif mode in [5, 6]:
        return len(matrix[0]) * 2 - 2
    else:
        return len(matrix[0]) * 2 - 2

def maxx(mode):
    if mode in [3, 4]:
        return len(matrix)-1
    else:
        return len(matrix[0])-1

def get(l, c, mode):
    if mode==1 and 0 <= l < len(matrix) and 0 <= c < len(matrix[l]):
        return matrix[l][c]
    elif mode==2 and 0 <= l < len(matrix):
        return get(l, len(matrix[l]) - c - 1, 1)
    elif mode==3 and 0 <= c < len(matrix) and 0 <= l < len(matrix[c]):
        return matrix[c][l]
    elif mode==4:
        return get(l, len(matrix)-c-1, 3)
    elif mode==5 and 0 <= l + c - len(matrix) + 1 < len(matrix) and 0 <= c < len(matrix[0]):
        return matrix[l+c-len(matrix)+1][c]
    elif mode==6 and 0 <= len(matrix[0])-1-c < len(matrix[0]):
        return get(l, len(matrix[0])-1-c, 5)
    elif mode==7 and 0 <= l - c < len(matrix[0]) and 0 <= c < len(matrix[0]):
        return matrix[l - c][c]
    elif mode==8 and 0 <= len(matrix[0])-1-c < len(matrix[0]):
        return get(l, len(matrix[0])-1-c, 7)
    else:
        return "*"

text = "XMAS"
found = 0
for mode in range(1, 9):
    for line in range(maxy(mode)+1):
        i = 0
        j = maxx(mode)-3
        while j >= 0:
            if get(line, j, mode) != text[i]:
                i, j = 0, j - i - 1
            else:
                i += 1
                j += 1
            if i == len(text):
                found += 1
                i, j = 0, j - i - len(text)

print(found)
