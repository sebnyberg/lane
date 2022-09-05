package lane

import (
	"fmt"
	"testing"
)

var v int
var ok bool

func BenchmarkStack(b *testing.B) {
	for _, sz := range []int{100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("Stack-%d", sz), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				stack := NewStack[int]()
				for i := 0; i < sz; i++ {
					stack.Push(i)
				}
				for i := 0; i < sz; i++ {
					v, ok = stack.Head()
					v, ok = stack.Pop()
				}
			}
		})
	}
}

func TestStack(t *testing.T) {
	t.Parallel()

	// Helpers
	requireTrue := func(t *testing.T, cond bool, fmt string, args ...interface{}) {
		if !cond {
			t.Fatalf(fmt, args...)
		}
	}
	requireFalse := func(t *testing.T, cond bool, fmt string, args ...interface{}) {
		requireTrue(t, !cond, fmt, args...)
	}

	stack := NewStack([]int{1, 2}...)

	v, ok := stack.Head()
	requireTrue(t, ok, "Head() must return ok when there are items")
	requireTrue(t, v == 2, "Head() must return last item, expected: %q, got: %q", 2, v)

	stack.Push(3)
	v, ok = stack.Head()
	requireTrue(t, ok, "Head() must return ok when there are items")
	requireTrue(t, v == 3, "Head() must return last item, expected: %q, got: %q", 3, v)

	x, ok := stack.Pop()
	requireTrue(t, ok, "Head() must return ok when there are items")
	requireTrue(t, x == 3, "Head() must return last item, expected: %q, got: %q", 3, v)

	// Empty the stack then try out-of-bounds Pop()/Head()
	stack.Pop()
	stack.Pop()
	x, ok = stack.Pop()
	requireFalse(t, ok, "Pop() must return not ok when there are no items")
	requireTrue(t, x == 0, "Pop() must return zero when there are no items")
	x, ok = stack.Head()
	requireFalse(t, ok, "Head() must return not ok when there are no items")
	requireTrue(t, x == 0, "Head() must return zero when there are no items")
}
