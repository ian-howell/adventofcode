#pragma once

#include <iostream>

typedef enum {
    NORTH = 0,
    EAST,
    SOUTH,
    WEST,
    NUM_DIRECTIONS,
} Direction;

std::ostream& operator<<(std::ostream& out, const Direction& direction);
