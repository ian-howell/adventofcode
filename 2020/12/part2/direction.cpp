#include "direction.h"

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
