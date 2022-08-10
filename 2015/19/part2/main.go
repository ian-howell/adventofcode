package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	rules, molecule := getInput()
	// separate(molecule)
	fmt.Println(countSteps(rules, molecule))
}

func countSteps(rules []Rule, molecule string) int {
	// Disclaimer, this algorithm comes from here:
	// https://www.reddit.com/r/adventofcode/comments/3xflz8/comment/cy4etju/
	numRn := strings.Count(molecule, "Rn")
	numAr := strings.Count(molecule, "Ar")
	numY := strings.Count(molecule, "Y")

	// Note the + 1 here. This is added because of the leading "C" in my
	// input. Since "C" is not part of the LHS of any production rule, we
	// need to make sure to count it here in the total number of tokens.
	numTokens := numRn + numAr + numY + 1
	seen := map[string]struct{}{}
	for _, r := range rules {
		if _, ok := seen[r.From()]; !ok {
			x := strings.Count(molecule, r.From())
			// fmt.Printf("There are %d instances of %v\n", x, r.From())
			numTokens += x
			seen[r.From()] = struct{}{}
		}
	}
	// fmt.Println()

	// fmt.Printf("There are %d instances of %v\n", numRn, "Rn")
	// fmt.Printf("There are %d instances of %v\n", numAr, "Ar")
	// fmt.Printf("There are %d instances of %v\n", numY, "Y")
	// fmt.Println()

	// fmt.Printf("There are a total of %d tokens in the molecule\n", numTokens)

	return (numTokens - 1) - (numRn + numAr) - (2 * numY)
}

func getInput() ([]Rule, string) {
	var molecule string
	rules := []Rule{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if strings.Contains(s.Text(), "=>") {
			rules = append(rules, NewRule(s.Text()))
		} else if len(s.Text()) > 0 {
			molecule = s.Text()
		}
	}
	return rules, molecule
}

func separate(molecule string) {
	first := true
	total := 0
	counts := map[string]int{}
	sb := strings.Builder{}
	for _, c := range molecule {
		if !first && unicode.IsUpper(c) {
			total++
			counts[sb.String()]++
			fmt.Printf("% 4d (% 3d): %v\n", total, counts[sb.String()], sb.String())
			sb.Reset()
			sb.WriteRune(c)
		} else {
			sb.WriteRune(c)
		}
		first = false
	}
	if sb.String() != "" {
		total++
		counts[sb.String()]++
		fmt.Printf("% 4d (% 3d): %v\n", total, counts[sb.String()], sb.String())
	}
}
