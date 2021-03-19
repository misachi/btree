package main

import "fmt"

func main() {
	degree := 4
	root := CreateNewNode(degree)
	root.Keys[0] = 3
	root.NumKeys += 1
	bt := BTree{Root: root, Degree: degree}
	bt.Insert(6)
	bt.Insert(4)
	bt.Insert(5)
	bt.Insert(2)
	bt.Insert(7)
	bt.Insert(8)
	bt.Insert(9)
	bt.Insert(10)
	bt.Insert(1)
	bt.Insert(11)
	bt.Insert(12)
	bt.Insert(2)
	bt.Insert(22)

	/*
		Check the root and child nodes
		Uncomment to see the first fee nodes including root
	*/
	// fmt.Println(bt.Root.Keys)
	// fmt.Println(bt.Root.Children[0].Keys, bt.Root.Children[0].NumKeys)
	// fmt.Println(bt.Root.Children[1].Keys, bt.Root.Children[1].NumKeys)

	node, idx := bt.Search(bt.Root, 22)
	if node != nil {
		fmt.Println(node.Keys[idx])
	} else {
		fmt.Println("Not Found")
	}
}
