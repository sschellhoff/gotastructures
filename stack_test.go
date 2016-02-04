package gotastructures

import "testing"

func TestStack(t *testing.T) {
	stack := new(Stack)
	if !stack.IsEmpty() {
		t.Errorf("IsEmpty() on a new stack should return true but returned false")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("IsEmpty() on a non empty stack should return false but returned true")
	}

	element := stack.Peek()

	if element == nil {
		t.Errorf("expected 1 as result of Peek() but result was nil")
	}

	if element.(int) != 1 {
		t.Errorf("expected 1 as result of Peek()")
	}

	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	element = stack.Pop()
	if element == nil {
		t.Errorf("expected 4 as result of Pop() but result was nil")
	}
	if element.(int) != 4 {
		t.Errorf("expected 4 as a result of Pop()")
	}

	element = stack.Pop()
	if element == nil {
		t.Errorf("expected 3 as result of Pop() but result was nil")
	}
	if element.(int) != 3 {
		t.Errorf("expected 3 as a result of Pop()")
	}

	element = stack.Pop()
	if element == nil {
		t.Errorf("expected 2 as result of Pop() but result was nil")
	}
	if element.(int) != 2 {
		t.Errorf("expected 2 as a result of Pop()")
	}

	element = stack.Pop()
	if element == nil {
		t.Errorf("expected 1 as result of Pop() but result was nil")
	}
	if element.(int) != 1 {
		t.Errorf("expected 1 as a result of Pop()")
	}

	element = stack.Peek()
	if element != nil {
		t.Errorf("expected nil as result of Peek() on an empty stack")
	}

	element = stack.Pop()
	if element != nil {
		t.Errorf("expected nil as result of Pop() on an empty stack")
	}
}
