package queuefifo

import "errors"

type ArrayQueue struct {
	v     []int
	front int
	back  int
	size  int
}

func (queue *ArrayQueue) Init(size int) {
	queue.v = make([]int, size)
	queue.front = 0
	queue.back = -1
	queue.size = 0

}

func (queue *ArrayQueue) Enqueue(val int) {
	if queue.size == len(queue.v)-1 {

	}
}

func (queue *ArrayQueue) Dequeue() (int, error) {
	return -1, errors.New("error msg")
}

func (queue *ArrayQueue) Front() (int, error) {
	return -1, errors.New("error msg")
}

func (queue *ArrayQueue) IsEmpty() bool {
	if queue.size == 0 {
		return true
	}

	return false
}

func (queue *ArrayQueue) Size() int {
	return queue.size
}
