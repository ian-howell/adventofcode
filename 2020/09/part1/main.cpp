#include <exception>
#include <iostream>
#include <map>
#include <vector>

#define PREAMBLE_LEN 25

void get_inputs(std::vector<int>& nums);
void add_sums(const std::vector<int>& nums, std::map<int, int>& sums, int index, int start, int len);
void remove_sums(const std::vector<int>& nums, std::map<int, int>& sums, int index, int start, int len);
bool check_valid(const std::map<int, int>& sums, int target);
void print_map(const std::map<int, int>& m);

int get_invalid_num(const std::vector<int>& nums);

int main() {
    std::vector<int> nums;
    get_inputs(nums);

    std::cout << get_invalid_num(nums) << std::endl;
    return 0;
}

void get_inputs(std::vector<int>& nums) {
    int num;
    while (std::cin >> num) {
        nums.push_back(num);
    }
}

void add_sums(const std::vector<int>& nums, std::map<int, int>& sums, int index, int start, int len) {
    for (int i = start; i < (start + len - 1); i++) {
        if (i != index) {
            int sum = nums.at(index) + nums.at(i);
            sums[sum]++;
        }
    }
}

void remove_sums(const std::vector<int>& nums, std::map<int, int>& sums, int index, int start, int len) {
    for (int i = start; i < (start + len - 1); i++) {
        if (i != index) {
            int sum = nums.at(index) + nums.at(i);
            sums[sum]--;
            if (sums[sum] == 0) {
                sums.erase(sum);
            }
        }
    }
}

bool check_valid(const std::map<int, int>& sums, int target) {
    return sums.count(target) > 0;
}

void print_map(const std::map<int, int>& m) {
    for (const auto& it : m) {
        std::cout << "<" << it.first << ", " << it.second << ">" << std::endl;
    }
}

int get_invalid_num(const std::vector<int>& nums) {
    std::map<int, int> sums;
    // Add the initial set of valid sums
    for (int i = 0; i < (PREAMBLE_LEN - 1); i++) {
        add_sums(nums, sums, i, i + 1, PREAMBLE_LEN - i);
    }

    for (size_t i = PREAMBLE_LEN; i < nums.size(); i++) {
        if (!check_valid(sums, nums.at(i))) {
            return nums.at(i);
        }

        int to_add = i;
        int to_remove = i - PREAMBLE_LEN;
        int lb = to_remove + 1;
        remove_sums(nums, sums, to_remove, lb, PREAMBLE_LEN);
        add_sums(nums, sums, to_add, lb, PREAMBLE_LEN);
    }

    throw std::runtime_error("Could not find an invalid number");
}
