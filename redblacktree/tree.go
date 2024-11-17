package redblacktree

type CompType[K comparable] func(x, y K) int

type Tree[K comparable, V any] struct {
	Root            *Node[K, V]
	Size            int
	Compare         CompType[K]
	AllowDuplicates bool
}

func NewTree[K comparable, V any](comparator CompType[K], allowDuplicates bool) *Tree[K, V] {
	return &Tree[K, V]{
		Root:            nil,
		Size:            0,
		Compare:         comparator,
		AllowDuplicates: allowDuplicates,
	}
}

func (tree *Tree[K, V]) IsEmpty() bool {
	return tree.Size == 0
}

func (tree *Tree[K, V]) Len() int {
	return tree.Size
}

func (tree *Tree[K, V]) Clear() {
	tree.Root = nil
	tree.Size = 0
}

func (tree *Tree[K, V]) RotateRight(node *Node[K, V]) *Node[K, V] {
	leftChild := node.Left
	leftRightChild := leftChild.Right
	parent := node.Parent

	leftChild.Parent = parent
	leftChild.Right = node

	node.Parent = leftChild
	node.Left = leftRightChild

	if parent.IsNil() {
		tree.Root = leftChild
		return leftChild
	}

	if parent.Right == node {
		parent.Right = leftChild
	} else if parent.Left == node {
		parent.Left = leftChild
	}

	return leftChild
}

func (tree *Tree[K, V]) RotateLeft(node *Node[K, V]) *Node[K, V] {
	rightChild := node.Right
	rightLeftChild := rightChild.Left
	parent := node.Parent

	rightChild.Parent = parent
	rightChild.Left = node

	node.Parent = rightChild
	node.Right = rightLeftChild

	if parent.IsNil() {
		tree.Root = rightChild
		return rightChild
	}

	if !parent.IsNil() {
		if parent.Right == node {
			parent.Right = rightChild
		} else if parent.Left == node {
			parent.Left = rightChild
		}
	}

	return rightChild
}

func (tree *Tree[K, V]) Insert(key K, value V) {
	nodeToInsert := NewNode(key, value, Red)
	tree.Size += 1

	if tree.Root.IsNil() {
		tree.Root = nodeToInsert
		tree.FixInsert(nodeToInsert)
		return
	}

	curNode := tree.Root

	for !curNode.IsNil() {
		compVal := tree.Compare(key, curNode.Key)

		if compVal == 0 {
			if !tree.AllowDuplicates {
				curNode.Key = key
				curNode.Value = value
				return
			}
		}

		if compVal < 0 {
			if curNode.Left.IsNil() {
				nodeToInsert.Parent = curNode
				curNode.Left = nodeToInsert
				tree.FixInsert(nodeToInsert)
				return
			}

			curNode = curNode.Left
		} else {
			if curNode.Right.IsNil() {
				nodeToInsert.Parent = curNode
				curNode.Right = nodeToInsert
				tree.FixInsert(nodeToInsert)
				return
			}

			curNode = curNode.Right
		}
	}
}

func (tree *Tree[K, V]) FixInsert(node *Node[K, V]) {
	if node.Parent.IsNil() {
		node.Color = Black
		return
	}

	if node.Parent.GetColor() == Black {
		return
	}

	uncle := node.GetUncle()

	if uncle.GetColor() == Red {
		node.Parent.Color = Black
		uncle.Color = Black
		grandPar := node.GetGrandParent()
		grandPar.Color = Red
		tree.FixInsert(grandPar)
		return
	}

	parent := node.Parent
	grandPar := node.GetGrandParent()

	if parent == grandPar.Left && node == parent.Right {
		parent = tree.RotateLeft(parent)
		node = parent.Left
	} else if parent == grandPar.Right && node == parent.Left {
		parent = tree.RotateRight(parent)
		node = parent.Right
	}

	node.Parent.Color = Black
	grandPar.Color = Red

	if parent == grandPar.Left && node == parent.Left {
		tree.RotateRight(grandPar)
	} else if parent == grandPar.Right && node == parent.Right {
		tree.RotateLeft(grandPar)
	}
}

func (tree *Tree[K, V]) ReplaceNodeKV(node1, node2 *Node[K, V]) {
	node1.Key = node2.Key
	node1.Value = node2.Value
}

func (tree *Tree[K, V]) FixDoubleBlack(node *Node[K, V]) {
	if node.Parent.IsNil() {
		return
	}

	sibling := node.GetSibling()

	if sibling.GetColor() == Black && sibling.AreChildrenBlack() {
		sibling.Color = Red

		if node.Parent.GetColor() == Black {
			tree.FixDoubleBlack(node.Parent)
		}

		node.Parent.Color = Black
		return
	}

	if sibling.GetColor() == Red {
		col := node.Parent.GetColor()
		node.Parent.Color = sibling.GetColor()
		sibling.Color = col

		if node == node.Parent.Left {
			tree.RotateLeft(node.Parent)
		} else if node == node.Parent.Right {
			tree.RotateRight(node.Parent)
		}

		tree.FixDoubleBlack(node)
		return
	}

	siblingLeftChild := sibling.Left

	if siblingLeftChild.GetColor() == Red {
		col := sibling.GetColor()
		sibling.Color = siblingLeftChild.GetColor()
		siblingLeftChild.Color = col

		if node == node.Parent.Left {
			tree.RotateRight(sibling)
		} else if node == node.Parent.Right {
			tree.RotateLeft(sibling)
		}
	}

	sibling = node.GetSibling()
	siblingRightChild := sibling.Right

	if siblingRightChild.GetColor() == Red {
		col := sibling.GetColor()
		sibling.Color = node.Parent.GetColor()
		node.Parent.Color = col
		siblingRightChild.Color = Black

		if node == node.Parent.Left {
			tree.RotateLeft(node.Parent)
		} else if node == node.Parent.Right {
			tree.RotateRight(node.Parent)
		}
	}

	node.Color = Black
}

func (tree *Tree[K, V]) DeleteNode(node *Node[K, V]) {
	if node.Parent.IsNil() {
		tree.Root = nil
		node = nil
		return
	}

	if node == node.Parent.Left {
		node.Parent.Left = nil
	} else if node == node.Parent.Right {
		node.Parent.Right = nil
	}

	node = nil
}

func (tree *Tree[K, V]) Delete(node *Node[K, V]) {
	if node.GetColor() == Black {
		tree.FixDoubleBlack(node)
	}

	tree.Size -= 1
	tree.DeleteNode(node)
}

func (tree *Tree[K, V]) Find(key K) (*Node[K, V], bool) {
	curNode := tree.Root

	for !curNode.IsNil() {
		compVal := tree.Compare(key, curNode.Key)

		switch {
		case compVal == 0:
			return curNode, true
		case compVal < 0:
			curNode = curNode.Left
		case compVal > 0:
			curNode = curNode.Right
		}
	}

	tree.Delete(curNode)
	return nil, false
}

func (tree *Tree[K, V]) Erase(key K) bool {
	node, found := tree.Find(key)

	if !found {
		return false
	}

	for !node.IsLeaf() {
		pre := node.GetPredeccessor()
		suc := node.GetSuccessor()

		preVal := pre.SubTreeSize()
		sucVal := suc.SubTreeSize()

		if preVal < sucVal {
			tree.ReplaceNodeKV(node, suc)
			node = suc
		} else {
			tree.ReplaceNodeKV(node, pre)
			node = pre
		}
	}

	tree.Delete(node)
	return true
}
