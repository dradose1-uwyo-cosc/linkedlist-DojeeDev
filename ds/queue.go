package ds

import (
	"errors"
)

type Queue struct {
    list LinkedList
}

// add tail node
func (q *Queue) Push(value string) {
	if q.list.IsEmpty() {
		q.list.Insert(value)
	} else {
		q.list.InsertAt(q.list.Size, value)
	}
}

// remove the head
func (q *Queue) Pop() (string, error) {
	if q.list.IsEmpty() {
		return "", errors.New(" is empty")
	} else {
		front := q.list.Head.data
		q.list.RemoveAt(0)
		return front, nil
	}
}
