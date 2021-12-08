import sys


def main():
    total = 0
    for line in sys.stdin.readlines():
        digit_str, signal_str = line.split('|')
        digit_map = get_digit_map(''.join(sorted(x)) for x in digit_str.split())
        signals = [''.join(sorted(x)) for x in signal_str.split()]
        for i in range(4):
            total += digit_map[signals[i]] * 10**(3-i)
    print(total)


def get_digit_map(digits):
    length_map = {i: [] for i in range(2, 8)}
    for d in digits:
        length_map[len(d)].append(d)

    digit_map = {
            1: length_map[2][0],
            4: length_map[4][0],
            7: length_map[3][0],
            8: length_map[7][0],
            }

    for d in length_map[6]:
        if contains_all(d, digit_map[1]):
            if contains_all(d, digit_map[4]):
                digit_map[9] = d
            else:
                digit_map[0] = d
        else:
            digit_map[6] = d

    for d in length_map[5]:
        if contains_all(d, digit_map[1]):
            digit_map[3] = d
        else:
            if contains_all(digit_map[9], d):
                digit_map[5] = d
            else:
                digit_map[2] = d

    return {digit_map[k]: k for k in digit_map}


def contains_all(s, cs):
    for c in cs:
        if c not in s:
            return False
    return True


main()
