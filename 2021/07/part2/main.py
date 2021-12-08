import math


def main():
    vals = [int(x) for x in input().split(',')]
    leftmost = min(vals)
    rightmost = max(vals)
    best_fuel_cost = math.inf
    for i in range(leftmost, rightmost+1):
        total_fuel_cost = 0
        for j in range(len(vals)):
            distance = abs(vals[j] - i)
            fuel_cost = sum_to(distance)
            total_fuel_cost += fuel_cost
        best_fuel_cost = min(best_fuel_cost, total_fuel_cost)
    print(best_fuel_cost)


def sum_to(n):
    return n*(n+1)//2


main()
