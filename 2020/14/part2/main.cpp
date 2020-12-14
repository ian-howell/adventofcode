#include <iostream>
#include <map>
#include <string>
#include <vector>

void set_bit(unsigned long& num, int bit_no);
void unset_bit(unsigned long& num, int bit_no);
void print_mask(unsigned long mask);
unsigned long calc_total(const std::map<unsigned long, unsigned long>& mem);
void generate_masks(const std::string& s, size_t i, unsigned long mask, std::vector<unsigned long>& masks);

std::vector<std::string> split(const std::string& s, char delim);

int main() {
    std::string line;
    std::map<unsigned long, unsigned long> mem;
    std::string mask;
    while(std::getline(std::cin, line)) {
        std::vector<std::string> parts = split(line, ' ');
        if (parts[0] == "mask") {
            mask = parts[2];
        } else {
            std::vector<unsigned long> masks;
            unsigned long loc = std::stol(split(split(parts[0], '[')[1], ']')[0]);
            generate_masks(mask, 0, loc, masks);
            unsigned long val = std::stol(parts[2]);
            for (const auto& it : masks) {
                mem[it] = val;
            }
        }
    }

    std::cout << calc_total(mem) << std::endl;
}

void set_bit(unsigned long& num, int bit_no) { num |= 1ul << bit_no; }
void unset_bit(unsigned long& num, int bit_no) { num &= (0xffffffffful ^ (1ul << bit_no)); }

void print_mask(unsigned long mask) {
    for (int i = 35; i >= 0; i--) {
        if (i % 8 == 7) {
            std::cout << " ";
        }
        std::cout << ((mask & (1ul << i)) >> i);
    }
    std::cout << std::endl;
}

unsigned long calc_total(const std::map<unsigned long, unsigned long>& mem) {
    unsigned long sum = 0ul;
    for (const auto& entry : mem) {
        sum += entry.second;
    }
    return sum;
}

std::vector<std::string> split(const std::string& s, char delim) {
    size_t lb = 0;
    size_t rb;
    std::vector<std::string> splits;
    do {
        rb = s.find(delim, lb);
        splits.push_back(s.substr(lb, (rb - lb)));
        lb = rb + 1;
    } while (rb != std::string::npos);
    return splits;
}

void generate_masks(const std::string& s, size_t i, unsigned long mask, std::vector<unsigned long>& masks) {
    if (i == s.size()) {
        masks.push_back(mask);
        return;
    }

    switch (s.at(i)) {
        case '1':
            set_bit(mask, 35 - i);
        case '0':
            generate_masks(s, i+1, mask, masks);
            break;
        default:
            unsigned long high_mask = mask;
            unsigned long low_mask = mask;

            set_bit(high_mask, 35 - i);
            unset_bit(low_mask, 35 - i);

            generate_masks(s, i+1, low_mask, masks);
            generate_masks(s, i+1, high_mask, masks);
    }
}
