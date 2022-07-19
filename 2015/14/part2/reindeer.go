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

	isRunning     bool
	timeTilSwitch int
	Distance      int
}

func New(s string) *Reindeer {
	parts := strings.Fields(s)
	r := &Reindeer{Name: parts[0], isRunning: true}

	var err error

	r.Speed, err = strconv.Atoi(parts[3])
	if err != nil {
		panic(err)
	}

	r.RunTime, err = strconv.Atoi(parts[6])
	if err != nil {
		panic(err)
	}
	r.timeTilSwitch = r.RunTime

	r.RestTime, err = strconv.Atoi(parts[13])
	if err != nil {
		panic(err)
	}

	return r
}

func (r *Reindeer) Run() {
	if r.isRunning {
		r.Distance += r.Speed
	}

	r.timeTilSwitch--
	if r.timeTilSwitch == 0 {
		if r.isRunning {
			r.timeTilSwitch = r.RestTime
		} else {
			r.timeTilSwitch = r.RunTime
		}
		r.isRunning = !r.isRunning
	}
}

func (r *Reindeer) String() string {
	return fmt.Sprintf("[<Name: %s> <Speed: %d> <RunTime: %d> <RestTime: %d>]",
		r.Name, r.Speed, r.RunTime, r.RestTime)
}
