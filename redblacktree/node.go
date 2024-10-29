package redblacktree

type Color bool

const (
	Black Color = true
	Red   Color = false
)

func (c Color) String() string {
	if c {
		return "Black"
	}

	return "Red"
}

type Node[K comparable, V any] struct {
	Key    K
	Value  V
	Color  Color
	Left   *Node[K, V]
	Right  *Node[K, V]
	Parent *Node[K, V]
}

func NewNode[K comparable, V any](key K, value V, color Color) *Node[K, V] {
	return &Node[K, V]{
		Key:    key,
		Value:  value,
		Color:  color,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}

func (node *Node[K, V]) IsNil() bool {
	return node == nil
}

func (node *Node[K, V]) IsLeaf() bool {
	if node.IsNil() {
		return false
	}

	return node.Left.IsNil() && node.Right.IsNil()
}

func (node *Node[K, V]) GetKey() K {
	return node.Key
}

func (node *Node[K, V]) GetValue() V {
	return node.Value
}

func (node *Node[K, V]) GetLeftmostNode() *Node[K, V] {
	if node.IsNil() {
		return nil
	}

	for !node.Left.IsNil() {
		node = node.Left
	}

	return node
}

func (node *Node[K, V]) GetRightmostNode() *Node[K, V] {
	if node.IsNil() {
		return nil
	}

	for !node.Right.IsNil() {
		node = node.Right
	}

	return node
}

func (node *Node[K, V]) GetSuccessor() *Node[K, V] {
	if node.IsNil() {
		return nil
	}

	return node.Right.GetLeftmostNode()
}

func (node *Node[K, V]) GetPredeccessor() *Node[K, V] {
	if node.IsNil() {
		return nil
	}

	return node.Left.GetRightmostNode()
}

func (node *Node[K, V]) GetGrandParent() *Node[K, V] {
	if node.IsNil() || node.Parent.IsNil() {
		return nil
	}

	return node.Parent.Parent
}

func (node *Node[K, V]) GetSibling() *Node[K, V] {
	if node.IsNil() || node.Parent.IsNil() {
		return nil
	}

	if node == node.Parent.Left {
		return node.Parent.Right
	}

	return node.Parent.Left
}

func (node *Node[K, V]) GetUncle() *Node[K, V] {
	if node.IsNil() || node.Parent.IsNil() || node.Parent.Parent.IsNil() {
		return nil
	}

	return node.Parent.GetSibling()
}

func (node *Node[K, V]) SubTreeSize() int {
	if node.IsNil() {
		return 0
	}

	return 1 + node.Left.SubTreeSize() + node.Right.SubTreeSize()
}

func (node *Node[K, V]) GetColor() Color {
	if node.IsNil() {
		return Black
	}

	return node.Color
}
