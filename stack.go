package main

import (
	"fmt"
)

const SIZE = 5

var stack [SIZE]int
var top = -1

func isFull() bool {
	return top == SIZE-1
}

func isEmpty() bool {
	return top == -1
}

func push(num int) {
	if isFull() {
		fmt.Println("\nstack overflow!")
	} else {
		top++
		stack[top] = num
	}
}

func pop() int {
	var n int
	if isEmpty() {
		fmt.Println("\nstack underflow!")
		return -1
	} else {
		n = stack[top]
		top--
	}
	return n
}

func display() {
	if isEmpty() {
		fmt.Println("\nstack is empty")
	} else {
		fmt.Println("\nstack contents:")
		for i := top; i >= 0; i-- {
			fmt.Println(stack[i])
		}
	}
}

func main() {
	var num, option int
	for option != 4 {
		fmt.Println("\n1. Push\n2. Pop\n3. Display\n4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Enter number to be pushed: ")
			fmt.Scan(&num)
			push(num)
		case 2:
			num = pop()
			if num != -1 {
				fmt.Printf("%d was popped out of stack\n", num)
			}
		case 3:
			display()
		}
	}
}
