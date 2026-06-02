package lists

import "errors"

type Node struct {
	val  int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (list *DoublyLinkedList) Add(value int) {
	newNode := &Node{val: value}

	if list.size == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
}

func (list *DoublyLinkedList) AddFirst(value int) {
	newNode := &Node{val: value}

	if list.size == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}

	list.size++
}

func (list *DoublyLinkedList) getNode(index int) (*Node, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index not valid")
	}

	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current, nil
}

func (list *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index < 0 || index > list.size {
		return errors.New("index not valid")
	}

	if index == 0 {
		list.AddFirst(value)
		return nil
	}

	if index == list.size {
		list.Add(value)
		return nil
	}

	current, err := list.getNode(index)
	if err != nil {
		return err
	}

	newNode := &Node{val: value}
	prevNode := current.prev

	newNode.prev = prevNode
	newNode.next = current

	prevNode.next = newNode
	current.prev = newNode

	list.size++
	return nil
}

func (list *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.size {
		return errors.New("index not valid")
	}

	if list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size--
		return nil
	}

	if index == 0 {
		list.head = list.head.next
		list.head.prev = nil
		list.size--
		return nil
	}

	if index == list.size-1 {
		list.tail = list.tail.prev
		list.tail.next = nil
		list.size--
		return nil
	}

	current, err := list.getNode(index)
	if err != nil {
		return err
	}

	prevNode := current.prev
	nextNode := current.next

	prevNode.next = nextNode
	nextNode.prev = prevNode

	list.size--
	return nil
}

func (list *DoublyLinkedList) Size() int {
	return list.size
}
