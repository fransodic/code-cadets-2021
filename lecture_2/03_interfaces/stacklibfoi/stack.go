package stacklibfoi

import "container/list"

type stack struct {
	list *list.List
}

func New() *stack {
	return &stack{
		list: list.New(),
	}
}

func (s *stack) Push(a int) {
	s.list.PushBack(a)
}

func (s *stack) Pop() (int, bool) {
	if s.list.Len() == 0 {
		return 0, false
	}

	lastElement := s.list.Back()
	lastValue := s.list.Remove(lastElement)

	res, ok := lastValue.(int)
	if !ok {
		return 0, false
	}

	return res, true
}
