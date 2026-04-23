package queuefifo

import "errors"

type ArrayQueue struct {
	v     []int
	front int
	back  int
	size  int
}

func (queue *ArrayQueue) Enqueue(val int) error {
	if queue.size == len(queue.v) {
		return errors.New("queue is full")
	}

	queue.back = (queue.back + 1) % len(queue.v)
	queue.v[queue.back] = val
	queue.size++

	return nil
}

func (queue *ArrayQueue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	val := queue.v[queue.front]
	queue.front = (queue.front + 1) % len(queue.v)
	queue.size--

	return val, nil
}

func (queue *ArrayQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	return queue.v[queue.front], nil
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *ArrayQueue) Size() int {
	return queue.size
}
