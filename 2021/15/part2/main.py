from queue import PriorityQueue


def main():
    grid = get_grid_from_stdin()
    larger_grid = generate_larger_grid(grid)

    start, end = (0, 0), (len(larger_grid)-1, len(larger_grid[0])-1)
    path = dijkstra(larger_grid, start, end)
    total = 0
    for pos in path[1:]:
        r, c = pos
        total += larger_grid[r][c]
    print(total)


def dijkstra(grid, start, end):
    frontier = PriorityQueue()
    frontier.put((0, start))
    came_from = {start: None}
    d = {start: 0}

    while not frontier.empty():
        cost, u = frontier.get()

        if u == end:
            return construct_path(came_from, end)

        for v in get_neighbors(grid, u):
            r, c = v
            new_cost = cost + grid[r][c]
            if (v not in came_from) or (new_cost < d[v]):
                d[v] = new_cost
                frontier.put((new_cost, v))
                came_from[v] = u


def get_neighbors(grid, pos):
    neighbors = set()
    r, c = pos
    for m in ((1, 0), (-1, 0), (0, 1), (0, -1)):
        nr, nc = r+m[0], c+m[1]
        if (0 <= nr < len(grid)) and (0 <= nc < len(grid[0])):
            neighbors.add((nr, nc))
    return neighbors


def construct_path(came_from, u):
    path = []
    while u is not None:
        path = [u] + path
        u = came_from[u]
    return path


def get_grid_from_stdin():
    grid = []
    while True:
        try:
            grid.append([int(x) for x in input()])
        except EOFError:
            return grid


def generate_larger_grid(grid):
    larger_grid = []
    for _ in range(len(grid) * 5):
        larger_grid.append([0 for _ in range(len(grid[0]) * 5)])
    for r, row in enumerate(grid):
        for c, val in enumerate(row):
            larger_grid[r][c] = val

    for r, _ in enumerate(grid):
        generate_row(larger_grid, r)
    for c, _ in enumerate(larger_grid):
        generate_col(larger_grid, c)

    return larger_grid


def generate_row(grid, r):
    cols = len(grid[0]) // 5
    for c in range(cols):
        for i in range(1, 5):
            target_col = i*cols + c
            previous_col = (i-1)*cols + c
            grid[r][target_col] = grid[r][previous_col] + 1
            if grid[r][target_col] == 10:
                grid[r][target_col] = 1


def generate_col(grid, c):
    rows = len(grid) // 5
    for r in range(rows):
        for i in range(1, 5):
            target_row = i*rows + r
            previous_row = (i-1)*rows + r
            grid[target_row][c] = grid[previous_row][c] + 1
            if grid[target_row][c] == 10:
                grid[target_row][c] = 1


def print_grid(grid):
    for row in grid:
        for val in row:
            print(val, end='')
        print()


main()
