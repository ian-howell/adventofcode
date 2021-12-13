def main():
    grid = Grid()
    folds = get_folds_from_stdin()

    grid.fold(*folds[0])
    print(len(grid.points))


class Grid:
    def __init__(self):
        self.x_max = 0
        self.y_max = 0
        self.points = set()
        while True:
            line = input()
            if line == "":
                break
            point = tuple(int(x) for x in line.split(','))
            self.points.add(point)
            self.x_max = max(self.x_max, point[0]+1)
            self.y_max = max(self.y_max, point[1]+1)

    def fold(self, axis, val):
        new_points = set()
        if axis == 'y':
            for point in self.points:
                if point[1] <= val:
                    new_points.add(point)
                else:
                    new_points.add((point[0], 2*val - point[1]))
        else:
            for point in self.points:
                if point[0] <= val:
                    new_points.add(point)
                else:
                    new_points.add((2*val - point[0], point[1]))
        self.points = new_points

        self.x_max = 0
        self.y_max = 0
        for point in self.points:
            self.x_max = max(self.x_max, point[0]+1)
            self.y_max = max(self.y_max, point[1]+1)

    def print(self):
        for y in range(self.y_max):
            for x in range(self.x_max):
                if (x, y) in self.points:
                    print('#', end='')
                else:
                    print('.', end='')
            print()


def get_folds_from_stdin():
    folds = []
    while True:
        try:
            fold = input().split()[-1].split('=')
            folds.append((fold[0], int(fold[1])))
        except EOFError:
            return folds


main()
