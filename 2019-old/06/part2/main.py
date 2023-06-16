import sys


def main():
    graph = read_input()
    inverted_graph = invert_graph(graph)
    common_ancestor = find_common_ancestor(inverted_graph, "YOU", "SAN")
    from_ancestor_to_you = path(inverted_graph, "YOU", common_ancestor)
    from_ancestor_to_santa = path(inverted_graph, "SAN", common_ancestor)
    print(len(from_ancestor_to_you) + len(from_ancestor_to_santa) - 2)


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


def invert_graph(graph):
    inverted_graph = dict()
    for u in graph:
        for v in graph[u]:
            inverted_graph[v] = u
        if u not in inverted_graph:
            inverted_graph[u] = set()
    return inverted_graph


def path(graph, a, b):
    if graph[a] != b:
        return path(graph, graph[a], b) + [a]
    return [a]


def find_common_ancestor(graph, a, b):
    path_to_a = path(graph, a, "COM")
    path_to_b = path(graph, b, "COM")

    i = 0
    while i < len(path_to_a) and i < len(path_to_b) and path_to_a[i] == path_to_b[i]:
        i += 1
    return path_to_a[i - 1]


main()
