import sys


def main():
    grid = []
    for line in sys.stdin.readlines():
        row = []
        for val in line.strip():
            row.append(int(val))
        grid.append(row)

    low_points = []
    for r, row in enumerate(grid):
        for c, val in enumerate(row):
            if val < min(get_neighbors((r, c), grid)):
                low_points.append(val)
    print(sum(low_points) + len(low_points))


def get_neighbors(pos, grid):
    r, c = pos
    neighbors = []
    for diff in ((1, 0), (-1, 0), (0, 1),(0, -1)):
        new_r = r + diff[0]
        new_c = c + diff[1]
        if (0 <= new_r < len(grid)) and (0 <= new_c < len(grid[0])):
            neighbors.append(grid[new_r][new_c])
    return neighbors


main()
