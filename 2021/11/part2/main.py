import sys


def main():
    grid = []
    for line in sys.stdin.readlines():
        row = []
        for val in line.strip():
            row.append(int(val))
        grid.append(row)

    i = 0
    while True:
        i += 1
        flashes_this_step = step(grid)
        if flashes_this_step == len(grid) * len(grid[0]):
            print(i)
            return


def step(grid):
    flashes = 0
    increment(grid)
    for r, row in enumerate(grid):
        for c in range(len(row)):
            if grid[r][c] == 10:
                flashes += propogate(grid, r, c)
    reset(grid)
    return flashes


def increment(grid):
    for r, row in enumerate(grid):
        for c in range(len(row)):
            grid[r][c] += 1


def propogate(grid, r, c):
    flashes = 1
    grid[r][c] = 11
    for neighbor in get_neighbors(grid, r, c):
        new_r, new_c = neighbor
        if grid[new_r][new_c] <= 9:
            grid[new_r][new_c] += 1
        if grid[new_r][new_c] == 10:
            flashes += propogate(grid, new_r, new_c)
    return flashes


def reset(grid):
    for r, row in enumerate(grid):
        for c in range(len(row)):
            if grid[r][c] >= 10:
                grid[r][c] = 0


def get_neighbors(grid, r, c):
    neighbors = []
    for m in ((+0, +1), (+0, -1), (+1, +0), (-1, +0), (+1, +1), (-1, -1), (+1, -1), (-1, +1)):
        new_r, new_c = r+m[0], c+m[1]
        if (0 <= new_r < len(grid)) and (0 <= new_c < len(grid[0])) and (grid[new_r][new_c] <= 9):
            neighbors.append((new_r, new_c))
    return neighbors


main()
