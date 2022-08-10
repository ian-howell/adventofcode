package main

import (
	"fmt"
	"strings"
)

type Rule [2]string

func NewRule(s string) Rule {
	from, to, _ := strings.Cut(s, " => ")
	return Rule{from, to}
}

func (r Rule) From() string {
	return r[0]
}

func (r Rule) To() string {
	return r[1]
}

func (r Rule) String() string {
	return fmt.Sprintf("<%s => %s>", r[0], r[1])
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
