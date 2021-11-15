package ds

import (
	"cf/queue"
	"container/heap"
	"fmt"
)

func Problem(n int) {
	items := []int{2, 3, 5}
	pq := make(queue.PriorityQueue, 0)

	heap.Init(&pq)

	for _, v := range items {
		heap.Push(&pq, &queue.Item{
			Value:    v,
			Priority: v,
		})
	}

	for ; n > 0; n-- {
		v := pq.Pop().(int)
		fmt.Println(v)
		for _, val := range items {
			newVal := v * val
			heap.Push(&pq, &queue.Item{
				Value:    newVal,
				Priority: newVal,
			})
		}
	}
}
