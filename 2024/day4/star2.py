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

# text = "XMAS"
# found = 0
# for mode in [5, 6, 7, 8]:
#     for y in range(maxy(mode)+1):
#         for x in range(maxx(mode)+1):
#             if get(x, y, mode) == "A" and get(x-1, y, mode) == "M" and get(x+1, y, mode) == "S":
#                 if (get(x, y-1, mode) == "M" and get(x, y+1, mode) == "S") or (get(x, y+1, mode) == "M" and get(x, y-1, mode) == "S"):
#                     found += 1
# print(found)

count = 0
for y in range(1, len(matrix)-1):
    for x in range(1, len(matrix[0])-1):
        if matrix[y][x] == "A":
            if (matrix[y-1][x-1]=="M" and matrix[y+1][x+1]=="S") or (matrix[y-1][x-1]=="S" and matrix[y+1][x+1]=="M"):
                if (matrix[y+1][x-1]=="M" and matrix[y-1][x+1]=="S") or (matrix[y+1][x-1]=="S" and matrix[y-1][x+1]=="M"):
                    count += 1
print(count)
