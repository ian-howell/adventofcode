#include <iostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
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

std::vector<Ticket> filter_valid_tickets(const std::vector<Ticket>& tickets, const std::unordered_map<std::string, Rule>& rules);
bool is_valid(const Ticket& ticket, const std::unordered_map<std::string, Rule>& rules);
bool matches_at_least_one_rule(int val, const std::unordered_map<std::string, Rule>& rules);
bool matches_rule(int val, const Rule& rule);

std::unordered_set<int> get_possible_fields(const Rule& rule, const std::vector<Ticket>& tickets);

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
    // Get my ticket's info
    std::getline(std::cin, line);
    Ticket my_ticket = create_ticket(line);

    // Eat the empty line
    std::getline(std::cin, line);
    // Eat the "nearby tickets:" line
    std::getline(std::cin, line);

    std::vector<Ticket> tickets;
    // Read the rest of the tickets
    while (std::getline(std::cin, line)) {
        tickets.push_back(create_ticket(line));
    }

    std::vector<Ticket> valid_tickets = filter_valid_tickets(tickets, rules);

    std::unordered_map<std::string, std::unordered_set<int>> possible_field_positions;
    for (const auto& [name, rule] : rules) {
        possible_field_positions[name] = get_possible_fields(rule, valid_tickets);
    }

    std::unordered_map<std::string, int> field_positions;
    while (!possible_field_positions.empty()) {
        std::vector<std::string> to_erase;
        for (const auto& [name, possible_positions] : possible_field_positions) {
            if (possible_positions.size() == 1) {
                // There's only one item possible for "name"
                field_positions[name] = *possible_positions.begin();

                // We should mark this for deletion so we don't check it again later
                to_erase.push_back(name);

                // We also need to remove this field from all other possibilities
                for (auto& [_, remove_from] : possible_field_positions) {
                    remove_from.erase(field_positions[name]);
                }
            }
        }

        // Remove all field positions that were completed this round
        for (const auto& it : to_erase) {
            possible_field_positions.erase(it);
        }
    }

    long departureProduct = 1l;
    for (const auto& [name, field] : field_positions) {
        if (split(name, ' ')[0] == "departure") {
            departureProduct *= my_ticket.at(field);
        }
    }
    std::cout << departureProduct << std::endl;

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

std::vector<Ticket> filter_valid_tickets(const std::vector<Ticket>& tickets, const std::unordered_map<std::string, Rule>& rules) {
    std::vector<Ticket> valid_tickets;
    for (const auto& ticket : tickets) {
        if (is_valid(ticket, rules)) {
            valid_tickets.push_back(ticket);
        }
    }
    return valid_tickets;
}

bool is_valid(const Ticket& ticket, const std::unordered_map<std::string, Rule>& rules) {
    for (const auto& ticket_val : ticket) {
        if (!matches_at_least_one_rule(ticket_val, rules)) {
            return false;
        }
    }
    return true;
}

bool matches_at_least_one_rule(int val, const std::unordered_map<std::string, Rule>& rules) {
    for (const auto& [_, rule] : rules) {
        if (matches_rule(val, rule)) {
            return true;
        }
    }
    return false;
}

bool matches_rule(int val, const Rule& rule) {
    for (const auto& range : rule) {
        if (in_range(val, range)) {
            return true;
        }
    }
    return false;
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

std::unordered_set<int> get_possible_fields(const Rule& rule, const std::vector<Ticket>& tickets) {
    size_t num_fields = tickets.at(0).size();
    std::vector<bool> field_table(num_fields, true);
    for (size_t i = 0; i < num_fields; i++) {
        for (const auto& ticket : tickets) {
            if (!matches_rule(ticket.at(i), rule)) {
                field_table.at(i) = false;
                break;
            }
        }
    }
    std::unordered_set<int> possible_fields;
    for (size_t i = 0; i < num_fields; i++) {
        if (field_table.at(i)) {
            possible_fields.insert(i);
        }
    }
    return possible_fields;
}
