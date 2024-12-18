package main

import (
	"slices"
)

type Item struct {
	Priority int
	State    State
}

type PriorityQueue struct {
	items []Item
}

func (p *PriorityQueue) Push(item Item) {
	p.items = append(p.items, item)
}

func (p *PriorityQueue) Pop() Item {
	indexOfLowest := 0
	for i := 1; i < len(p.items); i++ {
		if p.items[i].Priority < p.items[indexOfLowest].Priority {
			indexOfLowest = i
		}
	}
	val := p.items[indexOfLowest]
	p.items = slices.Delete(p.items, indexOfLowest, indexOfLowest+1)
	return val
}

func (p *PriorityQueue) Len() int {
	return len(p.items)
}
