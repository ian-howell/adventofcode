from collections import defaultdict
import math


def main():
    polymer_template = input()
    # eat the empty line
    input()

    template_pairs = create_template_pairs(polymer_template)
    rule_pairs = get_pairs_from_stdin()

    for i in range(40):
        template_pairs = insert(template_pairs, rule_pairs)

    frequencies = get_polymer_frequencies(template_pairs)

    min_value, max_value = math.inf, -math.inf
    for k in frequencies:
        min_value = min(min_value, frequencies[k])
        max_value = max(max_value, frequencies[k])
    print(max_value - min_value)


def create_template_pairs(template):
    pairs = defaultdict(int)
    for i in range(1, len(template)):
        pairs[template[i-1:i+1]] += 1
    return pairs


def get_pairs_from_stdin():
    pairs = {}
    while True:
        try:
            items = input().split(' -> ')
            pairs[items[0]] = items[1]
        except EOFError:
            return pairs


def insert(template_pairs, rule_pairs):
    new_template_pairs = defaultdict(int)
    for k in template_pairs:
        generated_polymer = rule_pairs[k]
        new_pair1 = k[0] + generated_polymer
        new_pair2 = generated_polymer + k[1]
        new_template_pairs[new_pair1] += template_pairs[k]
        new_template_pairs[new_pair2] += template_pairs[k]
    return new_template_pairs


def get_polymer_frequencies(template_pairs):
    frequencies = defaultdict(int)
    for k in template_pairs:
        frequencies[k[0]] += template_pairs[k]
        frequencies[k[1]] += template_pairs[k]
    return {k: int(frequencies[k]/2 + 0.5) for k in frequencies}


main()
