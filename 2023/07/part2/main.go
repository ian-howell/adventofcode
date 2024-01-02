package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var camelCardHands CamelCardHandSlice
	for s.Scan() {
		cardPart, bidPart, _ := strings.Cut(s.Text(), " ")
		camelCardHands = append(camelCardHands,
			CamelCardHand{
				camelCard: CamelCard(cardPart),
				bid:       atoi(bidPart),
			})
	}

	totalWinnings := 0
	sort.Sort(camelCardHands)
	for i, camelCardHand := range camelCardHands {
		totalWinnings += (1 + i) * camelCardHand.bid
		fmt.Println(camelCardHand, camelCardHand.camelCard.HandType())
	}
	fmt.Println(totalWinnings)
}

type CamelCardHand struct {
	camelCard CamelCard
	bid       int
}

type CamelCardHandSlice []CamelCardHand

func (c CamelCardHandSlice) Len() int           { return len(c) }
func (c CamelCardHandSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CamelCardHandSlice) Less(i, j int) bool { return c[i].camelCard.Less(c[j].camelCard) }

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
