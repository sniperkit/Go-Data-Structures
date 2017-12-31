package ArrayQueue

// A Queue is a first in, first out data structure.
type Queue struct {
}

// Queue constructor which creates an empty Queue.
func new() *Queue {
}

// Empty empties the Queue.
func (queue *Queue) Empty() {
}

// IsEmpty returns true if the Queue is empty and false if not.
func (queue *Queue) IsEmpty() bool {
}

// Len returns the length of the Queue.
func (queue *Queue) Len() int {
}

// Peek returns next Node to be dequeued.
func (queue *Queue) Peek() interface{} {
}

// Enqueue adds a Node to the end of the Queue.
func (queue *Queue) Enqueue(Data interface{}) {
}

// Dequeue removes the first Node in the Queue and returns it.
func (queue *Queue) Dequeue() interface{} {
}
