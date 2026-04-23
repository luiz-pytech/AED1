package stacks

import "errors"

type ArrayStack struct {
	values   []int
	inserted int
}

func (stack *ArrayStack) Push(value int) {
	if stack.inserted == len(stack.values) {
		newValues := make([]int, len(stack.values)*2+1)
		copy(newValues, stack.values)
		stack.values = newValues
	}

	stack.values[stack.inserted] = value
	stack.inserted++
}

func (stack *ArrayStack) Pop() (int, error) {
	if stack.inserted == 0 {
		return 0, errors.New("stack is empty")
	}

	stack.inserted--
	return stack.values[stack.inserted], nil
}

func (stack *ArrayStack) Peek() (int, error) {
	if stack.inserted == 0 {
		return 0, errors.New("stack is empty")
	}

	return stack.values[stack.inserted-1], nil
}

func (stack *ArrayStack) IsEmpty() bool {
	if stack.inserted == 0 {
		return true
	}
	return false
}

func (stack *ArrayStack) Size() int {
	return stack.inserted
}
