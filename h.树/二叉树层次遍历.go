package main

import (
	"fmt"
	"sync"
)

// 二叉树节点
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右子树
}

// 队列节点
type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode // 队列节点存储的值是二叉树节点
}

// 链表
type LinkQueue struct {
	root *LinkNode  // 链表起点
	size int        // 队列的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 层次遍历
func LayerOrder(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}

	// 新建队列
	queue := new(LinkQueue)
	// 根节点先入队
	queue.Add(treeNode)
	for queue.size > 0 {
		// 不断出队列
		element := queue.Remove()
		// 先打印节点值
		fmt.Print(element.Data, " ")

		// 左子树非空，入队列
		if element.Left != nil {
			queue.Add(element.Left)
		}

		// 右子树非空，入队列
		if element.Right != nil {
			queue.Add(element.Right)
		}

	}

}

// 入队
func (queue *LinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		newNode := new(LinkNode)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		// 新节点放在链表尾部
		nowNode.Next = newNode
	}
	queue.size = queue.size + 1
}

// 出队  (弹出的是二叉树节点)
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("over limit")
	}

	v := queue.root.Value
	queue.root = queue.root.Next
	queue.size = queue.size - 1
	return v
}

// 队列中元素数量
func (queue *LinkQueue) Size() int {
	return queue.size
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}
	fmt.Println("\n层次排序")
	LayerOrder(t)
}
