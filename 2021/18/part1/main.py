import sys


class Node:
    def __init__(self, left, right, parent = None):
        if type(left) == list:
            self.left = Node(left[0], left[1], self)
        else:
            self.left = left
        if type(right) == list:
            self.right = Node(right[0], right[1], self)
        else:
            self.right = right
        self.parent = parent

    def add_to_left(self, x):
        if type(self.left) == int:
            self.left += x
            return
        self.left.add_to_left(x)

    def add_to_right(self, x):
        if type(self.right) == int:
            self.right += x
            return
        self.right.add_to_right(x)

    def __repr__(self):
        return f"[{self.left}, {self.right}]"


def main():
    orig_snail_numbers = get_snail_numbers()
    total = orig_snail_numbers[0]
    for i in range(1, len(orig_snail_numbers)):
        total = snail_add(total, orig_snail_numbers[i])
    print(total)
    print(magnitude(total))


def get_snail_numbers():
    numbers = []
    for line in sys.stdin.readlines():
        as_list = eval(line.strip())
        sn = Node(as_list[0], as_list[1])
        numbers.append(sn)
    return numbers


def snail_add(a, b):
    sn = Node(a, b)
    a.parent = sn
    b.parent = sn
    simplify(sn)
    return sn


def simplify(sn):
    still_working = True
    while still_working:
        sn, still_working = simplify_step(sn)
        # print(sn)
    return sn


def simplify_step(sn):
    # print("trying to explode")
    sn, done = explode_step(sn)
    if done:
        return sn, done
    # print("done exploding...")
    return split_step(sn)


def explode_step(sn, level=0):
    if type(sn) != Node:
        return sn, False

    if level == 4:
        # print('exploding: ', end='')
        explode(sn)
        return 0, True

    sn.left, done = explode_step(sn.left, level + 1)
    if not done:
        sn.right, done = explode_step(sn.right, level + 1)
    return sn, done


def split_step(sn, level=0):
    if type(sn) == Node:
        sn.left, done = split_step(sn.left, level + 1)
        if not done:
            sn.right, done = split_step(sn.right, level + 1)
        if type(sn.left) == Node:
            sn.left.parent = sn
        if type(sn.right) == Node:
            sn.right.parent = sn
        return sn, done
    elif (type(sn) == int) and sn >= 10:
        smaller = sn // 2
        # print('splitting: ', end='')
        return Node(smaller, sn - smaller), True

    return sn, False


def explode(sn):
    parent = sn.parent
    child = sn
    while (parent is not None) and (parent.left == child):
        child = parent
        parent = parent.parent

    if parent is not None:
        if type(parent.left) == int:
            parent.left += sn.left
        else:
            parent.left.add_to_right(sn.left)

    parent = sn.parent
    child = sn
    while (parent is not None) and (parent.right == child):
        child = parent
        parent = parent.parent

    if parent is not None:
        if type(parent.right) == int:
            parent.right += sn.right
        else:
            parent.right.add_to_left(sn.right)


def magnitude(sn):
    if type(sn) == int:
        return sn
    return 3*magnitude(sn.left) + 2*magnitude(sn.right)


main()
