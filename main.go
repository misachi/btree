package main

import "fmt"

func main() {
	degree := 5
	root := CreateNewNode(degree)
	bt := BTree{Root: root, Degree: degree}
	bt.Insert(4)
	bt.Insert(5)

	node, idx := bt.Search(bt.Root, 5)
	if node != nil {
		fmt.Println(node.Keys[idx])
	} else {
		fmt.Println("Not Found")
	}
}
