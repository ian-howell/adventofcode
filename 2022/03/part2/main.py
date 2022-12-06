import sys


def main():
    total = 0
    lines = [set(line.strip()) for line in sys.stdin.readlines()]
    for i in range(0, len(lines), 3):
        total += get_value(find_common(lines[i:i+3]).pop())
    print(total)


def find_common(groups):
    return groups[0] & groups[1] & groups[2]


def get_value(c):
    return ((ord(c)&0xdf)-0x26) - (((ord(c)&0x20)>>5)*26)


if __name__ == "__main__":
    main()
