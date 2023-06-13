// 用递归来实现
package main

import "fmt"

// 二叉树可以使用链表来实现
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右子树
}

// 先序循环
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 根 左 右
	fmt.Print(tree.Data, " ")
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// 中序循环
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 左 根  右
	MidOrder(tree.Left)
	fmt.Print(tree.Data, " ")
	MidOrder(tree.Right)
}

// 后序循环
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 左 右 根
	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Print(tree.Data, " ")
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("先序排列：")
	PreOrder(t)
	fmt.Println("\n中序排列：")
	MidOrder(t)
	fmt.Println("\n后序排列：")
	PostOrder(t)
}
