with open('input.txt', 'r', encoding='utf-8') as file:
    input = file.read()

class Interval:
    def __init__(self, start, end):
        self.start = start
        self.end = end

    @classmethod
    def from_length(cls, start, length) -> "Interval":
        return cls(start, start + length - 1)

    def intersect(self, interval: "Interval"):
        return self.end < interval.start or interval.end < self.start

    def intersection(self, interval: "Interval"):
        """Retourne l'intersection de deux intervalles et None si l'intersection est vide"""
        if not self.intersect(interval):
            return None

        return self.__class__(
                max(self.start, interval.start),
                min(self.end, interval.end),
            )

    def __gt__(self, interval: "Interval") -> bool:
        return self.start > interval.end

    def copy(self) -> "Interval":
        return self.__class__(self.start, self.end)

class MappedInterval(Interval):
    def __init__(self, start, end, target):
        super().__init__(start, end)
        self.target = target

    @classmethod
    def from_length(cls, start, length, target) -> Interval:
        return cls(start, start + length - 1, target)

    def maps_to(self, element: int | Interval) -> int | Interval:
        if isinstance(element, int):
            return self.target + element - (self.end - self.start + 1)
        elif isinstance(element, Interval):
            return Interval(self.maps_to(element.start), self.maps_to(element.end))

class Map:
    def __init__(self, intervals, name):
        self.data: list[MappedInterval] = intervals
        self.name = name
    
    def image(self, interval: Interval):
        interval = interval.copy()

        for map in self.data:
            if map >
            if map.intersect(interval):
                yield

categories = input.split("\n\n")
seeds = [int(i) for i in categories[0].split(": ")[1].split(" ")]

maps = []

for category in categories[1:]:
    category = category.strip()
    lines = category.split("\n")
    name = lines[0][:-5]
    intervals = []
    for line in lines[1:]:
        target, source, length = (int(i) for i in line.split(" "))
        intervals.append(MappedInterval.from_length(source, length, target))
    intervals.sort(key=lambda interval: interval.start)
    maps.append(Map(intervals, name))
