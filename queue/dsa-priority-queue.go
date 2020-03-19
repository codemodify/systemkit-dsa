package queues

import (
	"container/heap"
)

// PriorityQueue -
type PriorityQueue []*Item

// An Item is something we manage in a priority queue.
type Item struct {
	Data     interface{}
	Priority int64 // The priority of the item in the queue.
	Index    int   // The index of the item in the heap.
	// The index is needed by update and is maintained by the heap.Interface methods.
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	//we are using this for timestamp, where oldest has the highest priority
	//so lowest numeric value is highest priority
	return pq[j].Priority > pq[i].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push -
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop -
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update - update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, value *interface{}, priority int64) {
	item.Data = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// Dump -
func (pq *PriorityQueue) Dump() []*Item {
	return *pq
}

// Peek - Get the top Priority to support making decision on priority from calling code
func (pq PriorityQueue) Peek() int64 {
	length := pq.Len()
	if length > 0 {
		return pq[length-1].Priority
	}
	return -1
}
