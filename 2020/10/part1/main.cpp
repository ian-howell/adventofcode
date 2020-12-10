#include <algorithm>
#include <iostream>
#include <map>
#include <vector>

void get_inputs(std::vector<int>& nums);

int main() {
    std::vector<int> nums;
    get_inputs(nums);

    // Add the joltage rating of the outlet
    nums.push_back(0);

    std::sort(nums.begin(), nums.end());

    // Add the joltage rating of the device
    nums.push_back(nums.at(nums.size() - 1) + 3);

    std::map<int, int> deltas;

    for (size_t i = 1; i < nums.size(); i++) {
        int delta = nums.at(i) - nums.at(i - 1);
        deltas[delta]++;
    }

    if (deltas.size() != 2) {
        std::cerr << "Got some invalid joltages" << std::endl;
    }

    /* for (const auto& it : deltas) { */
    /*     std::cout << it.first << ": " << it.second << std::endl; */
    /* } */

    std::cout << deltas[1] * deltas[3] << std::endl;
}

void get_inputs(std::vector<int>& nums) {
    int num;
    while (std::cin >> num) {
        nums.push_back(num);
    }
}
