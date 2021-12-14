from collections import defaultdict
import math
from pprint import pprint


def main():
    polymer_template = input()
    # eat the empty line
    input()

    pairs = get_pairs_from_stdin()

    for i in range(10):
        polymer_template = insert(polymer_template, pairs)

    frequencies = defaultdict(int)
    for letter in polymer_template:
        frequencies[letter] += 1

    max_value, min_value = -math.inf, math.inf
    for index in frequencies:
        max_value = max(frequencies[index], max_value)
        min_value = min(frequencies[index], min_value)
    print(max_value - min_value)


def get_pairs_from_stdin():
    pairs = {}
    while True:
        try:
            items = input().split(' -> ')
            pairs[items[0]] = items[1]
        except EOFError:
            return pairs


def insert(template, pairs):
    new_template = template[0]
    for i in range(1, len(template)):
        pair = template[i-1:i+1]
        new_template += pairs[pair] + template[i]
    return new_template


main()
