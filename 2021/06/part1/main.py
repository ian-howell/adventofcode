def main():
    current_fishes = [int(x) for x in input().split(',')]
    for i in range(80):
        next_fishes = []
        for fish in current_fishes:
            if fish == 0:
                next_fishes += [6, 8]
            else:
                next_fishes += [fish-1]
        current_fishes = next_fishes
    print(len(current_fishes))


main()
