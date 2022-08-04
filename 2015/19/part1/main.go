package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rules, molecule := getInput()
	mutations := getMutations(rules, molecule)
	fmt.Println(len(mutations))
}

func getMutations(rules []Rule, molecule string) map[string]struct{} {
	mutations := map[string]struct{}{}
	for _, rule := range rules {
		mutationsForRule := getMutationsForRule(rule, molecule)
		for m := range mutationsForRule {
			mutations[m] = struct{}{}
		}
	}
	return mutations
}

func getMutationsForRule(rule Rule, molecule string) map[string]struct{} {
	mutations := map[string]struct{}{}
	parts := strings.Split(molecule, rule.From())
	for i := 1; i < len(parts); i++ {
		first := strings.Join(parts[:i], rule.From())
		second := strings.Join(parts[i:], rule.From())
		full := strings.Join([]string{first, second}, rule.To())
		mutations[full] = struct{}{}
	}
	return mutations
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
