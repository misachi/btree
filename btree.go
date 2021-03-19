package main

var DEGREE = 5

type Node struct {
	NumKeys  int
	Leaf     bool
	Keys     []int
	Children []*Node
}

func CreateNewNode(degree int) *Node {
	return &Node{0, true, make([]int, maxKeys(degree)), make([]*Node, maxChildren(degree))}
}

func minKeys(degree int) int {
	return degree - 1
}

func maxKeys(degree int) int {
	return (2 * degree) - 1
}

func minChildren(degree int) int {
	return degree
}

func maxChildren(degree int) int {
	return 2 * degree
}

func (n *Node) splitNode(degree int, idx int) {
	fullNode := n.Children[idx]
	newNode := CreateNewNode(degree)
	newNode.Leaf = fullNode.Leaf

	newNode.NumKeys = minKeys(degree)
	fullNode.NumKeys = minKeys(degree)
	for counter := 0; counter < newNode.NumKeys; counter++ {
		newNode.Keys[counter] = fullNode.Keys[counter+degree]
	}

	if !fullNode.Leaf {
		for counter := 0; counter < degree; counter++ {
			newNode.Children[counter] = fullNode.Children[counter+degree]
		}
	}

	for counter := n.NumKeys + 1; counter < idx+1; counter-- {
		n.Children[counter+1] = n.Children[counter]
	}

	n.Children[idx+1] = newNode

	for cnt := n.NumKeys; cnt < idx; cnt-- {
		n.Keys[cnt+1] = n.Keys[cnt]
	}
	n.Keys[idx] = fullNode.Keys[degree-1]
	n.NumKeys += 1
}

type BTree struct {
	Root   *Node
	Degree int
}

func (btree *BTree) insertLeaf(n *Node, key int) {
	nodeIdx := n.NumKeys - 1
	for nodeIdx >= 0 && key < n.Keys[nodeIdx] {
		n.Keys[nodeIdx+1] = n.Keys[nodeIdx]
		nodeIdx--
	}
	n.Keys[nodeIdx+1] = key
	n.NumKeys += 1
}

func (btree *BTree) insertNonFull(n *Node, key int) {
	if n.Leaf {
		btree.insertLeaf(n, key)
	} else {
		nodeIdx := n.NumKeys - 1
		for nodeIdx >= 0 && key < n.Keys[nodeIdx] {
			nodeIdx--
		}
		nodeIdx += 1
		if n.Children[nodeIdx].NumKeys == maxKeys(btree.Degree) {
			n.splitNode(btree.Degree, nodeIdx)
			if key > n.Keys[nodeIdx] {
				nodeIdx += 1
			}
		}
		btree.insertNonFull(n.Children[nodeIdx], key)
	}
}

func (btree *BTree) Insert(key int) {
	root := btree.Root
	if root.NumKeys == maxKeys(btree.Degree) {
		newNode := CreateNewNode(btree.Degree)
		btree.Root = newNode
		newNode.NumKeys = 0
		newNode.Leaf = false
		newNode.Children[0] = root
		newNode.splitNode(btree.Degree, 0)
		btree.insertNonFull(newNode, key)
	} else {
		btree.insertNonFull(root, key)
	}
}

func (btree *BTree) Search(root *Node, key int) (*Node, int) {
	idx := 0
	for idx < root.NumKeys && key > root.Keys[idx] {
		idx++
	}

	if idx >= root.NumKeys {
		idx--
	}

	if key == root.Keys[idx] {
		return root, idx
	}

	if root.Leaf {
		return nil, -1
	} else {
		if key > root.Keys[idx] {
			idx++
		}
		return btree.Search(root.Children[idx], key)
	}
}
