#include <iostream>
#include <string>
#include <utility>
#include <vector>

void fill_grid_from_stdin(std::vector<std::string>& grid);
void print_grid(const std::vector<std::string>& grid);

int slide(int right, int down, const std::vector<std::string>& grid);

int main() {
    std::vector<std::string> grid;
    fill_grid_from_stdin(grid);
    /* print_grid(grid); */

    long trees_hit = 1;
    std::vector<std::pair<int, int>> slopes = {
        std::make_pair(1, 1),
        std::make_pair(3, 1),
        std::make_pair(5, 1),
        std::make_pair(7, 1),
        std::make_pair(1, 2),
    };

    for (const auto& it : slopes) {
        int s = slide(it.first, it.second, grid);
        std::cout << "for slope " << it.first << " " << it.second << ", "
            << s << " trees were hit" << std::endl;
        trees_hit *= s;
        std::cout << "total trees_hit: " << trees_hit << std::endl;
    }

    std::cout << trees_hit << std::endl;
}

void fill_grid_from_stdin(std::vector<std::string>& grid) {
    std::string line;
    while (std::cin >> line) {
        grid.push_back(line);
    }
}

void print_grid(const std::vector<std::string>& grid) {
    for (size_t i = 0; i < grid.size(); i++) {
        std::cout << grid.at(i) << "\n";
    }
    std::cout << std::flush;
}

int slide(int right, int down, const std::vector<std::string>& grid) {
    int current_row = 0;
    int current_col = 0;

    int width = grid.at(0).size();

    int trees = 0;
    while (current_row < int(grid.size()) - 1) {
        current_col = (current_col + right) % width;
        current_row += down;
        trees += grid.at(current_row)[current_col] == '#';
    }
    return trees;
}
