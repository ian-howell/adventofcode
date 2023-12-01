package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total += findCalibrationValue(s.Text())
	}
	fmt.Println(total)

	// fmt.Println(Trie("one"))
	// fmt.Println(Trie("two"))
	// fmt.Println(Trie("three"))
	// fmt.Println(Trie("four"))
	// fmt.Println(Trie("five"))
	// fmt.Println(Trie("six"))
	// fmt.Println(Trie("seven"))
	// fmt.Println(Trie("eight"))
	// fmt.Println(Trie("nine"))
	// fmt.Println()
	// fmt.Println(Trie("1"))
	// fmt.Println(Trie("2"))
	// fmt.Println(Trie("3"))
	// fmt.Println(Trie("4"))
	// fmt.Println(Trie("5"))
	// fmt.Println(Trie("6"))
	// fmt.Println(Trie("7"))
	// fmt.Println(Trie("8"))
	// fmt.Println(Trie("9"))
	// fmt.Println()
	// fmt.Println(Trie("invalid"))

}

func findCalibrationValue(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		if val := Trie(s[i:]); val > 0 {
			total = 10 * val
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if val := Trie(s[i:]); val > 0 {
			total += val
			break
		}
	}

	return total
}

func Trie(s string) int {
	trieStr := []byte(`{"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"o":{"n":{"e":1}},"t":{"w":{"o":2},"h":{"r":{"e":{"e":3}}}},"f":{"o":{"u":{"r":4}},"i":{"v":{"e":5}}},"s":{"i":{"x":6},"e":{"v":{"e":{"n":7}}}},"e":{"i":{"g":{"h":{"t":8}}}},"n":{"i":{"n":{"e":9}}}}`)

	trie := map[string]any{}
	if err := json.Unmarshal(trieStr, &trie); err != nil {
		panic(err)
	}

	currentNode := trie
	for _, c := range s {
		char := string(c)
		if val, ok := currentNode[char]; ok {
			if node, isMap := val.(map[string]any); isMap {
				currentNode = node
			} else {
				return int(val.(float64))
			}
		} else {
			break
		}
	}

	return -1
}
