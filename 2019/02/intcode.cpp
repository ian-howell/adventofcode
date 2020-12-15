#include "intcode.h"
#include "utility.h"

#include <exception>
#include <iomanip>
#include <sstream>

IntCode::IntCode() : pc(0) {}

IntCode::IntCode(const std::string& program) : IntCode() {
    set_program(program);
}

void IntCode::set_program(const std::string& program) {
    memory.clear();
    for (const auto& word : split(program, ',')) {
        memory.push_back(std::stoi(word));
    }
    pc = 0;
}

int IntCode::execute() {
    while (!is_finished()) {
        step();
    }
    return memory.at(0);
}

void IntCode::step() {
    switch (memory.at(pc)) {
        case OpCode::ADD:
            add();
            break;
        case OpCode::MULTIPLY:
            multiply();
            break;
        case OpCode::FINISH:
            finish();
            break;
        default:
            std::ostringstream err;
            err << "Unknown opcode on line " << pc << ": " << memory.at(pc);
            throw std::runtime_error(err.str());
    }
}

bool IntCode::is_finished() {
    return pc == -1;
}

void IntCode::set_memory(int index, int value) {
    memory.at(index) = value;
}

int IntCode::get_memory(int index) {
    return memory.at(index);
}

void IntCode::dump_memory(std::ostream& out) {
    out << std::right << std::setw(15) << "\\";
    for (int i = 0; i < 10; i++) {
        out << std::right << std::setw(13) << i;
    }
    out << std::endl;


    out << std::right << std::setw(15) << "-------";
    for (int i = 0; i < 10; i++) {
        out << "-------------";
    }
    out << std::endl;

    for (size_t i = 0; i < memory.size(); i += 10) {
        out << std::right << std::setw(13) << i << " |";
        for (size_t j = i; j < memory.size() && j < i + 10; j++) {
            out << std::right << std::setw(13) << memory.at(j);
        }
        out << std::endl;
    }
}

void IntCode::add() {
    pc++;
    int addend1 = memory.at(memory.at(pc));
    pc++;
    int addend2 = memory.at(memory.at(pc));
    pc++;
    int store_to = memory.at(pc);
    memory.at(store_to) = addend1 + addend2;
    pc++;
}

void IntCode::multiply() {
    pc++;
    int multiplicand1 = memory.at(memory.at(pc));
    pc++;
    int multiplicand2 = memory.at(memory.at(pc));
    pc++;
    int store_to = memory.at(pc);
    memory.at(store_to) = multiplicand1 * multiplicand2;
    pc++;
}

void IntCode::finish() {
    pc = -1;
}
