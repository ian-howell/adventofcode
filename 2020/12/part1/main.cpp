#include <iostream>
#include <string>
#include <vector>

struct Point {
    int x;
    int y;
};

typedef enum {
    NORTH = 0,
    EAST,
    SOUTH,
    WEST,
    NUM_DIRECTIONS,
} Direction;

struct Instruction {
    char type;
    int arg;
};

class Ferry {
    public:
        Direction direction;
        Point point;

        void execute(const Instruction& inst);
        void turn(char type, int degrees);
        void move(Direction dir, int dist);
};

std::istream& operator>>(std::istream& in, Instruction& inst);
std::ostream& operator<<(std::ostream& out, const Instruction& inst);

std::ostream& operator<<(std::ostream& out, const Direction& direction);

std::vector<Instruction> read_instructions();

Ferry execute_instructions(const std::vector<Instruction>& instructions);

int manhattan(const Point& a, const Point& b);

int main() {
    std::vector<Instruction> instructions = read_instructions();
    Ferry end = execute_instructions(instructions);
    std::cout << manhattan(end.point, {0, 0}) << std::endl;
    return 0;
}

int manhattan(const Point& a, const Point& b) {
    return std::abs(a.x - b.x) + std::abs(a.y - b.y);
}

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

std::vector<Instruction> read_instructions() {
    std::vector<Instruction> instructions;
    Instruction inst;
    while(std::cin >> inst) {
        instructions.push_back(inst);
    }
    return instructions;
}

Ferry execute_instructions(const std::vector<Instruction>& instructions) {
    Ferry ferry = {EAST, {0, 0}};
    for (const auto& inst : instructions) {
        ferry.execute(inst);
    }
    return ferry;
}

void Ferry::execute(const Instruction& inst) {
    switch (inst.type) {
        case 'L':
            this->turn(inst.type, 360 - inst.arg);
            break;
        case 'R':
            this->turn(inst.type, inst.arg);
            break;
        case 'N':
            this->move(NORTH, inst.arg);
            break;
        case 'E':
            this->move(EAST, inst.arg);
            break;
        case 'S':
            this->move(SOUTH, inst.arg);
            break;
        case 'W':
            this->move(WEST, inst.arg);
            break;
        case 'F':
            this->move(this->direction, inst.arg);
            break;
    }
}

void Ferry::turn(char type, int degrees) {
    if (degrees == 0) { return; }
    this->direction = Direction((this->direction + degrees/90) % NUM_DIRECTIONS);
}

void Ferry::move(Direction dir, int dist) {
    switch (dir) {
        case NORTH: this->point.y += dist; break;
        case EAST: this->point.x += dist; break;
        case SOUTH: this->point.y -= dist; break;
        case WEST: this->point.x -= dist; break;
        default: ;
    }
}

std::ostream& operator<<(std::ostream& out, const Direction& direction) {
    switch (direction) {
        case NORTH: out << "NORTH"; break;
        case EAST: out << "EAST"; break;
        case SOUTH: out << "SOUTH"; break;
        case WEST: out << "WEST"; break;
        default: out << "UNKNOWN";
    }
    return out;
}
