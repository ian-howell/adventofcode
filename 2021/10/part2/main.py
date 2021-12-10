import sys


def main():
    scores = []
    for line in sys.stdin.readlines():
        completing_chars = get_completing_chars(line.strip())
        if len(completing_chars) > 0:
            score = 0
            for char in completing_chars:
                score = (score * 5) + {')': 1, ']': 2, '}': 3, '>': 4}[char]
            scores.append(score)
    scores.sort()
    print(scores[len(scores) // 2])


def get_completing_chars(s):
    # This only looks for mismatched brackets; it doesn't care about closing
    # brackets that don't have a match at all.
    # All inputs contain either a mismatch or an incomplete, so we never need
    # to check the length of the stack.
    stack = []
    matches = {'(': ')', '[': ']', '{': '}', '<': '>'}
    for c in s:
        if c in '([{<':
            stack.append(c)
        elif matches[stack[-1]] == c:
            stack.pop()
        else:
            return ""
    rs = ""
    while len(stack) > 0:
        rs += matches[stack.pop()]
    return rs


main()
