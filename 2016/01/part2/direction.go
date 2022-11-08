package main

type DirectionType int

const (
	NoneDir DirectionType = iota
	North
	West
	South
	East
)

func (d DirectionType) String() string {
	switch d {
	case North:
		return "North"
	case West:
		return "West"
	case South:
		return "South"
	case East:
		return "East"
	default:
		return "none"
	}
}

func turn(s string, d DirectionType) DirectionType {
	if s == "L" {
		return turnLeft(d)
	} else if s == "R" {
		return turnRight(d)
	}
	return d
}

func turnLeft(d DirectionType) DirectionType {
	return (d+2)%4 + 1
}

func turnRight(d DirectionType) DirectionType {
	return d%4 + 1
}
