#include <exception>
#include <iostream>
#include <string>
#include <vector>

const int NUM_CYCLES = 6;

using Grid = std::vector<std::vector<char>>;

class ConwayCube {
    private:
        std::vector<Grid> grids;

    public:
        ConwayCube() {}

        ConwayCube(const ConwayCube& other) : grids(other.grids) {}

        ConwayCube(int x_size, int y_size, int z_size) {
            for (int z = 0; z < z_size; z++) {
                Grid grid;
                for (int x = 0; x < x_size; x++) {
                    std::vector<char> row;
                    for (int y = 0; y < y_size; y++) {
                        row.push_back('.');
                    }
                    grid.push_back(row);
                }
                grids.push_back(grid);
            }
        }

        explicit ConwayCube(const Grid& genesis) :
            ConwayCube(2*NUM_CYCLES + genesis.size(), 2*NUM_CYCLES + genesis.at(0).size(), 2*NUM_CYCLES + 1)
        {
            Grid& grid = grids.at(grids.size()/2);
            int start_row = (grid.size() - genesis.size()) / 2;
            int start_col = (grid.at(0).size() - genesis.at(0).size()) / 2;

            for (size_t x = 0; x < genesis.size(); x++) {
                for (size_t y = 0; y < genesis.at(0).size(); y++) {
                    grid.at(start_row + x).at(start_col + y) = genesis.at(x).at(y);
                }
            }
        }

        ConwayCube& operator=(const ConwayCube& other) {
            if (this != &other) {
                grids = other.grids;
            }
            return *this;
        }

        Grid& at(size_t i) {
            return grids.at(i);
        }

        size_t size() {
            return grids.size();
        }

        void print(std::ostream& out) {
            for (const auto& grid : grids) {
                for (const auto& row : grid) {
                    for (const auto& it : row) {
                        out << it;
                    }
                    out << std::endl;
                }
                out << std::endl;
            }
        }

        int count() {
            int num = 0;
            for (const auto& grid : grids) {
                for (const auto& row : grid) {
                    for (const auto& it : row) {
                        if (it == '#') {
                            num++;
                        }
                    }
                }
            }
            return num;
        }
};

ConwayCube game_of_life(ConwayCube& cube);
int get_num_neighbors(int x, int y, int z, ConwayCube& cube);

std::vector<std::vector<char>> read_initial_state();

int main() {
    auto initial_state = read_initial_state();

    ConwayCube conway(initial_state);
    conway = game_of_life(conway);

    std::cout << conway.count() << std::endl;
}


ConwayCube game_of_life(ConwayCube& cube) {
    ConwayCube other(cube);
    ConwayCube& flipper1 = cube;
    ConwayCube& flipper2 = other;

    for (int i = 0; i < NUM_CYCLES; i++) {
        for (size_t z = 0; z < cube.size(); z++) {
            for (size_t x = 0; x < cube.at(0).size(); x++) {
                for (size_t y = 0; y < cube.at(0).at(0).size(); y++) {
                    int num_neighbors = get_num_neighbors(x, y, z, flipper1);
                    if (flipper1.at(z).at(x).at(y) == '#') {
                        if (num_neighbors == 2 || num_neighbors == 3) {
                            flipper2.at(z).at(x).at(y) = '#';
                        } else {
                            flipper2.at(z).at(x).at(y) = '.';
                        }
                    } else if (num_neighbors == 3) {
                        flipper2.at(z).at(x).at(y) = '#';
                    } else {
                        flipper2.at(z).at(x).at(y) = '.';
                    }
                }
            }
        }
        ConwayCube& temp = flipper1;
        flipper1 = flipper2;
        flipper2 = temp;
    }

    return flipper1;
}

int get_num_neighbors(int x, int y, int z, ConwayCube& cube) {
    int num_neighbors = 0;
    for (int i = -1; i <= 1; i++) {
        for (int j = -1; j <= 1; j++) {
            for (int k = -1; k <= 1; k++) {
                if (i == 0 && j == 0 && k == 0) {
                    continue;
                }
                try {
                    if (cube.at(z + i).at(x + j).at(y + k) == '#') {
                        num_neighbors++;
                    }
                } catch (std::out_of_range& _) { }
            }
        }
    }
    return num_neighbors;
}

Grid read_initial_state() {
    Grid state;
    std::string line;
    while (std::cin >> line) {
        std::vector<char> row;
        for (const auto& it : line) {
            row.push_back(it);
        }
        state.push_back(row);
    }
    return state;
}
