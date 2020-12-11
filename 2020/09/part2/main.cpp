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

void create_fenwick(const std::vector<int>& nums, std::vector<int>& fenwick);
int get_sum(const std::vector<int>& fenwick, int start, int len);
int get_min(const std::vector<int>& nums, int lb, int ub);
int get_max(const std::vector<int>& nums, int lb, int ub);

int get_invalid_num(const std::vector<int>& nums);

int main() {
    std::vector<int> nums;
    get_inputs(nums);

    int key =  get_invalid_num(nums);

    std::vector<int> fenwick;
    create_fenwick(nums, fenwick);
    for (size_t range = 2; range < nums.size(); range++) {
        for (size_t i = 0; i < (nums.size() - range); i++) {
            if (key == get_sum(fenwick, i, range)) {
                int range_start = i;
                int range_end = i + range - 1;
                int smallest = get_min(nums, range_start, range_end);
                int largest = get_max(nums, range_start, range_end);
                std::cout << smallest + largest << std::endl;
            }
        }
    }
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

void create_fenwick(const std::vector<int>& nums, std::vector<int>& fenwick) {
    if (nums.empty()) {
        return;
    }
    fenwick.push_back(nums.at(0));
    for (size_t i = 1; i < nums.size(); i++) {
        fenwick.push_back(fenwick.at(i-1) + nums.at(i));
    }
}

int get_sum(const std::vector<int>& fenwick, int start, int len) {
    if (start == 0) {
        return fenwick.at(start + len - 1);
    }
    return fenwick.at(start + len - 1) - fenwick.at(start - 1);
}

int get_min(const std::vector<int>& nums, int lb, int ub) {
    int min = nums.at(lb);
    for (int i = lb + 1; i <= ub; i++) {
        min = min <= nums.at(i) ? min : nums.at(i);
    }
    return min;
}

int get_max(const std::vector<int>& nums, int lb, int ub) {
    int max = nums.at(lb);
    for (int i = lb + 1; i <= ub; i++) {
        max = max >= nums.at(i) ? max : nums.at(i);
    }
    return max;
}
