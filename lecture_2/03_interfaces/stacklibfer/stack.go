package stacklibfer

type stack struct {
	list []int
}

func New() *stack {
	return &stack{}
}

func (s *stack) Push(a int) {
	s.list = append(s.list, a)
}

func (s *stack) Pop() (int, bool) {
	if len(s.list) == 0 {
		return 0, false
	}

	res := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return res, true
}
