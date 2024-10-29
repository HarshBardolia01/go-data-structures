package redblacktree

type CompType[K comparable] func(x, y K) int

type Tree[K comparable, V any] struct {
	Root    *Node[K, V]
	Size    int
	Compare CompType[K]
}

func NewTree[K comparable, V any](comparator CompType[K]) *Tree[K, V] {
	return &Tree[K, V]{
		Root:    nil,
		Size:    0,
		Compare: comparator,
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

	if tree.Root.IsNil() {
		tree.Root = nodeToInsert
		tree.Size += 1
		tree.FixInsert(nodeToInsert)
		return
	}

	curNode := tree.Root

	for !curNode.IsNil() {
		compVal := tree.Compare(key, curNode.Key)

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
