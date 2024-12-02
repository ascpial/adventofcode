with open("input.txt", 'r', encoding='utf-8') as file:
    text = [line.strip() for line in file.readlines()]

DIGITS = [str(i) for i in range(10)]

class Grid:
    def __init__(self, data: list[str]):
        self.data = data
        self.max_x = len(data[0])
        self.max_y = len(data)
        print(f"The grid is of size {self.max_x}, {self.max_y}")

    def __getitem__(self, pos):
        x, y = pos
        if (0 <= x < self.max_x) and (0 <= y < self.max_y):
            return self.data[y][x]
        else:
            return "."

    def is_digit(self, x, y):
        return self[x, y] in DIGITS
    
    def is_symbol(self, x, y):
        return self[x, y] not in DIGITS and self[x, y] != "."

    def number_pos(self, x, y):
        """Retourne la valeur, la position et la longueur d'un éventuel nombre dont un chiffre est localisé à l'emplacement indiqué"""
        if not self.is_digit(x, y):
            return None
        start = x
        len = 1
        while self.is_digit(start - 1, y):
            start -= 1
            len += 1
        while self.is_digit(start + len, y):
            len += 1

        return (int(self.data[y][start:start+len]), (start, y), len)
    
    def box_contains_symbol(self, x, y, len):
        """Retourne True si la boîte entourant le texte à l'emplacement indiqué contient un symbol"""
        touches = False
        for x2 in range(x-1, x+len+1):
            touches = touches or self.is_symbol(x2, y-1)
            touches = touches or self.is_symbol(x2, y+1)
        touches = touches or self.is_symbol(x-1, y)
        touches = touches or self.is_symbol(x+len, y)

        return touches

    def touches_symbole(self, x, y):
        """Retourne le nombre à l'emplacement indiqué si il est entouré par un symbol"""
        data = self.number_pos(x, y)
        if data is None:
            return None
        id, pos, len = data
        if self.box_contains_symbol(*pos, len):
            return id, pos[0] + len - x - 1
        return None

    def touches_two_numbers(self, x, y):
        """Retourne les valeurs des nombres des deux uniques nombres situés autour du symbole ou None si le cas de figure n'est pas respecté"""
        adjacent_parts = set()
        for y2 in range(y-1, y+2):
            for x2 in range(x-1, x+2):
                value = self.number_pos(x2, y2)
                if value:
                    adjacent_parts.add(value[1]) # append the unique number pos

        if len(adjacent_parts) != 2:
            return None
        else:
            numbers = []
            for x, y in adjacent_parts:
                numbers.append(self.number_pos(x, y)[0])
            return numbers

    def show_number(self, x, y):
        """Shows the box containing the number if applicable"""
        if not self.is_digit(x, y):
            return

        id, pos, len = self.number_pos(x, y)
        x, y = pos
        for y2 in range(y-1, y+2):
            for x2 in range(x-1, x+len+1):
                print(self[x2, y2], end="")
            print()

        print()

grid = Grid(text)

if False:
    parts = []

    x, y = 0, 0
    while x < grid.max_x and y < grid.max_y:
        x += 1
        if x >= grid.max_x:
            x = 0
            y += 1

        value = grid.touches_symbole(x, y)
        if value:
            id, to_forward = value
            x += to_forward
            parts.append(id)

    
    print(parts)
    print(parts.count(395))
    print(sum(parts))

if True:
    ratios = []
    for y in range(grid.max_y):
        for x in range(grid.max_x):
            if grid[x, y] == "*":
                gears = grid.touches_two_numbers(x, y)
                if gears:
                    ratio = gears[0] * gears[1]
                    ratios.append(ratio)
    print(sum(ratios))
