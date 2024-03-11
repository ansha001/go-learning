package main

import "fmt"

func printPrevSmaller(arr []int) {
	n := len(arr)

	fmt.Print("-1, ")

	for i := 1; i < n; i++ {
		var j int
		for j = i - 1; j >= 0; j-- {
			if arr[j] < arr[i] {
				fmt.Printf("%d, ", arr[j])
				break
			}
		}
		if j == -1 {
			fmt.Print("-1, ")
		}
	}
}

func main() {
	arr := []int{2, 5, 3, 7, 8, 1, 4, 6, 9}
	printPrevSmaller(arr)
}
