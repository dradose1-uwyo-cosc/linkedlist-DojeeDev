package ds

type Stack struct {
    list LinkedList
}

// add new head node
func (s *Stack) Push(value string) {
	s.list.Insert(value)
}

// remove the head
func (s *Stack) Pop() (string, bool) {
	if s.list.IsEmpty() {
		return "", false
	} else {
		top := s.list.Head.data
		s.list.RemoveAt(0)
		return top, true
	}
}
