package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

type SeqMap struct {
	FirstTriplet rune
	Quintuplets  map[rune]struct{}
}

func main() {
	var input []byte
	fmt.Scan(&input)
	seqMapTable := initializeSeqMapTable(input)
	foundKeys := 0
	for i := 0; ; i++ {
		if isKey(i, seqMapTable) {
			// fmt.Printf("Found a key: %d\n", i)
			foundKeys++
		}

		a := strconv.Itoa(i + 1000)
		testString := append(input, a...)
		hash := stretch(md5.Sum(testString))

		seqMapTable = append(seqMapTable, toSeqMap(hash))

		// if i == 39 {
		// 	printHash(hash)
		// 	fmt.Println(seqMapTable[816])
		// 	break
		// }

		if foundKeys == 64 {
			fmt.Println(i)
			break
		}
	}
}

func isKey(i int, seqMapTable []SeqMap) bool {
	if seqMapTable[i].FirstTriplet == 0 {
		// There were no sequences with length >= 3
		return false
	}

	for j := i + 1; j < i+1000; j++ {
		if _, ok := seqMapTable[j].Quintuplets[seqMapTable[i].FirstTriplet]; ok {
			return true
		}
	}

	return false
}

func initializeSeqMapTable(input []byte) []SeqMap {
	seqMapTable := []SeqMap{}
	for i := 0; i < 1000; i++ {
		a := strconv.Itoa(i)
		testString := append(input, a...)
		hash := stretch(md5.Sum(testString))
		seqMapTable = append(seqMapTable, toSeqMap(hash))
	}
	return seqMapTable
}

func printHash(h [16]byte) {
	fmt.Println(toString(h))
}

func toString(h [16]byte) string {
	sb := strings.Builder{}
	for _, v := range h {
		sb.WriteString(fmt.Sprintf("%02x", v))
	}
	return sb.String()
}

func toSeqMap(h [16]byte) SeqMap {
	seqMap := SeqMap{
		FirstTriplet: 0,
		Quintuplets:  map[rune]struct{}{},
	}
	var last rune
	length := 1
	for _, c := range toString(h) {
		if c == last {
			length++
			if length >= 3 && seqMap.FirstTriplet == 0 {
				seqMap.FirstTriplet = c
			}
			if length >= 5 {
				seqMap.Quintuplets[c] = struct{}{}
			}
		} else {
			length = 1
			last = c
		}
	}

	return seqMap
}

func stretch(hash [16]byte) [16]byte {
	for i := 0; i < 2016; i++ {
		hash = md5.Sum([]byte(toString(hash)))
	}
	return hash
}
