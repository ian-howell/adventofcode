#include "ferry.h"

#include <iostream>

int main() {
    InstructionList instructions = read_instructions();
    Ferry ferry({0, 0}, {10, 1});
    ferry.execute_all(instructions);
    std::cout << manhattan(ferry.get_location(), {0, 0}) << std::endl;
    return 0;
}
