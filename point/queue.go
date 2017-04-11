package main

import (
	"errors"
	"fmt"
)

//自定义ERR
var (
	ErrQueFull  = errors.New("CirQue is Full")
	ErrQueEmpty = errors.New("CirQue is Empty")
)

//环形队列结构
type CirQue struct {
	data []int //队列数据存储区
	head int   //队列头
	tail int   //队列尾
}

//初始化环形队列
func NewCQ(vol int) *CirQue {
	return &CirQue{data: make([]int, vol), head: 0, tail: 0}
}

//出队列
func (cq *CirQue) Pop() (int, error) {
	if cq.head == cq.tail {
		return 0, ErrQueEmpty
	}
	n := cq.head % len(cq.data)
	v := cq.data[n]
	cq.head++
	return v, nil
}

//入队列
func (cq *CirQue) Push(v int) error {
	if len(cq.data) == (cq.tail - cq.head) {
		return ErrQueFull
	}
	n := cq.tail % len(cq.data)
	cq.data[n] = v
	cq.tail++
	return nil
}

func main() {
	cq := NewCQ(5)
	cq.Push(1)
	cq.Push(3)
	cq.Push(5)
	cq.Push(7)
	cq.Push(9)
	err := cq.Push(11)
	fmt.Println(err)

	for i := 0; i < len(cq.data); i++ {
		v, err := cq.Pop() // 1 3 5 7 9
		fmt.Println(v, err)
	}
	v, err := cq.Pop()
	fmt.Println(v, err) //0 CirQue is Empty
}
