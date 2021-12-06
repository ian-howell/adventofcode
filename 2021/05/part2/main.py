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
        start, end = self.a, self.b
        dx = end[0] - start[0]
        if dx != 0:
            dx = dx / abs(dx)
        dy = end[1] - start[1]
        if dy != 0:
            dy = dy / abs(dy)

        dx, dy = int(dx), int(dy)

        points = [start]
        p = start + (dx, dy)
        while points[-1] != end:
            p = (p[0] + dx, p[1] + dy)
            points.append(p)

        return points


main()
