package main

import (
	"fmt"
)

const SIZE = 5

var q [SIZE]int
var front, rear, num int

func isempty() bool {
	return rear == -1 && front == -1
}

func isfull() bool {
	return rear == SIZE-1
}

func enqueue() {
	if isfull() {
		fmt.Println("\nqueue is full")
	} else {
		fmt.Print("enter number to be pushed to queue: ")
		fmt.Scan(&num)
		if isempty() {
			front = 0
			rear++
			q[rear] = num
		} else {
			rear++
			q[rear] = num
		}
	}
}

func dequeue() {
	if isempty() {
		fmt.Println("\nqueue is empty")
	} else if rear == front {
		rear, front = -1, -1
	} else {
		front++
	}
}

func display() {
	if isempty() {
		fmt.Println("\nqueue is empty")
	} else {
		fmt.Println("\nqueue:")
		for i := front; i <= rear; i++ {
			fmt.Println(q[i])
		}
	}
}

func main() {
	var option int
	for option != 4 {
		fmt.Println("\n1. Enqueue\n2. Dequeue\n3. Display\n4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			enqueue()
		case 2:
			dequeue()
		case 3:
			display()
		case 4:
			break
		default:
			fmt.Println("invalid option")
		}
	}
}
