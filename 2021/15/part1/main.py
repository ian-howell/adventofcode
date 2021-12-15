from queue import PriorityQueue


def main():
    grid = get_grid_from_stdin()
    start, end = (0, 0), (len(grid)-1, len(grid[0])-1)
    path = dijkstra(grid, start, end)
    total = 0
    for pos in path[1:]:
        r, c = pos
        total += grid[r][c]
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


main()
