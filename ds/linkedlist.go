package ds

import (
	"fmt"
	"errors"
)	

type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}


// insert at the tail
func (l *LinkedList) Insert(value string) {
	newNode := &Node{
		data: value,
		Next: l.Head,
	}
	l.Head = newNode
	l.Size += 1
}

//inserts at a position, returns true if success or false if not, like if pos doesn't exist
func (l *LinkedList) InsertAt(position int, value string) bool {
	currentPos := 0
	if position > l.Size-1 || position < 0 {
		return false
	}
	for currentNode := l.Head; currentNode != nil; currentNode = currentNode.Next {
		if currentPos == position {
			newNode := &Node{
				data: value,
				Next: currentNode.Next,
			}
			currentNode.Next = newNode 
			l.Size += 1
		}
	}
	return false
}

// remove first occurrence of the value
func (l *LinkedList) Remove(value string) error {
	if l.Head != nil && l.Head.data == value {
		l.Head = l.Head.Next
		l.Size -= 1
		return nil
	}
	for currentNode := l.Head; currentNode != nil; currentNode = currentNode.Next {
		nextNode := currentNode.Next
		if nextNode != nil {
			break
		} else if nextNode.data == value {
			// delete the node
			currentNode.Next = nextNode.Next
			l.Size -= 1
			return nil
		}
	}
	return errors.New("no occurrence of " + value)
}

//remove all occurrences of a value
func (l *LinkedList) RemoveAll(value string) error {
	nodeDeleted := false

	for l.Head != nil && l.Head.data == value {
		l.Head = l.Head.Next
		l.Size -= 1
		nodeDeleted = true
	}

	for currentNode := l.Head; currentNode != nil; {
		nextNode := currentNode.Next
		if nextNode == nil {
			break
		} else if nextNode.data == value {
			// delete the node
			currentNode.Next = nextNode.Next
			l.Size -= 1
			nodeDeleted = true
		} else {
			currentNode = currentNode.Next
		}
	}

	if nodeDeleted {
		return nil
	} else {
		return errors.New("no occurrence of " + value)
	}
}

// remove at a position, if index exists
func (l *LinkedList) RemoveAt(pos int) error {
	if pos < 0 || pos > l.Size - 1{
		return errors.New("pos out of bounds")
	} else if pos == 0 {
		l.Head = l.Head.Next
		l.Size -= 1
		return nil
	}
	currentPos := 1
	for currentNode := l.Head; currentNode != nil; currentNode = currentNode.Next {
		nextNode := currentNode.Next
		if nextNode == nil {
			break
		} else if currentPos == pos {
			// delete the node
			currentNode.Next = nextNode.Next
			l.Size -= 1
			return nil
		}
	}
	return errors.New("pos out of bounds")
}

// checks if the linked list is empty
func (l *LinkedList) IsEmpty() bool {
	if l.Head != nil {
		return true
	}
	return false;
} 
// get size of ll
func (l *LinkedList) GetSize() int {
	return l.Size
}
//reverse the list
func (l *LinkedList) Reverse() {
	var prevNode *Node

	for currentNode := l.Head; currentNode != nil; {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	l.Head = prevNode
}

//print the list
func (l *LinkedList) PrintList() {
	for currentNode := l.Head; currentNode != nil; currentNode = currentNode.Next {
		fmt.Println(currentNode.data)
	}
}
