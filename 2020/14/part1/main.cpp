#include <iostream>
#include <map>
#include <string>
#include <vector>

void set_bit(unsigned long& num, int bit_no);
void unset_bit(unsigned long& num, int bit_no);
unsigned long create_or(std::string raw_mask);
unsigned long create_and(std::string raw_mask);
void print_mask(unsigned long mask);
std::map<unsigned long, unsigned long> read_memory();
unsigned long calc_total(const std::map<unsigned long, unsigned long>& mem);

std::vector<std::string> split(const std::string& s, char delim);

int main() {
    std::map<unsigned long, unsigned long> mem = read_memory();
    std::cout << calc_total(mem) << std::endl;
}

void set_bit(unsigned long& num, int bit_no) { num |= 1ul << bit_no; }
void unset_bit(unsigned long& num, int bit_no) { num &= (0xffffffffful ^ (1ul << bit_no)); }

unsigned long create_or(std::string raw_mask) {
    unsigned long mask = 0ul;
    for (size_t i = 0; i < raw_mask.size(); i++) {
        if (raw_mask.at(35 - i) == '1') {
            set_bit(mask, i);
        }
    }
    return mask;
}

unsigned long create_and(std::string raw_mask) {
    unsigned long mask = 0ul;
    for (size_t i = 0; i < raw_mask.size(); i++) {
        if (raw_mask.at(35 - i) != '0') {
            set_bit(mask, i);
        }
    }
    return mask;
}

void print_mask(unsigned long mask) {
    /* for (int i = 63; i >= 0; i--) { */
    for (int i = 35; i >= 0; i--) {
        if (i != 53 && i % 8 == 7) {
            std::cout << " ";
        }
        std::cout << ((mask & (1ul << i)) >> i);
    }
    std::cout << std::endl;
}

std::map<unsigned long, unsigned long> read_memory() {
    std::string line;
    unsigned long or_mask;
    unsigned long and_mask;
    std::map<unsigned long, unsigned long> mem;
    while(std::getline(std::cin, line)) {
        std::vector<std::string> parts = split(line, ' ');
        if (parts[0] == "mask") {
            or_mask = create_or(parts[2]);
            and_mask = create_and(parts[2]);
        } else {
            unsigned long loc = std::stol(split(split(parts[0], '[')[1], ']')[0]);
            mem[loc] = std::stol(parts[2]);
            mem[loc] &= and_mask;
            mem[loc] |= or_mask;
        }
    }
    return mem;
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
