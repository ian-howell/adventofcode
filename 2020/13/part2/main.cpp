#include <cctype>
#include <iostream>
#include <queue>
#include <utility>
#include <vector>

std::vector<std::string> split(const std::string& s, char delim);
std::priority_queue<std::pair<long, long>> create_queue(const std::vector<std::string>& strings);
long chinese_remainder_theorem(std::priority_queue<std::pair<long, long>> nums);

long wait_time(long target, long interval);

int main() {
    int target;
    std::string line;
    std::cin >> target >> line;

    std::priority_queue<std::pair<long, long>> departure_times = create_queue(split(line, ','));
    /* while (!departure_times.empty()) { */
    /*     std::pair<long, long> u = departure_times.top(); */
    /*     std::cout << u.first << " " << u.second << std::endl; */
    /*     departure_times.pop(); */
    /* } */

    std::cout << chinese_remainder_theorem(departure_times) << std::endl;
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

std::priority_queue<std::pair<long, long>> create_queue(const std::vector<std::string>& strings) {
    std::priority_queue<std::pair<long, long>> ints;
    for (size_t i = 0; i < strings.size(); i++) {
        if (strings.at(i).size() > 0 && isdigit(strings.at(i).at(0))) {
            long num = std::stol(strings.at(i));
            ints.push({wait_time(i, num), num});
        }
    }
    return ints;
}

long wait_time(long target, long interval) {
    long v = target % interval;
    if (v == 0) {
        return 0;
    }
    return interval - (target % interval);
}

long chinese_remainder_theorem(std::priority_queue<std::pair<long, long>> nums) {
    long t = nums.top().first;
    long q = nums.top().second;

    while (nums.size() > 1) {
        nums.pop();
        while (t % nums.top().second != nums.top().first) {
            t += q;
        }
        q *= nums.top().second;
    }
    return t;
}
