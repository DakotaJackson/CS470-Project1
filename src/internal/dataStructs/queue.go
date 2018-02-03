package dataStructs

// Following is a basic queue with basic functionality, following
//	the FIFO standard.

// Queue is the data structure used for algorithms such as bfs.
type Queue struct {
	length int
	start  *nodeQ
	end    *nodeQ
}

// node is used in the queue data structure (only using ints for map)
type nodeQ struct {
	value int
	next  *nodeQ
}

// InitQueue creates a blank queue for use elsewhere.
func InitQueue() *Queue {
	return &Queue{
		0,
		nil,
		nil,
	}
}

// Enqueue adds a node to the end of the queue.
func (q *Queue) Enqueue(value int) int {
	n := &nodeQ{value, nil}
	if q.length != 0 {
		// queue has nodes in it
		q.end.next = n
		q.end = n
	} else {
		// queue has nothing in it
		q.end = n
		q.start = n
	}
	// add one to length
	q.length++
	return value
}

// Dequeue removed a node from the front of the queue.
// 	It also returns false if the queue has no items to return from.
func (q *Queue) Dequeue() (int, bool) {
	if q.length == 0 {
		return 0, false
	}
	n := q.start
	if q.length > 1 {
		// queue has more than one item, proceed as normal
		q.start = q.start.next
	} else {
		// queue has only one item and needs special case
		q.start = nil
		q.end = nil
	}
	// decrement length
	q.length--
	return n.value, true
}

// Peek returns the value of the node at the front of the queue.
// 	It also has a boolean value if the node exists or not (false if empty, true if not).
func (q *Queue) Peek() (int, bool) {
	if q.length != 0 {
		return q.start.value, true
	}
	return 0, false
}

// GetLenQ returns the length of the queue.
func (q *Queue) GetLenQ() int {
	return q.length
}

// IsEmpty returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	if q.length <= 0 {
		return true
	}
	return false
}
