#include <exception>
#include <iostream>
#include <set>
#include <string>
#include <sstream>
#include <vector>

typedef enum {
    UNKNOWN,
    NOP,
    ACC,
    JMP,
} OpCode;

std::istream& operator>>(std::istream& in, OpCode& op_code);
std::ostream& operator<<(std::ostream& out, const OpCode& op_code);

struct Instruction {
    OpCode op_code;
    int argument;
};

std::istream& operator>>(std::istream& in, Instruction& instruction);
std::ostream& operator<<(std::ostream& out, const Instruction& instruction);

void read_program(std::vector<Instruction>& program);
void print_program(const std::vector<Instruction>& program);
int execute_program(const std::vector<Instruction>& program);
int execute_program_from(const std::vector<Instruction>& program, const std::set<int>& previsited, size_t program_counter);

int main() {
    std::vector<Instruction> program;
    read_program(program);
    /* print_program(program); */
    int return_value = execute_program(program);
    std::cout << return_value << std::endl;
    return 0;
}

void read_program(std::vector<Instruction>& program) {
    Instruction instruction;
    while (std::cin >> instruction) {
        program.push_back(instruction);
    }
}

void print_program(const std::vector<Instruction>& program) {
    for (auto instruction : program) {
        std::cout << instruction << std::endl;
    }
}

int execute_program(const std::vector<Instruction>& program) {
    int accumulator = 0;
    size_t program_counter = 0;
    std::set<int> visited;

    while (program_counter < program.size()) {
        Instruction instruction = program.at(program_counter);
        visited.insert(program_counter);
        switch (instruction.op_code) {
            case NOP:
                try {
                    // Change this to a jmp
                    return accumulator + execute_program_from(program, visited, program_counter + instruction.argument);
                } catch (std::runtime_error& e) { }
                program_counter++;
                break;
            case ACC:
                accumulator += instruction.argument;
                program_counter++;
                break;
            case JMP:
                try {
                    // Change this to a nop
                    return accumulator + execute_program_from(program, visited, program_counter + 1);
                } catch (std::runtime_error& e) { }
                program_counter += instruction.argument;
                break;
            default:
                std::ostringstream err;
                err << "Unknown opcode on line " << program_counter << ": " << instruction.op_code;
                throw std::runtime_error(err.str());
        }
    }
    return accumulator;
}

int execute_program_from(const std::vector<Instruction>& program, const std::set<int>& previsited, size_t program_counter) {
    int accumulator = 0;
    std::set<int> visited;

    while (program_counter < program.size()) {
        Instruction instruction = program.at(program_counter);
        if ((visited.count(program_counter) != 0) || (previsited.count(program_counter) != 0)) {
            // Infinite loop detected
            std::ostringstream err;
            err << "Revisiting line " << program_counter << ": "
                << instruction << ", resulting in an infinite loop";
            throw std::runtime_error(err.str());
        }
        visited.insert(program_counter);
        switch (instruction.op_code) {
            case NOP:
                program_counter++;
                break;
            case ACC:
                accumulator += instruction.argument;
                program_counter++;
                break;
            case JMP:
                program_counter += instruction.argument;
                break;
            default:
                std::ostringstream err;
                err << "Unknown opcode on line " << program_counter << ": " << instruction.op_code;
                throw std::runtime_error(err.str());
        }
    }
    return accumulator;
}

std::istream& operator>>(std::istream& in, OpCode& op_code) {
    std::string raw_op_code;
    in >> raw_op_code;
    if (raw_op_code == "nop") {
        op_code = NOP;
    } else if (raw_op_code == "acc") {
        op_code = ACC;
    } else if (raw_op_code == "jmp") {
        op_code = JMP;
    } else {
        op_code = UNKNOWN;
    }
    return in;
}

std::ostream& operator<<(std::ostream& out, const OpCode& op_code) {
    switch (op_code) {
        case NOP: out << "nop"; break;
        case ACC: out << "acc"; break;
        case JMP: out << "jmp"; break;
        default: out << "unknown"; break;
    }
    return out;
}

std::istream& operator>>(std::istream& in, Instruction& instruction) {
    in >> instruction.op_code >> instruction.argument;
    return in;
}

std::ostream& operator<<(std::ostream& out, const Instruction& instruction) {
    out << "<" << instruction.op_code << " " << instruction.argument << ">";
    return out;
}
