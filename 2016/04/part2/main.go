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
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		room := NewRoom(s.Text())
		if isReal(room) {
			if decrypt(room.Name, room.SectorID) == "northpoleobjectstorage" {
				// Note: To find the key string above, I simply grepped for "north" on the
				// encrypted output. The prompt doesn't really tell you what to look for
				// otherwise...
				fmt.Println(room.SectorID)
				return
			}
		}
	}
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

func decrypt(s string, id int) string {
	shift := id % 26
	sb := strings.Builder{}

	for i := 0; i < len(s); i++ {
		a := s[i]
		if s[i] != ' ' {
			a = rotate(a, shift)
		}
		sb.WriteByte(a)
	}
	return sb.String()
}

func rotate(a byte, shift int) byte {
	return (((a - 'a') + byte(shift)) % 26) + 'a'
}
