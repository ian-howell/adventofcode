#include <exception>
#include <iostream>
#include <string>
#include <vector>

const int NUM_CYCLES = 6;

using Grid = std::vector<std::vector<char>>;
using Cube = std::vector<Grid>;

class ConwayCube {
    private:
        std::vector<Cube> cubes;

    public:
        ConwayCube() {}

        ConwayCube(const ConwayCube& other) : cubes(other.cubes) {}

        ConwayCube(int x_size, int y_size, int z_size, int w_size) {
            for (int w = 0; w < w_size; w++) {
                Cube cube;
                for (int z = 0; z < z_size; z++) {
                    Grid grid;
                    for (int x = 0; x < x_size; x++) {
                        std::vector<char> row;
                        for (int y = 0; y < y_size; y++) {
                            row.push_back('.');
                        }
                        grid.push_back(row);
                    }
                    cube.push_back(grid);
                }
                cubes.push_back(cube);
            }
        }

        explicit ConwayCube(const Grid& genesis) :
            ConwayCube(2*NUM_CYCLES + genesis.size(), 2*NUM_CYCLES + genesis.at(0).size(), 2*NUM_CYCLES + 1, 2*NUM_CYCLES + 1)
        {
            Grid& grid = cubes.at(cubes.size()/2).at(cubes.at(0).size()/2);
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
                cubes = other.cubes;
            }
            return *this;
        }

        Cube& at(size_t i) {
            return cubes.at(i);
        }

        size_t size() {
            return cubes.size();
        }

        void print(std::ostream& out) {
            int w = -size()/2 + 1;
            for (const auto& cube : cubes) {
                int z = -size()/2 + 1;
                for (const auto& grid : cube) {
                    std::cout << "z = " << z << ", w = " << w << std::endl;
                    for (const auto& row : grid) {
                        for (const auto& it : row) {
                            out << it;
                        }
                        out << std::endl;
                    }
                    out << std::endl;
                    z++;
                }
                w++;
            }
        }

        int count() {
            int num = 0;
            for (const auto& cube : cubes) {
                for (const auto& grid : cube) {
                    for (const auto& row : grid) {
                        for (const auto& it : row) {
                            if (it == '#') {
                                num++;
                            }
                        }
                    }
                }
            }
            return num;
        }

        int get_num_neighbors(int x, int y, int z, int w) {
            int num_neighbors = 0;
            for (int l = -1; l <= 1; l++) {
                for (int i = -1; i <= 1; i++) {
                    for (int j = -1; j <= 1; j++) {
                        for (int k = -1; k <= 1; k++) {
                            if (i == 0 && j == 0 && k == 0 && l == 0) {
                                continue;
                            }
                            try {
                                if (at(w + l).at(z + i).at(x + j).at(y + k) == '#') {
                                    num_neighbors++;
                                }
                            } catch (std::out_of_range& _) { }
                        }
                    }
                }
            }
            return num_neighbors;
        }
};

ConwayCube game_of_life(ConwayCube& cube);

std::vector<std::vector<char>> read_initial_state();

int main() {
    auto initial_state = read_initial_state();
    ConwayCube conway(initial_state);
    conway = game_of_life(conway);
    std::cout << conway.count() << std::endl;
}

ConwayCube game_of_life(ConwayCube& conway_cube) {
    ConwayCube other(conway_cube);
    ConwayCube& flipper1 = conway_cube;
    ConwayCube& flipper2 = other;

    for (int i = 0; i < NUM_CYCLES; i++) {
        for (size_t w = 0; w < conway_cube.size(); w++) {
            for (size_t z = 0; z < conway_cube.at(0).size(); z++) {
                for (size_t x = 0; x < conway_cube.at(0).at(0).size(); x++) {
                    for (size_t y = 0; y < conway_cube.at(0).at(0).at(0).size(); y++) {
                        int num_neighbors = flipper1.get_num_neighbors(x, y, z, w);
                        if (flipper1.at(w).at(z).at(x).at(y) == '#') {
                            if (num_neighbors == 2 || num_neighbors == 3) {
                                flipper2.at(w).at(z).at(x).at(y) = '#';
                            } else {
                                flipper2.at(w).at(z).at(x).at(y) = '.';
                            }
                        } else if (num_neighbors == 3) {
                            flipper2.at(w).at(z).at(x).at(y) = '#';
                        } else {
                            flipper2.at(w).at(z).at(x).at(y) = '.';
                        }
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
