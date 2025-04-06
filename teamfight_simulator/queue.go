package main

type Queue interface {
	Enqueue(val int)
	Top() int
	Pop() int
	Size() int
}

type queueData struct {
	current *queueNode
	tail    *queueNode
	size    int
}

type queueNode struct {
	value int
	next  *queueNode
}

func NewQueue() Queue {
	return new(queueData)
}

func (q *queueData) Enqueue(val int) {

	n := &queueNode{
		value: val,
	}

	if q.size == 0 {
		q.current = n
		q.tail = n
		q.size += 1
	} else {
		q.tail.next = n
		q.tail = q.tail.next
		q.size += 1
	}

}

func (q *queueData) Size() int {
	return q.size
}

func (q *queueData) Top() int {
	return q.current.value
}

func (q *queueData) Pop() int {
	if q.size == 0 {
		panic("Pop on an empty queue")
	}

	out := q.current.value
	q.size -= 1

	if q.size > 0 {
		q.current = q.current.next
	}

	return out
}
