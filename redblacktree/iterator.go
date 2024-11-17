package redblacktree

type Iterator[K comparable, V any] struct {
	tree  *Tree[K, V]
	node  *Node[K, V]
	index int
}

func (tree *Tree[K, V]) Begin() *Iterator[K, V] {
	return &Iterator[K, V]{
		tree:  tree,
		node:  tree.Root.GetLeftmostNode(),
		index: 0,
	}
}

func (tree *Tree[K, V]) End() *Iterator[K, V] {
	return &Iterator[K, V]{
		tree:  tree,
		node:  nil,
		index: tree.Size,
	}
}

func (tree *Tree[K, V]) Iterator() *Iterator[K, V] {
	return tree.Begin()
}

func (it *Iterator[K, V]) GetNext() *Iterator[K, V] {
	if it == nil || it.node.IsNil() {
		return nil
	}

	nxtNode := it.node.GetSuccessor()

	if nxtNode != nil {
		return &Iterator[K, V]{
			tree:  it.tree,
			node:  nxtNode,
			index: it.index + 1,
		}
	}

	curNode := it.node

	for !curNode.Parent.IsNil() && it.tree.Compare(curNode.Parent.Key, it.node.Key) <= 0 {
		curNode = curNode.Parent
	}

	if it.tree.Compare(curNode.Key, it.node.Key) <= 0 {
		return it.tree.End()
	}

	return &Iterator[K, V]{
		tree:  it.tree,
		node:  curNode,
		index: it.index + 1,
	}
}
