#include <algorithm>
#include <iostream>
#include <map>
#include <vector>

using list = std::vector<int>;
using adjacency_list = std::map<int, list>;

long count_paths(const list& nums);

list get_inputs();
adjacency_list create_graph(const list& nums);
void print_graph(const adjacency_list& graph);

int main() {
    list nums = get_inputs();

    // Add the joltage rating of the outlet
    nums.push_back(0);

    std::sort(nums.begin(), nums.end());

    // Add the joltage rating of the device
    nums.push_back(nums.at(nums.size() - 1) + 3);

    std::cout << count_paths(nums) << std::endl;
}

list get_inputs() {
    list nums;
    int num;
    while (std::cin >> num) {
        nums.push_back(num);
    }
    return nums;
}

adjacency_list create_graph(const list& nums) {
    adjacency_list graph;
    for (size_t i = 0; i < nums.size()-1; i++) {
        graph[nums.at(i)] = list();
        for (size_t j = i+1; (j < nums.size()) && (j < i + 4); j++) {
            int delta = nums.at(j) - nums.at(i);
            if (delta > 3) {
                break;
            }
            graph[nums.at(i)].push_back(nums.at(j));
        }
    }
    return graph;
}

void print_graph(const adjacency_list& graph) {
    for (const auto& it : graph) {
        std::cout << it.first << ": [";
        for (const auto& it2 : it.second) {
            std::cout << " " << it2;
        }
        std::cout << " ]" << std::endl;
    }
}

long count_paths(const list& nums) {
    adjacency_list graph = create_graph(nums);

    std::map<int, long> path_map;
    path_map[nums.at(nums.size()-1)] = 1;

    for (int i = nums.size()-1; i >= 0; i--) {
        int val = nums.at(i);
        for (int j = 1; j <= 3; j++) {
            if (path_map.count(val+j)) {
                path_map[val] += path_map[val+j];
            }
        }
    }

    return path_map[nums.at(0)];
}
