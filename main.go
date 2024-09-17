package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	lowerHalf MaxHeap
	upperHalf MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		lowerHalf: MaxHeap{},
		upperHalf: MinHeap{},
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.lowerHalf.Len() == 0 || num <= mf.lowerHalf.Peek() {
		heap.Push(&mf.lowerHalf, num)
	} else {
		heap.Push(&mf.upperHalf, num)
	}

	if mf.lowerHalf.Len() > mf.upperHalf.Len()+1 {
		heap.Push(&mf.upperHalf, heap.Pop(&mf.lowerHalf))
	} else if mf.upperHalf.Len() > mf.lowerHalf.Len() {
		heap.Push(&mf.lowerHalf, heap.Pop(&mf.upperHalf))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.lowerHalf.Len() > mf.upperHalf.Len() {
		return float64(mf.lowerHalf.Peek())
	}
	return (float64(mf.lowerHalf.Peek()) + float64(mf.upperHalf.Peek())) / 2.0
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MaxHeap) Peek() int {
	return h[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], j = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MinHeap) Peek() int {
	return h[0]
}

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian()) // 1.5
	mf.AddNum(3)
	fmt.Println(mf.FindMedian()) // 2
}