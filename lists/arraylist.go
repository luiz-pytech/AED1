package lists

import (
	"errors"
)

type ArrayList struct {
	values   []int
	inserted int
}

func (list *ArrayList) Add(value int) {
	if list.inserted == len(list.values) {
		newValues := make([]int, len(list.values)*2+1)
		copy(newValues, list.values)
		list.values = newValues
	}

	list.values[list.inserted] = value
	list.inserted++
}

func (list *ArrayList) AddOnIndex(value int, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index not valid")
	}

	if list.inserted == len(list.values) {
		newValues := make([]int, len(list.values)*2+1)
		copy(newValues, list.values)
		list.values = newValues
	}

	for i := list.inserted; i > index; i-- {
		list.values[i] = list.values[i-1]
	}
	list.values[index] = value
	list.inserted++
	return nil
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index not valid")
	}

	list.values[index] = 0

	for i := index; i < list.inserted-1; i++ {
		list.values[i] = list.values[i+1]
	}

	list.inserted--
	return nil
}

func (list *ArrayList) Get(index int) (int, error) {
	if index < 0 || index >= list.inserted {
		return 0, errors.New("index not valid")
	}

	return list.values[index], nil
}

func (list *ArrayList) Set(value int, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New("index not valid")
	}

	list.values[index] = value
	return nil
}

func (list *ArrayList) Size() int {
	return list.inserted
}
