def main():
    total = 0
    for _ in range(100):
        total += fuel_per_module(int(input()))
    print(total)


def fuel_per_module(weight):
    total = 0
    required_fuel = weight // 3 - 2
    while required_fuel > 0:
        total += required_fuel
        required_fuel = required_fuel // 3 - 2
    return total

main()
