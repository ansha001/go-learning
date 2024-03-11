package main

import (
	"fmt"
)

type dl struct {
	stack  []int
	queue  []int
	middle int
}

func (m *dl) Push(val int) {
	m.stack = append(m.stack, val)
	m.queue = append(m.queue, val)
	if len(m.stack)%2 == 1 {
		m.middle = len(m.stack) / 2
	}
}

func (m *dl) Pop() {
	if len(m.stack) == 0 {
		return
	}
	m.stack = m.stack[:len(m.stack)-1]
	m.queue = m.queue[1:]
	if len(m.stack)%2 == 0 {
		m.middle = len(m.stack) / 2
	}
}

func (m *dl) FindMiddle() int {
	if len(m.stack) == 0 {
		return -1
	}
	return m.stack[m.middle]
}

func (m *dl) DeleteMiddle() {
	if len(m.stack) == 0 {
		return
	}
	copy(m.stack[m.middle:], m.stack[m.middle+1:])
	m.stack = m.stack[:len(m.stack)-1]
	if len(m.stack)%2 == 0 {
		m.middle = len(m.stack) / 2
	}
}

func main() {
	m := dl{}
	m.Push(10)
	m.Push(30)
	m.Push(50)
	fmt.Println("Middle element:", m.FindMiddle())
	m.DeleteMiddle()
	fmt.Println("middle element after deletion:", m.FindMiddle())
}
