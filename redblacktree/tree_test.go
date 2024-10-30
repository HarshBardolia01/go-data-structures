package redblacktree

import (
	"fmt"
	"testing"
)

func DFS(node *Node[int, int], path string) {
	if node.IsNil() {
		return
	}

	DFS(node.Left, path+"L")

	fmt.Printf("Key: %d\n", node.Key)
	fmt.Printf("Color: %s\n", node.Color)
	fmt.Printf("Path: %s\n\n", path)

	DFS(node.Right, path+"R")
}

func TestTreeDelete(t *testing.T) {
	t.Run("tc1", func(t *testing.T) {
		tree := NewTree[int, int](func(x, y int) int {
			return x - y
		})

		tree.Insert(10, 0)

		tree.Root.Left = NewNode(5, 0, Red)
		tree.Root.Left.Parent = tree.Root

		tree.Root.Left.Left = NewNode(2, 0, Black)
		tree.Root.Left.Left.Parent = tree.Root.Left

		tree.Root.Left.Right = NewNode(9, 0, Black)
		tree.Root.Left.Right.Parent = tree.Root.Left

		tree.Root.Right = NewNode(30, 0, Red)
		tree.Root.Right.Parent = tree.Root

		tree.Root.Right.Left = NewNode(25, 0, Black)
		tree.Root.Right.Left.Parent = tree.Root.Right

		tree.Root.Right.Right = NewNode(40, 0, Black)
		tree.Root.Right.Right.Parent = tree.Root.Right

		tree.Root.Right.Right.Left = NewNode(38, 0, Red)
		tree.Root.Right.Right.Left.Parent = tree.Root.Right.Right

		DFS(tree.Root, "0")
		tree.Erase(30)
		fmt.Println("---------------------")
		DFS(tree.Root, "0")
	})

	t.Run("tc2", func(t *testing.T) {
		tree := NewTree[int, int](func(x, y int) int {
			return x - y
		})

		tree.Insert(10, 0)

		tree.Root.Left = NewNode(5, 0, Black)
		tree.Root.Left.Parent = tree.Root

		tree.Root.Right = NewNode(20, 0, Red)
		tree.Root.Right.Parent = tree.Root

		tree.Root.Right.Left = NewNode(15, 0, Black)
		tree.Root.Right.Left.Parent = tree.Root.Right

		tree.Root.Right.Right = NewNode(30, 0, Black)
		tree.Root.Right.Right.Parent = tree.Root.Right

		DFS(tree.Root, "0")
		tree.Erase(15)
		fmt.Println("---------------------")
		DFS(tree.Root, "0")
	})

	t.Run("tc3", func(t *testing.T) {
		tree := NewTree[int, int](func(x, y int) int {
			return x - y
		})

		tree.Insert(10, 0)

		tree.Root.Left = NewNode(5, 0, Black)
		tree.Root.Left.Parent = tree.Root

		tree.Root.Right = NewNode(20, 0, Black)
		tree.Root.Right.Parent = tree.Root

		tree.Root.Left.Left = NewNode(1, 0, Black)
		tree.Root.Left.Left.Parent = tree.Root.Left

		tree.Root.Left.Right = NewNode(7, 0, Black)
		tree.Root.Left.Right.Parent = tree.Root.Left

		tree.Root.Right.Left = NewNode(15, 0, Black)
		tree.Root.Right.Left.Parent = tree.Root.Right

		tree.Root.Right.Right = NewNode(30, 0, Black)
		tree.Root.Right.Right.Parent = tree.Root.Right

		DFS(tree.Root, "0")
		tree.Erase(15)
		fmt.Println("---------------------")
		DFS(tree.Root, "0")
	})

	t.Run("tc4", func(t *testing.T) {
		tree := NewTree[int, int](func(x, y int) int {
			return x - y
		})

		tree.Insert(10, 0)

		tree.Root.Left = NewNode(5, 0, Black)
		tree.Root.Left.Parent = tree.Root

		tree.Root.Right = NewNode(20, 0, Black)
		tree.Root.Right.Parent = tree.Root

		tree.Root.Left.Left = NewNode(1, 0, Black)
		tree.Root.Left.Left.Parent = tree.Root.Left

		tree.Root.Left.Right = NewNode(7, 0, Black)
		tree.Root.Left.Right.Parent = tree.Root.Left

		tree.Root.Right.Left = NewNode(15, 0, Black)
		tree.Root.Right.Left.Parent = tree.Root.Right

		tree.Root.Right.Right = NewNode(30, 0, Red)
		tree.Root.Right.Right.Parent = tree.Root.Right

		tree.Root.Right.Right.Left = NewNode(25, 0, Black)
		tree.Root.Right.Right.Left.Parent = tree.Root.Right.Right

		tree.Root.Right.Right.Right = NewNode(40, 0, Black)
		tree.Root.Right.Right.Right.Parent = tree.Root.Right.Right

		DFS(tree.Root, "0")
		tree.Erase(15)
		fmt.Println("---------------------")
		DFS(tree.Root, "0")
	})

	t.Run("tc5", func(t *testing.T) {
		tree := NewTree[int, int](func(x, y int) int {
			return x - y
		})

		tree.Insert(10, 0)

		tree.Root.Left = NewNode(5, 0, Black)
		tree.Root.Left.Parent = tree.Root

		tree.Root.Right = NewNode(30, 0, Black)
		tree.Root.Right.Parent = tree.Root

		tree.Root.Left.Left = NewNode(1, 0, Black)
		tree.Root.Left.Left.Parent = tree.Root.Left

		tree.Root.Left.Right = NewNode(7, 0, Black)
		tree.Root.Left.Right.Parent = tree.Root.Left

		tree.Root.Right.Left = NewNode(25, 0, Red)
		tree.Root.Right.Left.Parent = tree.Root.Right

		tree.Root.Right.Right = NewNode(40, 0, Black)
		tree.Root.Right.Right.Parent = tree.Root.Right

		tree.Root.Right.Left.Left = NewNode(20, 0, Black)
		tree.Root.Right.Left.Left.Parent = tree.Root.Right.Left

		tree.Root.Right.Left.Right = NewNode(28, 0, Black)
		tree.Root.Right.Left.Right.Parent = tree.Root.Right.Left

		DFS(tree.Root, "0")
		tree.Erase(1)
		fmt.Println("---------------------")
		DFS(tree.Root, "0")
	})
}
