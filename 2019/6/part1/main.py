import sys


def main():
    graph = read_input()
    print(count_edges(graph, "COM"))


def read_input():
    graph = dict()
    for line in sys.stdin.readlines():
        u, v = line.strip().split(')')
        if u in graph:
            graph[u].add(v)
        else:
            graph[u] = {v}
        if v not in graph:
            graph[v] = set()
    return graph


def count_edges(graph, node, level=1):
    if len(graph[node]) == 0:
        return 0

    total = level * len(graph[node])
    for v in graph[node]:
        total += count_edges(graph, v, level+1)
    return total


main()
