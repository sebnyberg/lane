package lane

// Stack implements a Last In First Out data structure.
//
// Every operation's has a time complexity of *O(1)*.
type Stack[T any] struct {
	items []T
}

// NewStack produces a new Stack instance.
//
// When providing initialization items, those will be inserted as-is: lower
// index being the head of the stack.
func NewStack[T any](items ...T) *Stack[T] {
	var s Stack[T]
	s.items = items
	return &s
}

// Push adds on an item on the top of the Stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item on the top of the Stack.
//
// If the stack is empty, a default-initialized item and false is returned.
func (s *Stack[T]) Pop() (item T, ok bool) {
	n := len(s.items)
	if n == 0 {
		return item, false
	}
	x := s.items[n-1]
	s.items = s.items[:n-1]
	return x, true
}

// Head returns the item on the top of the Stack.
//
// If the stack is empty, a default-initialized item and false is returned.
func (s Stack[T]) Head() (item T, ok bool) {
	n := len(s.items)
	if n == 0 {
		return item, false
	}
	return s.items[n-1], true
}

// Size returns the size of the Stack.
func (s Stack[T]) Size() int { return len(s.items) }
