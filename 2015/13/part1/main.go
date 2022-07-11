package main

import "fmt"

func main() {
	graph := readGraph()

	guestList := []string{}
	for guest := range graph {
		guestList = append(guestList, guest)
	}

	best := evaluateSeatingChart(guestList, graph)
	allGuestLists := generateAllLists(guestList)

	for _, list := range allGuestLists {
		current := evaluateSeatingChart(list, graph)
		if current > best {
			best = current
		}
	}
	fmt.Println(best)
}

func evaluateSeatingChart(guestList []string, graph Graph) int {
	l := len(guestList)
	total := graph[guestList[0]][guestList[l-1]] + graph[guestList[l-1]][guestList[0]]
	for i := 1; i < l; i++ {
		total += graph[guestList[i]][guestList[i-1]] + graph[guestList[i-1]][guestList[i]]
	}
	return total
}

func generateAllLists(ss []string) [][]string {
	if len(ss) == 1 {
		return [][]string{{ss[0]}}
	}

	result := [][]string{}
	for _, s := range ss {
		for _, subList := range generateAllLists(listSubtract(ss, s)) {
			result = append(result, append([]string{s}, subList...))
		}
	}
	return result
}

func listSubtract(ss []string, s string) []string {
	newList := []string{}
	for _, v := range ss {
		if v != s {
			newList = append(newList, v)
		}
	}
	return newList
}
