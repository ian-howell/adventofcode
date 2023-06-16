#include "intcode.h"

#include <iostream>

int main() {
    std::string program;
    std::cin >> program;
    IntCode ic;
    for (int i = 0; i < 100; i++) {
        for (int j = 0; j < 100; j++) {
            ic.set_program(program);
            ic.set_memory(1, i);
            ic.set_memory(2, j);
            int result = ic.execute();
            if (result == 19690720) {
                std::cout << (100*i + j) << std::endl;
                return 0;
            }
        }
    }
}
