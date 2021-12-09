import sys


def main():
    grid = []
    for line in sys.stdin.readlines():
        row = []
        for val in line.strip():
            row.append(int(val))
        grid.append(row)

    # a basin is described by 2 attributes, size and total of values
    largest_basins = []
    for r, row in enumerate(grid):
        for c, val in enumerate(row):
            if val != 9:
                basin = dfs((r, c), grid)
                largest_basins = list(sorted(largest_basins + [basin]))[-3:]
    print(product(largest_basins))


def get_neighbors(pos, grid):
    r, c = pos
    neighbors = []
    for diff in ((1, 0), (-1, 0), (0, 1),(0, -1)):
        new_r = r + diff[0]
        new_c = c + diff[1]
        if (0 <= new_r < len(grid)) and (0 <= new_c < len(grid[0])) and (grid[new_r][new_c] != 9):
            neighbors.append((new_r, new_c))
    return neighbors


def dfs(pos, grid):
    basin = 0
    frontier = [pos]
    visited = {pos}
    while len(frontier) != 0:
        u, frontier = frontier[-1], frontier[:-1]
        r, c = u
        basin += 1
        grid[r][c] = 9
        for v in get_neighbors(u, grid):
            if v not in visited:
                visited.add(v)
                frontier.append(v)
    return basin


def product(iterable):
    total = 1
    for x in iterable:
        total *= x
    return total


main()
