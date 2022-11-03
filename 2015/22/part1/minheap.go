package main

import "errors"

type minHeapNode[T any] struct {
	value    T
	priority int
}

// A MinHeap implements heap.Interface and holds Items.
type MinHeap[T any] []*minHeapNode[T]

func (mh MinHeap[T]) Len() int           { return len(mh) }
func (mh MinHeap[T]) Less(i, j int) bool { return mh[i].priority < mh[j].priority }

func (mh *MinHeap[T]) Push(x T, priority int) {
	newNode := &minHeapNode[T]{value: x, priority: priority}
	*mh = append(*mh, newNode)

	index := len(*mh) - 1
	parentIndex := (index - 1) / 2
	for index > 0 && mh.Less(index, parentIndex) {
		mh.swap(index, parentIndex)
		index = parentIndex
		parentIndex = (index - 1) / 2
	}

}

func (mh *MinHeap[T]) Pop() (T, int, error) {
	if len(*mh) == 0 {
		return *new(T), -1, errors.New("heap is empty")
	}

	val, priority := (*mh)[0].value, (*mh)[0].priority
	(*mh)[0] = (*mh)[mh.Len()-1]
	(*mh)[mh.Len()-1] = nil // avoid memory leaks
	*mh = (*mh)[:mh.Len()-1]
	mh.heapify(0)
	return val, priority, nil
}

func (mh *MinHeap[T]) swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *MinHeap[T]) heapify(i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < mh.Len() && mh.Less(left, smallest) {
		smallest = left
	}

	if right < mh.Len() && mh.Less(right, smallest) {
		smallest = right
	}

	if smallest != i {
		mh.swap(smallest, i)
		mh.heapify(smallest)
	}
}
