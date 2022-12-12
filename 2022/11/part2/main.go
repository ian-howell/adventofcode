package main

import (
	"fmt"
	"sort"
	"strconv"
)

const DEBUG = false

func main() {
	monkeys := getMonkeys()
	worryMod := 1
	for _, monkey := range monkeys {
		worryMod *= monkey.Test
	}

	inspections := make([]int, len(monkeys))
	for round := 0; round < 10_000; round++ {
		for i := range monkeys {
			DebugF("Monkey %d:\n", i)
			for monkeys[i].HasItems() {
				inspections[i]++
				val := monkeys[i].Pop()
				DebugF("  Monkey inspects an item with a worry level of %d.\n", val)
				val = monkeys[i].Operation(val)
				val %= worryMod
				DebugF("    Worry level is changed to %d.\n", val)
				DebugF("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", val)
				if val%monkeys[i].Test == 0 {
					DebugF("    Current worry level is divisible by %d.\n", monkeys[i].Test)
					monkeys[monkeys[i].True].Push(val)
					DebugF("    Item with worry level %d is thrown to Monkey %d.\n", val, monkeys[i].True)
				} else {
					DebugF("    Current worry level is not divisible by %d.\n", monkeys[i].Test)
					monkeys[monkeys[i].False].Push(val)
					DebugF("    Item with worry level %d is thrown to Monkey %d.\n", val, monkeys[i].False)
				}
			}
		}
	}

	sort.Ints(inspections)
	l := len(inspections)
	fmt.Println(inspections[l-1] * inspections[l-2])
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func DebugF(format string, args ...any) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}
