#include <iostream>
#include <unordered_map>
#include <string>
#include <vector>

std::vector<int> get_initial_nums();
std::vector<std::string> split(const std::string& s, char delim);

int main() {
    auto initial_nums = get_initial_nums();
    std::unordered_map<int, int> last_time_spoken;
    last_time_spoken.reserve(30000000);
    std::vector<bool> seen(30000000, false);

    for (size_t i = 0; i < initial_nums.size() - 1; i++) {
        seen.at(initial_nums.at(i)) = true;
        last_time_spoken[initial_nums.at(i)] = i;
    }

    int last_spoken_num = initial_nums.at(initial_nums.size() - 1);
    for (int current_turn = initial_nums.size(); current_turn < 30000000; current_turn++) {
        if (!seen.at(last_spoken_num)) {
            seen.at(last_spoken_num) = true;
            last_time_spoken[last_spoken_num] = current_turn - 1;
            last_spoken_num = 0;
        } else {
            int diff = current_turn - last_time_spoken.at(last_spoken_num) - 1;
            last_time_spoken[last_spoken_num] = current_turn - 1;
            last_spoken_num = diff;
        }
    }
    std::cout << last_spoken_num << std::endl;

    return 0;
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

std::vector<int> get_initial_nums() {
    std::string line;
    std::cin >> line;
    std::vector<int> nums;
    for (const auto& it : split(line, ',')) {
        nums.push_back(std::stoi(it));
    }
    return nums;
}
