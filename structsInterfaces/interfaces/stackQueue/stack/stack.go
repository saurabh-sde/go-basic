package stack

import "fmt"

type Stack struct {
	Data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n int) {
	s.Data = append(s.Data, n)
}

func (s *Stack) Pop() int {
	if len(s.Data) == 0 {
		return -1
	}
	top := s.Data[len(s.Data)-1]
	s.Data = s.Data[0 : len(s.Data)-1]
	return top
}

func (s *Stack) Print() {
	fmt.Println("Data: ", s.Data)
}
