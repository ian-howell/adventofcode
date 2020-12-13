#include <cctype>
#include <iostream>
#include <vector>

std::vector<std::string> split(const std::string& s, char delim);
std::vector<int> filter_ints(const std::vector<std::string>& strings);

int wait_time(int target, int interval);
int get_shortest_wait_time(int target, std::vector<int> departure_times);

int main() {
    int target;
    std::string line;
    std::cin >> target >> line;

    std::vector<int> departure_times = filter_ints(split(line, ','));
    int bus_id = get_shortest_wait_time(target, departure_times);
    std::cout << bus_id * wait_time(target, bus_id) << std::endl;
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

std::vector<int> filter_ints(const std::vector<std::string>& strings) {
    std::vector<int> ints;
    for (const auto& s : strings) {
        if (s.size() > 0 && isdigit(s[0])) {
            ints.push_back(std::stoi(s));
        }
    }
    return ints;
}

int wait_time(int target, int interval) {
    return interval - (target % interval);
}

int get_shortest_wait_time(int target, std::vector<int> departure_times) {
    int best_departure_time = departure_times.at(0);
    int shortest_wait_time = wait_time(target, best_departure_time);
    for (size_t i = 1; i < departure_times.size(); i++) {
        int this_wait_time = wait_time(target, departure_times.at(i));
        if (this_wait_time < shortest_wait_time) {
            shortest_wait_time = this_wait_time;
            best_departure_time = departure_times.at(i);
        }
    }
    return best_departure_time;
}
