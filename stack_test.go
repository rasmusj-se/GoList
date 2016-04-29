package GoList

import "testing"

func TestStack(t *testing.T) {
	stack := &Stack{}
	stack.Push("Hello")
	if stack.Pop() != "Hello" {
		t.Error("Stack did not return expected value")
	}
	if stack.Count() != 0 {
		t.Error("Stack count was faulty")
	}
	stack.Push(1234)
	if stack.Pop() != 1234 {
		t.Error("Stack did not handle int value")
	}
	type test struct {
		Value string
	}
	stack.Push(&test{Value: "test"})
	if stack.Pop().(*test).Value != "test" {
		t.Error("Stack did not handle custom value")
	}
	stack.Push("Hello")
	if stack.Peek() != "Hello" || stack.Count() != 1 {
		t.Error("Stack did handle peek correctly")
	}
	if stack.Peek() != "Hello" || stack.Count() != 1 {
		t.Error("Stack did handle peek correctly")
	}
	stack.Pop()
	for i := 0; i <= 512; i++ {
		stack.Push(i)
	}
	for i := 512; i > 0; i-- {
		if stack.Pop() != i || stack.Count() != i {
			t.Error("Stack did not handle large size correctly")
		}
	}
}
