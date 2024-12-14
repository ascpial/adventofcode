import re
import queue

input = """Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279"""

with open('input.txt', 'r') as file:
    input = file.read()

class Machine:
    def __init__(self, ax, ay, bx, by, px, py):
        self.ax, self.ay = int(ax), int(ay)
        self.bx, self.by = int(bx), int(by)
        self.px, self.py = int(px) + 10000000000000, int(py) + 10000000000000

machines: list[Machine] = []
p = re.compile(r"Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)")
for bloc in input.split("\n\n"):
    data = p.match(bloc)
    machines.append(Machine(*data.groups()))

tokens = 0
for m in machines:
    n = m.px * m.ay - m.py * m.ax
    d = m.bx * m.ay - m.by * m.ax
    if n % d == 0:
        B = n // d
        if B > 0 and (m.py - B * m.by) % m.ay == 0:
            A = (m.py - B * m.by) // m.ay
            if A > 0:
                tokens += A * 3 + B

print(tokens)

