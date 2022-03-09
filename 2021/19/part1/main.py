from collections import defaultdict
from pprint import pprint


def main():
    scanners = get_scanners()
    d0map = {}
    d1map = {}
    for i, xi in enumerate(scanners[0]):
        for j, xj in enumerate(scanners[0][i+1:], i+1):
            d = euclidean_distance(xi, xj)
            d0map[int(1e4 * d)] = (i, j)

    for i, xi in enumerate(scanners[1]):
        for j, xj in enumerate(scanners[1][i+1:], i+1):
            d = euclidean_distance(xi, xj)
            d1map[int(1e4 * d)] = (i, j)

    common_beacons_from_scanner0s_pov = {}
    for d in d0map:
        if d in d1map:
            x = d0map[d]
            y = d1map[d]
            x_pair = (scanners[0][x[0]], scanners[0][x[1]])
            y_pair = (scanners[0][y[0]], scanners[0][y[1]])
            common_beacons_from_scanner0s_pov[x_pair] = y_pair
    pprint(common_beacons_from_scanner0s_pov)



def get_scanners():
    scanners = []
    count = 1
    while True:
        try:
            scanners.append(get_one_scanner())
            count += 1
        except EOFError:
            return scanners


def get_one_scanner():
    # Eat the first line
    input()

    scanner = []
    while True:
        line = input()
        if not line.strip():
            break
        scanner.append(tuple(int(x) for x in line.split(",")))
    return scanner


def euclidean_distance(a, b):
    return ((a[0] - b[0])**2 + (a[1] - b[1])**2 + (a[2] - b[2])**2)**(1/2)


def rotate(beacon, m):
    return (
            m[0][0]*beacon[0] + m[0][1]*beacon[1] + m[0][2]*beacon[2],
            m[1][0]*beacon[0] + m[1][1]*beacon[1] + m[1][2]*beacon[2],
            m[2][0]*beacon[0] + m[2][1]*beacon[1] + m[2][2]*beacon[2]
            )


def rotatation_matrices = (
    (
        (1, 0, 0),
        (0, 1, 0),
        (0, 0, 1)
    ),
    (
        (1, 0, 0),
        (0, 1, 0),
        (0, 0, -1)
    ),
    (
        (1, 0, 0),
        (0, -1, 0),
        (0, 0, 1)
    ),
    (
        (1, 0, 0),
        (0, -1, 0),
        (0, 0, -1)
    ),
    (
        (-1, 0, 0),
        (0, 1, 0),
        (0, 0, 1)
    ),
    (
        (-1, 0, 0),
        (0, 1, 0),
        (0, 0, -1)
    ),
    (
        (-1, 0, 0),
        (0, -1, 0),
        (0, 0, 1)
    ),
    (
        (-1, 0, 0),
        (0, -1, 0),
        (0, 0, -1)
    ),


    (
        (1, 0, 0),
        (0, 0, 1),
        (0, 1, 0)
    ),
    (
        (1, 0, 0),
        (0, 0, 1),
        (0, -1, 0)
    ),
    (
        (1, 0, 0),
        (0, 0, -1),
        (0, 1, 0)
    ),
    (
        (1, 0, 0),
        (0, 0, -1),
        (0, -1, 0)
    ),

    (
        (-1, 0, 0),
        (0, 1, 0),
        (0, 0, 1)
    ),
    (
        (-1, 0, 0),
        (0, 1, 0),
        (0, 0, -1)
    ),
    (
        (-1, 0, 0),
        (0, 0, 1),
        (0, 1, 0)
    ),
    (
        (-1, 0, 0),
        (0, 0, 1),
        (0, -1, 0)
    ),
    (
        (-1, 0, 0),
        (0, 0, -1),
        (0, 1, 0)
    ),
    (
        (-1, 0, 0),
        (0, 0, -1),
        (0, -1, 0)
    ),

    (
        (0, 1, 0),
        (1, 0, 0),
        (0, 0, 1)
    ),
    (
        (0, 1, 0),
        (1, 0, 0),
        (0, 0, -1)
    ),
    (
        (0, 1, 0),
        (-1, 0, 0),
        (0, 0, 1)
    ),
    (
        (0, 1, 0),
        (-1, 0, 0),
        (0, 0, -1)
    ),
    (
        (0, 1, 0),
        (0, 0, 1),
        (1, 0, 0)
    ),
    (
        (0, 1, 0),
        (0, 0, 1),
        (-1, 0, 0)
    ),
    (
        (0, 1, 0),
        (0, 0, -1),
        (1, 0, 0)
    ),
    (
        (0, 1, 0),
        (0, 0, -1),
        (-1, 0, 0)
    ),
)

main()
