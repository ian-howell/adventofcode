#include <iostream>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

struct Range {
    int lo;
    int hi;
};
Range create_range(const std::string& s);
std::ostream& operator<<(std::ostream& out, const Range& range);
bool in_range(int x, const Range& range);

using Rule = std::vector<Range>;
Rule create_rule(const std::string& s);
std::ostream& operator<<(std::ostream& out, const Rule& rule);

using Ticket = std::vector<int>;
Ticket create_ticket(const std::string& s);
std::ostream& operator<<(std::ostream& out, const Ticket& ticket);

std::vector<std::string> split(const std::string& s, char delim);

int main() {
    std::string line;
    std::unordered_map<std::string, Rule> rules;
    std::getline(std::cin, line);
    while(line != "") {
        std::vector<std::string> parts = split(line, ':');
        rules[parts[0]] = create_rule(parts[1]);
        std::getline(std::cin, line);
    }

    // Eat the "your ticket:" line
    std::getline(std::cin, line);
    // Get my ticket's info (not used in part 1)
    std::getline(std::cin, line);

    // Eat the empty line
    std::getline(std::cin, line);
    // Eat the "nearby tickets:" line
    std::getline(std::cin, line);

    std::vector<Ticket> tickets;
    // Read the rest of the tickets
    while (std::getline(std::cin, line)) {
        tickets.push_back(create_ticket(line));
    }

    int ticket_scanning_error_rate = 0;
    for (const auto& ticket : tickets) {
        for (const auto& it : ticket) {
            bool valid = false;
            for (const auto& [_, ranges] : rules) {
                for (const auto& range : ranges) {
                    if (in_range(it, range)) {
                        valid = true;
                        goto outer_loop;
                    }
                }
            }
outer_loop:
            if (!valid) {
                ticket_scanning_error_rate += it;
            }
        }
    }

    std::cout << ticket_scanning_error_rate << std::endl;
}

Range create_range(const std::string& s) {
    auto parts = split(s, '-');
    return Range{std::stoi(parts[0]), std::stoi(parts[1])};
}

std::ostream& operator<<(std::ostream& out, const Range& range) {
    return out << "[" << range.lo << ", " << range.hi << "]";
}

bool in_range(int x, const Range& range) {
    return range.lo <= x && x <= range.hi;
}

Rule create_rule(const std::string& s) {
    auto parts = split(s, ' ');
    std::string raw_range1 = parts[1];
    std::string raw_range2 = parts[3];
    Range range1 = create_range(raw_range1);
    Range range2 = create_range(raw_range2);
    return Rule{range1, range2};
}

std::ostream& operator<<(std::ostream& out, const Rule& rule) {
    if (rule.empty()) {
        return out << "[empty]";
    }

    out << rule.at(0);
    for (size_t i = 1; i < rule.size(); i++) {
        out << " or " << rule.at(i);
    }
    return out;
}

Ticket create_ticket(const std::string& s) {
    Ticket ticket;
    ticket.reserve(19);
    for (const auto& it : split(s, ',')) {
        ticket.push_back(std::stoi(it));
    }
    return ticket;
}

std::ostream& operator<<(std::ostream& out, const Ticket& ticket) {
    out << "[ ";
    for (const auto& it : ticket) {
        out << it << " ";
    }
    return out << "]";
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
