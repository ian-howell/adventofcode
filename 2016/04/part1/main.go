package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	Name     string
	SectorID int
	CheckSum string
}

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		room := NewRoom(s.Text())
		if isReal(room) {
			total += room.SectorID
		}
	}
	fmt.Println(total)
}

func NewRoom(s string) Room {
	parts := strings.Split(s, "[")
	sectorID, _ := strconv.Atoi(parts[0][len(parts[0])-3:])
	return Room{
		Name:     strings.ReplaceAll(parts[0][:len(parts[0])-3], "-", ""),
		SectorID: sectorID,
		CheckSum: strings.TrimRight(parts[1], "]"),
	}
}

func isReal(room Room) bool {
	freq := map[rune]int{}
	for _, letter := range room.Name {
		freq[letter]++
	}

	largest := 0
	inverseFreq := map[int]string{}
	for letter := range freq {
		if freq[letter] > largest {
			largest = freq[letter]
		}
		inverseFreq[freq[letter]] += string(letter)
	}

	sb := strings.Builder{}

	for ; largest > 0; largest-- {
		sb.WriteString(alphabetize(inverseFreq[largest]))
	}

	shouldBe := alphabetize(sb.String()[:5])
	sortedCheckSum := alphabetize(room.CheckSum)

	return shouldBe == sortedCheckSum
}

func alphabetize(s string) string {
	r := []rune(s)
	sort.Sort(RuneSlice(r))
	return string(r)
}

type RuneSlice []rune

func (s RuneSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s RuneSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s RuneSlice) Len() int {
	return len(s)
}
