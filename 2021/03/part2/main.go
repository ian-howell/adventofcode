package main

import "fmt"

// Use for small input
// const wordLength = 5
const wordLength = 12

func main() {
	origList := []int{}
	var temp int
	for _, err := fmt.Scanf("%b", &temp); err == nil; _, err = fmt.Scanf("%b", &temp) {
		origList = append(origList, temp)
	}

	oxygenList := make([]int, len(origList))
	copy(oxygenList, origList)
	co2List := make([]int, len(origList))
	copy(co2List, origList)

	for i := wordLength - 1; i >= 0; i-- {
		if len(oxygenList) > 1 {
			dominantBit := dominatingBit(i, oxygenList)
			mask := dominantBit << i
			currentOxygenList := []int{}
			for _, val := range oxygenList {
				masked := val & (1 << i)
				if masked^mask == 0 {
					currentOxygenList = append(currentOxygenList, val)
				}
			}
			oxygenList = currentOxygenList
		}

		if len(co2List) > 1 {
			dominantBit := dominatingBit(i, co2List)
			mask := dominantBit << i
			currentCO2List := []int{}
			for _, val := range co2List {
				masked := val & (1 << i)
				if masked^mask != 0 {
					currentCO2List = append(currentCO2List, val)
				}
			}
			co2List = currentCO2List
		}

		if len(co2List) == 1 && len(oxygenList) == 1 {
			break
		}

	}

	oxygenGeneratorRating := oxygenList[0]
	co2ScrubberRating := co2List[0]

	fmt.Println(oxygenGeneratorRating * co2ScrubberRating)
}

func dominatingBit(pos int, nums []int) int {
	ones := 0
	mask := 1 << pos
	for _, val := range nums {
		if val&mask != 0 {
			ones++
		}
	}
	if ones*2 >= len(nums) {
		return 1
	}
	return 0
}
