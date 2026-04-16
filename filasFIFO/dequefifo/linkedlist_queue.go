package queuefifo

import "errors"

type LinkedListQueue struct {
	front *Node
	back  *Node
	size  int
}

type Node struct {
	val  int
	next *Node
}

func (queue *LinkedListQueue) Enqueue(val int) {
	newNode := &Node{val: val, next: nil}

	if queue.size == 0 {
		queue.front = newNode
	} else {
		queue.back.next = newNode
	}

	queue.back = newNode
	queue.size++
}

func (queue *LinkedListQueue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	val := queue.front.val
	queue.front = queue.front.next
	queue.size--

	if queue.size == 0 {
		queue.back = nil
	}

	return val, nil
}

func (queue *LinkedListQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	return queue.front.val, nil
}

func (queue *LinkedListQueue) IsEmpty() bool {
	if queue.size == 0 {
		return true
	}
	return false
}

func (queue *LinkedListQueue) Size() int {
	return queue.size
}
