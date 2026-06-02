package stacks

import "errors"

type LinkedListStack struct {
	top  *stackNode
	size int
}

type stackNode struct {
	val  int
	next *stackNode
}

func (stack *LinkedListStack) Push(value int) {
	newNode := &stackNode{val: value, next: stack.top}
	stack.top = newNode
	stack.size++
}

func (stack *LinkedListStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	val := stack.top.val
	stack.top = stack.top.next
	stack.size--

	return val, nil
}

func (stack *LinkedListStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return stack.top.val, nil
}

func (stack *LinkedListStack) IsEmpty() bool {
	if stack.size == 0 {
		return true
	}
	return false

}

func (stack *LinkedListStack) Size() int {
	return stack.size
}
