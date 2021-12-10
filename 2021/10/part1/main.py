import sys


def main():
    score = 0
    for line in sys.stdin.readlines():
        broken_char = get_broken_char(line.strip())
        if broken_char is not None:
            score += {')': 3, ']': 57, '}': 1197, '>': 25137}[broken_char]
    print(score)


def get_broken_char(s):
    # This only looks for mismatched brackets; it doesn't care about closing
    # brackets that don't have a match at all.
    # All inputs contain either a mismatch or an incomplete, so we never need
    # to check the length of the stack.
    stack = []
    for c in s:
        if c in '([{<':
            stack.append(c)
        elif matches(c, stack[-1]):
            stack.pop()
        else:
            return c
    return None


def matches(closing, opening):
    return {'(': ')', '[': ']', '{': '}', '<': '>'}[opening] == closing


main()
