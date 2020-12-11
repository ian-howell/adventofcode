def main():
    wire1_instructions = input().split(',')
    wire2_instructions = input().split(',')

    wire1_path = get_path(wire1_instructions)
    wire2_path = get_path(wire2_instructions)

    intersections = wire1_path & wire2_path

    distances1 = get_distances(wire1_instructions, intersections)
    distances2 = get_distances(wire2_instructions, intersections)

    closest = min([distances1[x] + distances2[x] for x in intersections])
    print(closest)


def get_path(instructions):
    pos = (0, 0)
    path = {pos}
    for instr in instructions:
        direction = get_direction(instr)
        for i in range(1, int(instr[1:])+1):
            pos = add_tuple(pos, direction)
            path.add(pos)
    path.remove((0, 0))
    return path


def add_tuple(x, y):
    return x[0] + y[0], x[1] + y[1]


def get_direction(instr):
    if instr[0] == 'R':
        return 1, 0
    if instr[0] == 'L':
        return -1, 0
    if instr[0] == 'U':
        return 0, 1
    if instr[0] == 'D':
        return 0, -1


def get_distances(instructions, intersections):
    distances = {}
    pos = (0, 0)
    steps = 0
    for instr in instructions:
        direction = get_direction(instr)
        for i in range(1, int(instr[1:])+1):
            steps += 1
            pos = add_tuple(pos, direction)
            if pos in intersections and pos not in distances.keys():
                distances[pos] = steps
    return distances


main()
