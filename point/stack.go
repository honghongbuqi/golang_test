package main

import (
	"errors"
	"fmt"
)

var (
	ErrStackFull  = errors.New("stack full")
	ErrStackEmpty = errors.New("stack empty")
)

type Stack struct {
	ptr []int
	sp  int
}

//cap:	堆棧容量
//sp: 堆栈剩余大小
func NewStack(vol int, sp int) *Stack {
	return &Stack{ptr: make([]int, vol), sp: sp}
}

//入栈
func (s *Stack) Push(v int) error {
	if s.sp == 0 {
		return ErrStackFull
	}
	s.sp--
	s.ptr[s.sp] = v
	return nil
}

//出栈
func (s *Stack) Pop() (int, error) {
	if s.sp == len(s.ptr) {
		return 0, ErrStackEmpty
	}
	v := s.ptr[s.sp]
	s.sp++
	return v, nil
}

func main() {
	s := NewStack(5, 5)
	s.Push(3)
	s.Push(5)
	s.Push(7)
	s.Push(9)
	s.Push(11)
	s.Push(13)
	s.Push(15)

	for i := 0; i < len(s.ptr); i++ {
		v, _ := s.Pop()
		fmt.Println(v) // 11 9 7 5 3
	}
}
