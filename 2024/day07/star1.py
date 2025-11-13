input = """190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"""

with open('input.txt', 'r') as file:
    input = file.read()

equations = []
for line in input.strip().split("\n"):
    result, raw_numbers = line.split(": ")
    result = int(result)
    numbers = tuple(map(int, raw_numbers.split()))
    equations.append((result, numbers))

print(equations)

count = 0
for supposed_result, numbers in equations:
    def aux(result, i):
        if i == 0:
            return result == numbers[0]
        if result % numbers[i] == 0:
            if aux(result // numbers[i], i-1):
                return True
        if result - numbers[i] > 0:
            if aux(result - numbers[i], i-1):
                return True
        return False
    if aux(supposed_result, len(numbers)-1):
        count += supposed_result

print(count)
