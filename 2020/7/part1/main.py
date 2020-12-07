import sys


class Edge:
    def __init__(self, node, weight):
        self.node = node
        self.weight = weight

    def __repr__(self):
        return f"{self.node, self.weight}"


def main():
    adj_list = {}
    for line in sys.stdin.readlines():
        outer, edges = parse_edge(line)
        adj_list[outer] = edges

    inverted_list = invert_adj_list(adj_list)
    print(count_descendants(inverted_list, "shiny gold"))


def parse_edge(s):
    outer, rest = s[:-2].split(" bags contain ")
    inners = rest.split(", ")
    edges = []
    for inner in inners:
        parts = inner.split()
        if parts[0] != 'no':
            number = int(parts[0])
            color = " ".join(parts[1:-1])
            edges.append(Edge(color, number))
    return outer, edges


def invert_adj_list(adj_list):
    inverted_list = {}
    for u in adj_list:
        if u not in inverted_list:
            inverted_list[u] = []
        for edge in adj_list[u]:
            if edge.node in inverted_list:
                inverted_list[edge.node].append(Edge(u, edge.weight))
            else:
                inverted_list[edge.node] = [Edge(u, edge.weight)]
    return inverted_list


def get_descendants(adj_list, node):
    if len(adj_list[node]) == 0:
        return set()

    descendants = set(x.node for x in adj_list[node])
    for child in adj_list[node]:
        descendants |= get_descendants(adj_list, child.node)

    return descendants


def count_descendants(adj_list, node):
    return len(get_descendants(adj_list, node))


if __name__ == "__main__":
    main()
