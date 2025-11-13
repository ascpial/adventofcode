input="""7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"""

with open("input.txt", 'r', encoding='utf8') as file:
    input = file.read()

reports = [[int(i) for i in line.split(" ")] for line in input.strip().split("\n")]

def isreportvalid(report):
    direction = -1 if report[0] > report[1] else 1
    previous = report[0]
    for record in report[1:]:
        if abs(previous - record) < 1 or abs(previous - record) > 3:
            break
        elif (direction == -1 and previous < record) or (direction == 1 and previous > record):
            break
        previous = record
    else:
        return True
    return False

safe = 0
for report in reports:
    if isreportvalid(report):
        safe += 1
    else:
        for i in range(len(report)):
            new_report = report[:i] + report[i+1:]
            print(new_report)
            if isreportvalid(new_report):
                safe += 1
                break
   
print(safe)
