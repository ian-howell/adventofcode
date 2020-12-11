#include <iostream>
#include <map>
#include <string>
#include <utility>
#include <vector>

using Grid = std::vector<std::vector<char>>;
using NeighborGrid = std::vector<std::vector<std::vector<std::pair<int, int>>>>;

Grid read_grid();
void print_grid(const Grid& grid);
bool step(Grid& grid, const NeighborGrid& neighbor_grid);
int count_neighbors(const Grid& grid, const NeighborGrid& neighbor_grid, int row, int col);
bool is_valid(const Grid& grid, int row, int col);
int count_occupied_seats(const Grid& grid);
NeighborGrid create_neighbor_grid(const Grid& grid);
std::vector<std::pair<int, int>> get_neighbors(const Grid& grid, int row, int col);
void print_neighbor_grid(const NeighborGrid& neighbor_grid);

int main() {
    Grid grid = read_grid();
    NeighborGrid neighbor_grid = create_neighbor_grid(grid);

    while(step(grid, neighbor_grid));
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

bool step(Grid& grid, const NeighborGrid& neighbor_grid) {
    Grid old = grid;
    bool changed = false;
    for (size_t r = 0; r < old.size(); r++) {
        for (size_t c = 0; c < old.at(0).size(); c++) {
            int num_neighbors = count_neighbors(old, neighbor_grid, r, c);
            if (old.at(r).at(c) == 'L') {
                if (num_neighbors == 0) {
                    grid.at(r).at(c) = '#';
                    changed = true;
                }
            } else if (old.at(r).at(c) == '#') {
                if (num_neighbors >= 5) {
                    grid.at(r).at(c) = 'L';
                    changed = true;
                }
            }
        }
    }
    return changed;
}

int count_neighbors(const Grid& grid, const NeighborGrid& neighbor_grid, int row, int col) {
    int num_neighbors = 0;
    for (const auto& it : neighbor_grid.at(row).at(col)) {
        int new_row = it.first;
        int new_col = it.second;
        if (grid.at(new_row).at(new_col) == '#') {
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

NeighborGrid create_neighbor_grid(const Grid& grid) {
    NeighborGrid neighbor_grid;
    for (size_t r = 0; r < grid.size(); r++) {
        neighbor_grid.push_back({});
        for (size_t c = 0; c < grid.at(0).size(); c++) {
            if (grid.at(r).at(c) == 'L') {
                neighbor_grid.at(r).push_back(get_neighbors(grid, r, c));
            } else {
                neighbor_grid.at(r).push_back({});
            }
        }
    }
    return neighbor_grid;
}

std::vector<std::pair<int, int>> get_neighbors(const Grid& grid, int row, int col) {
    std::vector<std::pair<int, int>> neighbors;
    std::vector<std::vector<int>> m = {
        {+0, +1}, {+0, -1}, {+1, +0}, {-1, +0},
        {+1, +1}, {+1, -1}, {-1, -1}, {-1, +1},
    };

    for (const auto& it : m) {
        int new_row = row + it.at(0);
        int new_col = col + it.at(1);
        while (is_valid(grid, new_row, new_col)) {
            if (grid.at(new_row).at(new_col) == 'L') {
                neighbors.push_back({new_row, new_col});
                break;
            }
            new_row += it.at(0);
            new_col += it.at(1);
        }
    }
    return neighbors;
}

void print_neighbor_grid(const NeighborGrid& neighbor_grid) {
    for (size_t r = 0; r < neighbor_grid.size(); r++) {
        for (size_t c = 0; c < neighbor_grid.at(0).size(); c++) {
            std::cout << "NEIGHBORS OF (" << r << ", " << c << "): [";
            for (const auto& it : neighbor_grid.at(r).at(c)) {
                std::cout << " (" << it.first << ", " << it.second << ")";
            }
            std::cout << " ]" << std::endl;
        }
    }
}
