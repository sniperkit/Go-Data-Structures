package ArrayStack

// A Node is the data stored inside of a Stack.
type Node struct {
}

// A Stack is a first in, last out data structure.
type Stack struct {
}

// Stack constructor which creates an empty Stack.
func new() *Stack {
}

// Empty empties the Stack.
func (stack *Stack) Empty() {
}

// IsEmpty returns true if the Stack is empty.
func (stack *Stack) IsEmpty() bool {
}

// Len returns the length of the Stack.
func (stack *Stack) Len() int {
}

// Peek returns the top Node.
func (stack *Stack) Peek() interface{} {
}

// Push adds a Node to the top of the Stack.
func (stack *Stack) Push(Data interface{}) {
}

// Pop removes the top Node in the stack and returns it.
func (stack *Stack) Pop() interface{} {
}
