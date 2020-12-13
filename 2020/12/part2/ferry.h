#pragma once

#include "direction.h"
#include "instruction.h"
#include "point.h"

class Ferry {
    public:
        Ferry() {}
        Ferry(Point location, Point waypoint);

        Point get_location() { return location; }

        void execute_all(const InstructionList& instructions);
        void execute(const Instruction& inst);

    private:
        Point location;
        Point waypoint;

        void turn(int degrees);
        void turn_right(int degrees);
        void turn_left(int degrees);
        void move_waypoint(Direction dir, int dist);
        void move(int steps);
};
