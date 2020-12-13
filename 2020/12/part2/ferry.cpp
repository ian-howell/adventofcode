#include "ferry.h"

Ferry::Ferry(Point location, Point waypoint) :
    location(location), waypoint(waypoint) { }

void Ferry::execute_all(const InstructionList& instructions) {
    for (const auto& inst : instructions) {
        this->execute(inst);
    }
}

void Ferry::execute(const Instruction& inst) {
    switch (inst.type) {
        case 'L':
            this->turn_left(inst.arg);
            break;
        case 'R':
            this->turn_right(inst.arg);
            break;
        case 'N':
            this->move_waypoint(NORTH, inst.arg);
            break;
        case 'E':
            this->move_waypoint(EAST, inst.arg);
            break;
        case 'S':
            this->move_waypoint(SOUTH, inst.arg);
            break;
        case 'W':
            this->move_waypoint(WEST, inst.arg);
            break;
        case 'F':
            this->move(inst.arg);
            break;
    }
}

void Ferry::turn_right(int degrees) {
    this->turn(degrees);
}

void Ferry::turn_left(int degrees) {
    this->turn(360 - degrees);
}

void Ferry::turn(int degrees) {
    switch (degrees) {
        case 90:
            this->waypoint = {this->waypoint.y, -this->waypoint.x};
            break;
        case 180:
            this->waypoint = {-this->waypoint.x, -this->waypoint.y};
            break;
        case 270:
            this->waypoint = {-this->waypoint.y, this->waypoint.x};
            break;
        default:
            ;
    }
}

void Ferry::move_waypoint(Direction dir, int dist) {
    switch (dir) {
        case NORTH: this->waypoint.y += dist; break;
        case EAST: this->waypoint.x += dist; break;
        case SOUTH: this->waypoint.y -= dist; break;
        case WEST: this->waypoint.x -= dist; break;
        default: ;
    }
}

void Ferry::move(int steps) {
    this->location.y += steps * this->waypoint.y;
    this->location.x += steps * this->waypoint.x;
}
