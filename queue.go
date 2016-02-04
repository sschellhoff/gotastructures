package gotastructures

type Queue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

type QueueNode struct {
	data interface{}
	next *QueueNode
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.head.data
}

func (q *Queue) Enqueue(data interface{}) {
	if q.IsEmpty() {
		q.head = &QueueNode{data, nil}
		q.tail = q.head
	} else {
		q.tail.next = &QueueNode{data, nil}
		q.tail = q.tail.next
	}
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.head.data
	q.size--
	if q.IsEmpty() { // size was reduced now
		q.head = nil
		q.tail = nil
		return data
	}
	q.head = q.head.next
	return data
}
