package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Ingredient struct {
	Name       string
	Properties [5]int
}

func New(s string) *Ingredient {
	parts := strings.Fields(s)
	g := &Ingredient{Name: strings.TrimRight(parts[0], ":")}

	var err error

	g.Properties[0], err = strconv.Atoi(strings.TrimRight(parts[2], ","))
	if err != nil {
		panic(err)
	}

	g.Properties[1], err = strconv.Atoi(strings.TrimRight(parts[4], ","))
	if err != nil {
		panic(err)
	}

	g.Properties[2], err = strconv.Atoi(strings.TrimRight(parts[6], ","))
	if err != nil {
		panic(err)
	}

	g.Properties[3], err = strconv.Atoi(strings.TrimRight(parts[8], ","))
	if err != nil {
		panic(err)
	}

	g.Properties[4], err = strconv.Atoi(parts[10])
	if err != nil {
		panic(err)
	}

	return g
}

func (g *Ingredient) String() string {
	return fmt.Sprintf("<Name:%s> <Capacity:%d> <Durability:%d> <Flavor:%d> <Texture:%d> <Calories:%d>",
		g.Name, g.Properties[0], g.Properties[1], g.Properties[2], g.Properties[3], g.Properties[4])
}
