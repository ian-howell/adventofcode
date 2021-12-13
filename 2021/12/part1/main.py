import sys


def main():
    graph = get_graph_from_stdin()
    print(dfs(graph))


def get_graph_from_stdin():
    adj_set = {}
    for line in sys.stdin.readlines():
        to_cave, from_cave = line.strip().split("-")
        if to_cave not in adj_set:
            adj_set[to_cave] = {from_cave}
        else:
            adj_set[to_cave].add(from_cave)
        if from_cave not in adj_set:
            adj_set[from_cave] = {to_cave}
        else:
            adj_set[from_cave].add(to_cave)
    # The end cave is a sink
    adj_set['end'] = set()
    return adj_set


def dfs(graph):
    return _r_dfs(graph, 'start', {'start'})


def _r_dfs(graph, u, visited):
    if u == 'end':
        return 1
    paths = 0
    for v in graph[u]:
        if v not in visited:
            new_visited = {x for x in visited}
            if v.islower():
                new_visited |= {v}
            paths += _r_dfs(graph, v, new_visited)
    return paths


main()
