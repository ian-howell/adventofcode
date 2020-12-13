#pragma once

#include <iostream>
#include <vector>

struct Instruction {
    char type;
    int arg;
};

using InstructionList = std::vector<Instruction>;

std::istream& operator>>(std::istream& in, Instruction& inst);
std::ostream& operator<<(std::ostream& out, const Instruction& inst);

std::vector<Instruction> read_instructions();
