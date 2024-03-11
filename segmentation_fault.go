// 1. Cases for segementation fault in Go:
// 		a. Stack Overflow
// 		b. Dereferencing nil pointer
// 		c. Accessing out-of-array index bounds

package main

import "fmt"

func random() int {
	return random()
}

func main() {

	//case a: stack overflow (at run time)
	random()

	//case b: nil pointer dereference segfault (runtime error)

	var nilptr *int = nil
	*nilptr = 100

	//case c: index-out of bound error during compile time

	arr := [3]int{1, 3, 5}
	fmt.Println(arr[3])

}
