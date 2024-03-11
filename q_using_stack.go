package main

import (
	"fmt"
)

const SIZE = 5

var stack1, stack2 [SIZE]int
var top1, top2 int = -1, -1

func isFull() bool {
	return top1 == SIZE-1
}

func isEmpty() bool {
	return top1 == -1
}

func enqueue(num int) {
	if isFull() {
		fmt.Println("\nQueue is full")
	} else {
		top1++
		stack1[top1] = num
	}
}

func dequeue() int {
	if isEmpty() {
		fmt.Println("\nQueue is empty")
		return -1
	} else {
		if top2 == -1 {
			for top1 != -1 {
				top2++
				stack2[top2] = stack1[top1]
				top1--
			}
		}
		front := stack2[top2]
		top2--
		stack1 = stack2
		top1 = top2
		return front
	}
}

func display() {
	if isEmpty() {
		fmt.Println("\nQueue is empty")
	} else {
		fmt.Println("\nQueue contents:")
		for i := top2; i >= 0; i-- {
			fmt.Println(stack2[i])
		}
		for i := 0; i <= top1; i++ {
			fmt.Println(stack1[i])
		}
	}
}

func main() {
	var num, option int
	for option != 4 {
		fmt.Println("\n1. Enqueue\n2. Dequeue\n3. Display\n4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Enter number to be enqueued: ")
			fmt.Scan(&num)
			enqueue(num)
		case 2:
			num = dequeue()
			if num != -1 {
				fmt.Printf("%d was dequeued from the queue\n", num)
			}
		case 3:
			display()
		case 4:
			break
		default:
			fmt.Println("Invalid option")
		}
	}
}

