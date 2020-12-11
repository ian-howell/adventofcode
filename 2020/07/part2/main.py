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
    print(count_descendants(adj_list, "shiny gold"))


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


def count_descendants(adj_list, node):
    if len(adj_list[node]) == 0:
        return 0

    total = 0
    for child in adj_list[node]:
        total += child.weight + child.weight * count_descendants(adj_list, child.node)

    return total


if __name__ == "__main__":
    main()
