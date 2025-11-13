import re
import itertools

input = """p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3"""
w = 11
h = 7

with open('input.txt', 'r') as file:
    input = file.read()
    w = 101
    h = 103

class Robot:
    def __init__(self, w, h, posx, posy, vx, vy):
        self.posx, self.posy = int(posx), int(posy)
        self.vx, self.vy = int(vx), int(vy)
        self.w, self.h = w, h

    def move(self):
        self.posx = (self.posx + self.vx) % self.w
        self.posy = (self.posy + self.vy) % self.h

p = re.compile(r"p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)")
robots = []
for line in input.strip().split("\n"):
    robots.append(
            Robot(
            w, h,
            *p.match(line).groups(),
            ))

def show_map():
    map = [[0 for _ in range(w)] for _ in range(h)]
    for robot in robots:
        map[robot.posy][robot.posx] += 1

    print('\n'.join([''.join(['.' if pos == 0 else hex(pos)[2:] for pos in line]) for line in map]))
    print('-'*w)

show_map()

for _ in range(100):
    for robot in robots:
        robot.move()

result = 1
n = 2
for qx, qy in itertools.product(range(n), range(n)):
    robot_n = 0
    for robot in robots:
        if qy*(h//n+1) <= robot.posy < h//2 + qy * (h//2+1) and qx*(w//2+1) <= robot.posx < w//2 + qx * (w//2+1):
            robot_n += 1
    result *= robot_n

show_map()
print(result)
