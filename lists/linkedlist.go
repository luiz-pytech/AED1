package lists

import "errors"

type LinkedList struct {
	head *listNode
	size int
}

type listNode struct {
	val  int
	next *listNode
}

func (list *LinkedList) Add(value int) {
	newNode := &listNode{val: value, next: nil}
	if list.size == 0 {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	list.size++
}

func (list *LinkedList) AddOnIndex(value int, index int) error {
	if index < 0 || index >= list.size {
		return errors.New("index not valid")
	}

	newNode := &listNode{val: value, next: nil}
	if index == 0 {
		newNode.next = list.head
		list.head = newNode
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
	list.size++
	return nil
}

func (list *LinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.size {
		return errors.New("index not valid")
	}

	if index == 0 {
		list.head = list.head.next
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	list.size--
	return nil
}

func (list *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.size {
		return 0, errors.New("index not valid")
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.val, nil
}

func (list *LinkedList) Set(value int, index int) error {
	if index < 0 || index >= list.size {
		return errors.New("index not valid")
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.val = value
	return nil
}

func (list *LinkedList) Size() int {
	return list.size
}
