package main

import (
	"fmt"
)

type node struct {
	left  *node
	right *node
	data  int
}

type BinaryTree struct {
	root *node
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *node) insert(data int) {
	if n == nil {
		return
	} else if data >= n.data { //BST logic is implemented
		if n.right == nil {
			n.right = &node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	} else {
		if n.left == nil {
			n.left = &node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	}
}

func (tree *BinaryTree) inorder(n *node) { //non-recursive
	if n==nil{
		fmt.Print("No data")
		return
	}
	stack := []*node{}
	curr := n

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(curr.data," ")

		curr = curr.right
	}
}

func (tree * BinaryTree ) preorder(n *node){
	if n==nil{
		fmt.Print("No data")
		return
	}
	stack:=[]*node{}
	curr:=n

	for curr!=nil|| len(stack)>0{
		for curr!=nil{
			stack=append(stack,curr)
			fmt.Print(curr.data," ")
			curr = curr.left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr = curr.right

	}
}

func (tree *BinaryTree) postorder(n *node) {
	if n==nil{
		fmt.Print("No data")
		return
	}
    stack := []*node{}
    curr := n
    var prev *node

    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.left
        }
        curr = stack[len(stack)-1]

        if curr.right == nil || curr.right == prev {
            fmt.Print(curr.data, " ")
            stack = stack[:len(stack)-1]
            prev = curr
            curr = nil
        } else {
            curr = curr.right
        }
    }
}

func verticalSum(root *node, hd int, sums map[int]int) {
	if root == nil {
		return
	}

	verticalSum(root.left, hd-1, sums)
	sums[hd] += root.data
	verticalSum(root.right, hd+1, sums)
}

func calculate_vs(root *node) {
	sums := make(map[int]int)
	verticalSum(root, 0, sums)

	for hd, sum := range sums {
		fmt.Printf("(%d): %d\n", hd, sum)
	}
}

func main() {
	tree := &BinaryTree{}
	var item, num, choice int

	for choice != 6 {
		fmt.Println("\n\n1.Custom input")
		fmt.Println("2.print inorder")
		fmt.Println("3.print preorder")
		fmt.Println("4.print postorder")
		fmt.Println("5.vertical sum")
		fmt.Println("6.exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter number of inputs: ")
			fmt.Scan(&num)
			for i := 0; i < num; i++ {
				fmt.Print("Enter Data:")
				fmt.Scan(&item)
				tree.insert(item)
			}
		case 2:
			fmt.Println("\ninorder:")
			tree.inorder(tree.root)
		case 3:
			fmt.Println("\npreorder:")
			tree.preorder(tree.root)
		case 4:
			fmt.Println("\npostorder:")
			tree.postorder(tree.root)
		case 5:
			fmt.Println("\nposition of column wrt root: vertical sum")
			calculate_vs(tree.root)
		case 6:
			break
		default:
			fmt.Println("Invalid input")
		}
	}
}
