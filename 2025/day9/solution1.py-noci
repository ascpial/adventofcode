print(
    (lambda points:
        (lambda ps: (abs(ps[0][0]-ps[1][0])+1)*(abs(ps[0][1]-ps[1][1])+1)
        (sorted([(points[i], points[j]) for i in range(len(points)) for j in range(i+1, len(points))], key=lambda ps: (abs(ps[0][0]-ps[1][0])+1)*(abs(ps[0][1]-ps[1][1])+1), reverse=True))[0]
    )
    ([(int(rawpoint[0]), int(rawpoint[1])) for rawpoint in (rawrawpoint.split(",") for rawrawpoint in open('input.txt', 'r').read().strip().split('\n'))])
)
