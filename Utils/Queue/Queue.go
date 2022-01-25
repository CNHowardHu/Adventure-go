package Queue

import (
	"errors"
)

type Queue []interface{}

func (Q *Queue) Size() int {
	return len(*Q)
}

func (Q *Queue) Empty() bool {
	return len(*Q) == 0
}

func (Q *Queue) Push(e interface{}) {
	*Q = append(*Q, e)
}

func (Q *Queue) Pop() error {
	if len(*Q) == 0 {
		return errors.New("Queue is empty!")
	}
	*Q = (*Q)[1:]
	return nil
}

func (Q *Queue) Front() (interface{}, error) {
	if len(*Q) == 0 {
		return 0, errors.New("Queue is empty!")
	}
	return (*Q)[0], nil
}

func (Q *Queue) Back() (interface{}, error) {
	if L := len(*Q); L == 0 {
		return 0, errors.New("Queue is empty!")
	} else {
		return (*Q)[L-1], nil
	}
}
