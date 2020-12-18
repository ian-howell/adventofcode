import sys


def compute(tokens, pos):
    if pos == len(tokens)-1:
        return int(tokens[pos])

    val = 0
    if tokens[pos] == '(':
        val, pos = compute(tokens, pos+1)
    else:
        val = int(tokens[pos])

    i = pos + 1
    while i < len(tokens):
        if tokens[i] == ')':
            return val, i

        if tokens[i] == '+':
            i += 1
            if tokens[i] == '(':
                v, i = compute(tokens, i+1)
                val += v
            else:
                val += int(tokens[i])
        else:
            i += 1
            if tokens[i] == '(':
                v, i = compute(tokens, i+1)
                val *= v
            else:
                val *= int(tokens[i])
        i += 1
    return val, -1


total = 0
for line in sys.stdin.readlines():
    tokens = ''.join(line.split(' ')).strip()
    val, _ = compute(tokens, 0)
    total += val
print(total)
