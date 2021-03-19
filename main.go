package main

import "fmt"

func main() {
	degree := 4
	root := CreateNewNode(degree)
	root.Keys[0] = 3.3
	root.NumKeys += 1
	bt := BTree{Root: root, Degree: degree}
	bt.Insert(6.0)
	bt.Insert(4.2)
	bt.Insert(5.3)
	bt.Insert(2.4)
	bt.Insert(7.0)
	bt.Insert(8.4)
	bt.Insert(9.0)
	bt.Insert(10.0)
	bt.Insert(1.0)
	bt.Insert(11.0)
	bt.Insert(12.0)
	bt.Insert(2.0)
	bt.Insert(22.0)

	/*
		Check the root and child nodes
		Uncomment to see the first fee nodes including root
	*/
	// fmt.Println(bt.Root.Keys)
	// fmt.Println(bt.Root.Children[0].Keys, bt.Root.Children[0].NumKeys)
	// fmt.Println(bt.Root.Children[1].Keys, bt.Root.Children[1].NumKeys)

	node, idx := bt.Search(bt.Root, 22.0)
	foo := []int{}
	fmt.Println(foo)
	if node != nil {
		fmt.Println(node.Keys[idx])
	} else {
		fmt.Println("Not Found")
	}
}
