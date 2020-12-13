#include "instruction.h"

#include <string>

std::istream& operator>>(std::istream& in, Instruction& inst) {
    std::string raw;
    if (in >> raw) {
        inst.type = raw[0];
        inst.arg = std::stoi(raw.substr(1));
    }
    return in;
}

std::ostream& operator<<(std::ostream& out, const Instruction& inst) {
    out << "<" << inst.type << " " << inst.arg << ">";
    return out;
}

InstructionList read_instructions() {
    InstructionList instructions;
    Instruction inst;
    while(std::cin >> inst) {
        instructions.push_back(inst);
    }
    return instructions;
}
