package main

func atoi(a byte) int {
	return int(a & 0x0f)
}
