package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Reindeer struct {
	Name     string
	Speed    int
	RunTime  int
	RestTime int
}

func New(s string) *Reindeer {
	parts := strings.Fields(s)
	r := &Reindeer{Name: parts[0]}

	var err error

	r.Speed, err = strconv.Atoi(parts[3])
	if err != nil {
		panic(err)
	}

	r.RunTime, err = strconv.Atoi(parts[6])
	if err != nil {
		panic(err)
	}

	r.RestTime, err = strconv.Atoi(parts[13])
	if err != nil {
		panic(err)
	}

	return r
}

func (r *Reindeer) Run(s int) int {
	fullCycles := s / (r.RunTime + r.RestTime)
	partialCycle := s % (r.RunTime + r.RestTime)

	distanceFromFullCycles := fullCycles * (r.Speed * r.RunTime)

	timeRunningInFinalCycle := partialCycle
	if timeRunningInFinalCycle > r.RunTime {
		timeRunningInFinalCycle = r.RunTime
	}
	distanceFromFinalCycle := timeRunningInFinalCycle * r.Speed

	return distanceFromFullCycles + distanceFromFinalCycle
}

func (r *Reindeer) String() string {
	return fmt.Sprintf("[<Name: %s> <Speed: %d> <RunTime: %d> <RestTime: %d>]",
		r.Name, r.Speed, r.RunTime, r.RestTime)
}
