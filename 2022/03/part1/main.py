import sys


def main():
    total = 0
    for line in sys.stdin.readlines():
        total += get_value(find_common(line.strip()))
    print(total)


def find_common(line):
    return (set(line[:len(line)//2]) & set(line[len(line)//2:])).pop()


def get_value(c):
    return ((ord(c)&0xdf)-0x26) - (((ord(c)&0x20)>>5)*26)


if __name__ == "__main__":
    main()
