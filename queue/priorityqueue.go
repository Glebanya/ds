package queue

type Item struct {
	Value    interface{}
	Priority int
}

type PriorityQueue []*Item

func New(items ...*Item) *PriorityQueue {

	prQueue := make(PriorityQueue, len(items))
	for i, v := range items {
		prQueue[i] = v
	}
	link := &prQueue

	return link

}

func (prQueue PriorityQueue) Len() int {
	return len(prQueue)
}

func (prQueue PriorityQueue) Less(i, j int) bool {
	return prQueue[i].Priority > prQueue[j].Priority
}

func (prQueue PriorityQueue) Swap(i, j int) {
	prQueue[i], prQueue[j] = prQueue[j], prQueue[i]
}

func (prQueue *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*prQueue = append(*prQueue, item)
}

func (prQueue *PriorityQueue) Pop() interface{} {
	old := *prQueue
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*prQueue = old[0 : n-1]
	return item
}
