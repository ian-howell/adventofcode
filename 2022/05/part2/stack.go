package main

type Stack []string

func (s *Stack) Push(item string) {
	*s = append(*s, item)
}

func (s *Stack) Pop() string {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s Stack) Top() string {
	return s[len(s)-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}
