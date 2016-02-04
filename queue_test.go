package gotastructures

import "testing"

func TestQueue(t *testing.T) {
	queue := new(Queue)

	if !queue.IsEmpty() {
		t.Errorf("IsEmpty() on empty queue should return true but returns false")
	}

	if queue.Size() != 0 {
		t.Errorf("Size() on empty queue should return 0")
	}

	queue.Enqueue(1)

	if queue.IsEmpty() {
		t.Errorf("IsEmpty() on non empty queue should return false but returns true")
	}

	if queue.Size() != 1 {
		t.Errorf("Size() on queue with one element should return 1")
	}

	element := queue.Peek()

	if element.(int) != 1 {
		t.Errorf("expected 1 as result of Peek()")
	}
	if queue.Size() != 1 {
		t.Errorf("Size() on queue with one element should return 1 after Peek()")
	}

	element = queue.Dequeue()

	if element.(int) != 1 {
		t.Errorf("expected 1 as result of Dequeue()")
	}
	if !queue.IsEmpty() {
		t.Errorf("IsEmpty() should return true after Dequeue on a queue with only one element")
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	for expected := 1; !queue.IsEmpty(); expected++ {
		result := queue.Dequeue().(int)
		if expected != result {
			t.Errorf("expected %d as result of Dequeue() but got %d", expected, result)
		}
	}

	if !queue.IsEmpty() {
		t.Errorf("queue should be empty after removing all items but has %d items left", queue.Size())
	}
}
