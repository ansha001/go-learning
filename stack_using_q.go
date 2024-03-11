package main

import (
	"fmt"
)

const SIZE = 5

var q1 [SIZE]int
var q2 [SIZE]int
var front1, rear1, front2, rear2 int = -1,-1,-1,-1
var num int


func isEmpty() bool {
	return front1 == -1 && rear1 == -1
}

func isFull() bool {
	return rear1 == SIZE-1
}

func push() {
	if isFull() {
		fmt.Println("\nstack is full")
	} else {
		fmt.Print("Enter number to be pushed: ")
		fmt.Scan(&num)
		if isEmpty() {
			front1, rear1 = 0, 0
			q1[rear1] = num
		} else {
			rear1++
			q1[rear1] = num
		}
	}
}

func pop() {
	if isEmpty() {
		fmt.Println("\nstack is empty")
	} else {
		for front1 != rear1 {
			rear2++
			q2[rear2] = q1[front1]
			front1++
		}
		fmt.Printf("Popped element: %d\n", q1[front1])
		front1, rear1 = -1, -1
		front1, rear1 = front2, rear2
		front2, rear2 = -1, -1
	}
}

func display() {
	if isEmpty() {
		fmt.Println("\nstack is empty")
	} else {
		fmt.Println("\nstack contents:")
		for i := front1; i <= rear1; i++ {
			fmt.Println(q1[i])
		}
	}
}

func main() {
	var option int
	for option != 4 {
		fmt.Println("\n1. Push\n2. Pop\n3. Display\n4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			push()
		case 2:
			pop()
		case 3:
			display()
		case 4:
			break
		default:
			fmt.Println("Invalid option")
		}
	}
}
