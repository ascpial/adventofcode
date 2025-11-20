with open('input.txt', 'r', encoding='utf-8') as file:
    text = file.readlines()

# Part 1
if False:
    callibrations = []

    DIGITS = [str(i) for i in range(10)]

    def isdigit(char):
        return char in DIGITS

    for line in text:
        line = line.strip()
        first = None
        for char in line:
            if isdigit(char):
                first = char
                break
        second = None
        for char in reversed(line):
            if isdigit(char):
                second = char
                break
        callibrations.append(int(first + second))

    print(sum(callibrations))

# Part 2

if True:
    callibrations = []

    DIGITS = [
            ('one', 1),
            ('two', 2),
            ('three', 3),
            ('four', 4),
            ('five', 5),
            ('six', 6),
            ('seven', 7),
            ('eight', 8),
            ('nine', 9),
            ]
    
    DIGITS.extend([(str(i), i) for i in range(10)])

    def get_digit(string: str, start: int):
        """Return the first digit encountered in the string natural order"""
    
        for digit, value in DIGITS:
            if string[start:].startswith(digit):
                return value
        return None
    
    for line in text:
        line = line.strip()

        first = None
        for i in range(len(line)):
            value = get_digit(line, i)
            if value:
                first = str(value)
                print(line, value)
                break

        second = None
        for i in range(len(line)-1, -1, -1):
            value = get_digit(line, i)
            if value:
                second = str(value)
                print(line, value)
                break

        callibrations.append(int(first + second))

    print(sum(callibrations))



