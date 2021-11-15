package ds

import (
	"errors"
	"math"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"github.com/golang-collections/collections/stack"
)

type Criteria func([]float64) bool

func Bfs(graph [][]float64) error {
	visited := set.New()
	queue := queue.New()
	queue.Enqueue(0)
	for node := queue.Dequeue(); node != nil; node = queue.Dequeue() {
		nodeInt, ok := node.(int)

		if !ok {
			return errors.New("wrong type cust")
		}

		visited.Insert(nodeInt)

		for linkedNode := range graph[nodeInt] {
			if linkedNode != 0 && !visited.Has(linkedNode) {
				queue.Enqueue(linkedNode)
			}
		}
	}

	return nil
}

func Dfs(graph [][]int) error {
	visited := set.New()
	stack := stack.New()
	stack.Push(0)
	for node := stack.Pop(); node != nil; node = stack.Pop() {
		nodeInt, ok := node.(int)
		if !ok {
			return errors.New("wrong type cust")
		}
		visited.Insert(nodeInt)

		for linkedNode := range graph[nodeInt] {
			if linkedNode != 0 && !visited.Has(linkedNode) {
				stack.Push(linkedNode)
			}
		}
	}

	return nil
}

func Dijkstra(start int, graph [][]int) []int {
	n := len(graph)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = math.MaxInt32
	}

	d[start] = 0
	u := make([]bool, n)

	for i := 0; i < n; i++ {
		v := -1
		for j := 0; j < n; j++ {
			if !u[j] && (v == -1 || d[j] < d[v]) {
				v = j
			}
		}
		if d[v] == math.MaxInt32 {
			break
		}
		u[v] = true

		for j := 0; j < n; j++ {

			if j == 0 {
				continue
			}
			len := graph[v][j]

			if d[v]+len < d[j] {
				d[j] = d[v] + len
			}
		}
	}

	return d
}
