#include <iostream>
#include <string>
#include <vector>

using Grid = std::vector<std::vector<char>>;

Grid read_grid();
void print_grid(const Grid& grid);
bool step(Grid& grid);
int count_neighbors(const Grid& grid, int row, int col);
bool is_valid(const Grid& grid, int row, int col);
int count_occupied_seats(const Grid& grid);

int main() {
    Grid grid = read_grid();
    while(step(grid));
    std::cout << count_occupied_seats(grid) << std::endl;
    return 0;
}

Grid read_grid() {
    Grid grid;
    std::string line;
    while (std::cin >> line) {
        std::vector<char> row;
        for (const auto& it : line) {
            row.push_back(it);
        }
        grid.push_back(row);
    }
    return grid;
}

void print_grid(const Grid& grid) {
    for (const auto& row : grid) {
        for (const auto& it : row) {
            std::cout << it;
        }
        std::cout << std::endl;
    }
}

bool step(Grid& grid) {
    Grid old = grid;
    bool changed = false;
    for (size_t r = 0; r < old.size(); r++) {
        for (size_t c = 0; c < old.at(0).size(); c++) {
            int num_neighbors = count_neighbors(old, r, c);
            if (old.at(r).at(c) == 'L') {
                if (num_neighbors == 0) {
                    grid.at(r).at(c) = '#';
                    changed = true;
                }
            } else if (old.at(r).at(c) == '#') {
                if (num_neighbors >= 4) {
                    grid.at(r).at(c) = 'L';
                    changed = true;
                }
            }
        }
    }
    return changed;
}

int count_neighbors(const Grid& grid, int row, int col) {
    int num_neighbors = 0;
    std::vector<std::vector<int>> m = {
        {+0, +1}, {+0, -1}, {+1, +0}, {-1, +0},
        {+1, +1}, {+1, -1}, {-1, -1}, {-1, +1},
    };

    for (const auto& it : m) {
        int new_row = row + it.at(0);
        int new_col = col + it.at(1);
        if (is_valid(grid, new_row, new_col) && grid.at(new_row).at(new_col) == '#') {
            num_neighbors++;
        }
    }
    return num_neighbors;
}

bool is_valid(const Grid& grid, int row, int col) {
    return row >= 0 && row < int(grid.size()) && col >= 0 && col < int(grid.at(0).size());
}

int count_occupied_seats(const Grid& grid) {
    int num_occupied_seats = 0;
    for (const auto& row : grid) {
        for (const auto& it : row) {
            if (it == '#') {
                num_occupied_seats++;
            }

        }
    }
    return num_occupied_seats;
}
