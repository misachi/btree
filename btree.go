package main

type Node struct {
	NumKeys  int
	Leaf     bool
	Keys     []interface{}
	Children []*Node
}

func CreateNewNode(degree int) *Node {
	return &Node{0, true, make([]interface{}, maxKeys(degree)), make([]*Node, maxChildren(degree))}
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

/*
Split full node(with 2*Degree-1 and 2*Degree children) into
2 partially full nodes
*/
func (n *Node) splitNode(degree int, idx int) {
	fullNode := n.Children[idx]
	newNode := CreateNewNode(degree)
	newNode.Leaf = fullNode.Leaf

	newNode.NumKeys = minKeys(degree)
	fullNode.NumKeys = minKeys(degree)
	for counter := 0; counter < newNode.NumKeys; counter++ {
		newNode.Keys[counter] = fullNode.Keys[counter+degree]
	}

	/* Don't bother updating child node if node is a leaf */
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

	/* Get median key from child and push it to parent node */
	n.Keys[idx] = fullNode.Keys[degree-1]
	n.NumKeys += 1
}

type BTree struct {
	Root   *Node
	Degree int
}

func (btree *BTree) insertLeaf(n *Node, key interface{}) {
	nodeIdx := n.NumKeys - 1
	for nodeIdx >= 0 && compareLessThan(n, key, nodeIdx) {
		n.Keys[nodeIdx+1] = n.Keys[nodeIdx]
		nodeIdx--
	}
	n.Keys[nodeIdx+1] = key
	n.NumKeys += 1
}

func (btree *BTree) insertNonFull(n *Node, key interface{}) {
	if n.Leaf {
		btree.insertLeaf(n, key)
	} else {
		nodeIdx := n.NumKeys - 1
		for nodeIdx >= 0 && compareLessThan(n, key, nodeIdx) {
			nodeIdx--
		}
		nodeIdx += 1
		if n.Children[nodeIdx].NumKeys == maxKeys(btree.Degree) {
			n.splitNode(btree.Degree, nodeIdx)
			if compareGreaterThan(n, key, nodeIdx) {
				nodeIdx += 1
			}
		}
		btree.insertNonFull(n.Children[nodeIdx], key)
	}
}

func (btree *BTree) Insert(key interface{}) {
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

func compareGreaterThan(n *Node, key interface{}, idx int) bool {
	switch key.(type) {
	case int:
		return key.(int) > n.Keys[idx].(int)
	case float64:
		return key.(float64) > n.Keys[idx].(float64)
	case string:
		return key.(string) > n.Keys[idx].(string)
	}
	return false
}

func compareLessThan(n *Node, key interface{}, idx int) bool {
	switch key.(type) {
	case int:
		return key.(int) < n.Keys[idx].(int)
	case float64:
		return key.(float64) < n.Keys[idx].(float64)
	case string:
		return key.(string) < n.Keys[idx].(string)
	}
	return false
}

func (btree *BTree) Search(root *Node, key interface{}) (*Node, int) {
	idx := 0
	for idx < root.NumKeys && compareGreaterThan(root, key, idx) {
		idx++
	}

	if idx >= root.NumKeys {
		idx--
	}

	if key == root.Keys[idx] {
		return root, idx
	}

	if root.Leaf {
		// We reached the last node. Key does not exist
		return nil, -1
	} else {
		if compareGreaterThan(root, key, idx) {
			idx++
		}
		return btree.Search(root.Children[idx], key)
	}
}
