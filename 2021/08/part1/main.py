import sys


def main():
    total_1478 = 0
    for line in sys.stdin.readlines():
        _, signal_str = line.split('|')
        signals = signal_str.split()
        total_1478 += sum(1 for x in signals if len(x) in {2, 3, 4, 7})
    print(total_1478)


main()
