package queue

import "fmt"

type Queue struct {
	Data []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Queue) Push(n int) {
	s.Data = append(s.Data, n)
}

func (s *Queue) Pop() int {
	if len(s.Data) == 0 {
		return -1
	}
	front := s.Data[0]
	s.Data = s.Data[1:len(s.Data)]
	return front
}

func (s *Queue) Print() {
	fmt.Println("Data: ", s.Data)
}
