#include <iostream>
#include <string>
#include <vector>

void fill_grid_from_stdin(std::vector<std::string>& grid);
void print_grid(const std::vector<std::string>& grid);

int slide(int right, int down, const std::vector<std::string>& grid);

int main() {
    std::vector<std::string> grid;
    fill_grid_from_stdin(grid);
    /* print_grid(grid); */

    int trees_hit = slide(3, 1, grid);
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
        current_row++;
        trees += grid.at(current_row)[current_col] == '#';
    }
    return trees;
}
