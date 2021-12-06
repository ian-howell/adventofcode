from collections import defaultdict
import sys


def main():
    lines = []
    for line in sys.stdin.readlines():
        raw_a, raw_b = line.split(' -> ')
        a = tuple(int(x) for x in raw_a.split(','))
        b = tuple(int(x) for x in raw_b.split(','))
        lines.append(Line(a, b))

    intersecting_points = defaultdict(int)
    for line in lines:
        for point in line.get_points():
            intersecting_points[point] += 1

    score = 0
    for point in intersecting_points:
        if intersecting_points[point] > 1:
            score += 1
    print(score)


class Line:
    def __init__(self, a, b):
        self.a = a
        self.b = b

    def get_points(self):
        if self.a[0] == self.b[0]:
            index = 1
        elif self.a[1] == self.b[1]:
            index = 0
        else:
            return []
        start, end = self.a[index], self.b[index]
        if end < start:
            start, end = end, start
        points = []
        for i in range(start, end+1):
            if index == 1:
                points.append((self.a[0], i))
            else:
                points.append((i, self.b[1]))
        return points

main()
