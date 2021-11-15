package heap

import (
	"ds/interfaces"
)

type Heap []interfaces.Comparable

func (heap Heap) Len() int { return len(heap) }

func (heap Heap) Less(i, j int) bool {

	return heap[i].Compare(heap[j]) < 0
}

func (heap Heap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *Heap) Push(x interface{}) {
	item := x.(interfaces.Comparable)
	*heap = append(*heap, item)
}

func (heap *Heap) Pop() interface{} {
	old := *heap
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*heap = old[0 : n-1]
	return item
}
