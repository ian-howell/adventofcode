def main():
    fishes = [int(x) for x in input().split(',')]
    day_counter = [0] * 9
    for fish in fishes:
        day_counter[fish] += 1

    for i in range(256):
        number_spawning, day_counter = day_counter[0], day_counter[1:]
        day_counter[6] += number_spawning
        day_counter += [number_spawning]
    print(f'{sum(day_counter)}')


main()
