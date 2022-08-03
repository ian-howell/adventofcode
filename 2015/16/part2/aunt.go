package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AuntSue struct {
	Children    int
	Cats        int
	Samoyeds    int
	Pomeranians int
	Akitas      int
	Vizslas     int
	Goldfish    int
	Trees       int
	Cars        int
	Perfumes    int
}

func NewAuntSue(input string) *AuntSue {
	auntSue := &AuntSue{
		Children:    -1,
		Cats:        -1,
		Samoyeds:    -1,
		Pomeranians: -1,
		Akitas:      -1,
		Vizslas:     -1,
		Goldfish:    -1,
		Trees:       -1,
		Cars:        -1,
		Perfumes:    -1,
	}
	workingStr, rest, _ := strings.Cut(input, ", ")
	for workingStr != "" {
		compound, amountStr, _ := strings.Cut(workingStr, ": ")
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			panic(err)
		}

		switch compound {
		case "children":
			auntSue.Children = amount
		case "cats":
			auntSue.Cats = amount
		case "samoyeds":
			auntSue.Samoyeds = amount
		case "pomeranians":
			auntSue.Pomeranians = amount
		case "akitas":
			auntSue.Akitas = amount
		case "vizslas":
			auntSue.Vizslas = amount
		case "goldfish":
			auntSue.Goldfish = amount
		case "trees":
			auntSue.Trees = amount
		case "cars":
			auntSue.Cars = amount
		case "perfumes":
			auntSue.Perfumes = amount
		}
		workingStr, rest, _ = strings.Cut(rest, ", ")
	}
	return auntSue
}

func (as *AuntSue) String() string {
	sb := strings.Builder{}
	sb.WriteString("< ")
	if as.Children > -1 {
		sb.WriteString(fmt.Sprintf("children:%d ", as.Children))
	}
	if as.Cats > -1 {
		sb.WriteString(fmt.Sprintf("cats:%d ", as.Cats))
	}
	if as.Samoyeds > -1 {
		sb.WriteString(fmt.Sprintf("samoyeds:%d ", as.Samoyeds))
	}
	if as.Pomeranians > -1 {
		sb.WriteString(fmt.Sprintf("pomeranians:%d ", as.Pomeranians))
	}
	if as.Akitas > -1 {
		sb.WriteString(fmt.Sprintf("akitas:%d ", as.Akitas))
	}
	if as.Vizslas > -1 {
		sb.WriteString(fmt.Sprintf("vizslas:%d ", as.Vizslas))
	}
	if as.Goldfish > -1 {
		sb.WriteString(fmt.Sprintf("goldfish:%d ", as.Goldfish))
	}
	if as.Trees > -1 {
		sb.WriteString(fmt.Sprintf("trees:%d ", as.Trees))
	}
	if as.Cars > -1 {
		sb.WriteString(fmt.Sprintf("cars:%d ", as.Cars))
	}
	if as.Perfumes > -1 {
		sb.WriteString(fmt.Sprintf("perfumes:%d ", as.Perfumes))
	}
	sb.WriteString(">")
	return sb.String()
}
