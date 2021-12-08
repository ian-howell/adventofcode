import math


def main():
    vals = [int(x) for x in input().split(',')]
    leftmost = min(vals)
    rightmost = max(vals)
    best_fuel_cost = math.inf
    for i in range(len(vals)):
        fuel_cost = sum(abs(vals[j] - vals[i]) for j in range(len(vals)) if i != j)
        best_fuel_cost = min(best_fuel_cost, fuel_cost)
    print(best_fuel_cost)


main()
