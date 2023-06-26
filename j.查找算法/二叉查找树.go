package main

import "fmt"

// 二叉查找树
type BinarySearchTree struct {
	Root *BinarySearchTreeNode // 树根节点
}

// 二叉查找树节点
type BinarySearchTreeNode struct {
	Value int64                 // 值
	Times int64                 // 值重复的次数
	Left  *BinarySearchTreeNode // 左子树
	Right *BinarySearchTreeNode // 右子树
}

// 初始化一个二叉查找树
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

// 添加元素
func (tree *BinarySearchTree) Add(value int64) {
	// 如果没有树根，证明是颗空树，添加树根后返回
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}

	// 将值添加进去
	tree.Root.Add(value)
}

func (node *BinarySearchTreeNode) Add(value int64) {
	// 如果插入的值比节点的值小，那么要插入到该节点的左子树
	if value < node.Value {
		if node.Left == nil { // 左子树为空直接添加
			node.Left = &BinarySearchTreeNode{Value: value}
		} else { // 否则递归
			node.Left.Add(value)
		}
	} else if value > node.Value { // 如果插入的值比节点的值大，那么要插入到该节点的右子树
		if node.Right == nil { // 右子树为空直接添加
			node.Right = &BinarySearchTreeNode{Value: value}
		} else { // 否则递归
			node.Right.Add(value)
		}
	} else { // 值相同
		node.Times++
	}
}

// 找出最小值的节点
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil // 如果是空树，返回空
	}

	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// 找出最大值的节点
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil // 如果是空树，返回空
	}

	return tree.Root.FindMaxValue()
}

func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// 查找指定元素
func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil // 如果是空树，返回空
	}
	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {
	if value > node.Value { // 大于节点值
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(value) // 递归
	} else if value < node.Value { // 小于节点值
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(value) // 递归
	} else { // 等于节点值
		return node
	}
}

// 查找指定元素的父亲
func (tree *BinarySearchTree) FindParent(value int64) *BinarySearchTreeNode {
	if tree.Root == nil { // 如果是空树，返回空
		return nil
	}
	if tree.Root.Value == value { // 如果根节点等于该值，根节点其没有父节点，返回nil
		return nil
	}
	return tree.Root.FindParent(value)
}

func (node *BinarySearchTreeNode) FindParent(value int64) *BinarySearchTreeNode {
	// 外层没有值相等的判定,因为在上一个函数中已经进行了判断
	if value < node.Value { // 如果查找的值小于节点值，从节点的左子树开始找
		leftTree := node.Left
		if leftTree == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		// 左子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value) // 递归接着找
		}
	} else { // 如果查找的值大于节点值，从节点的右子树开始找
		rightTree := node.Right
		if rightTree == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		// 右子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value) // 递归接着找
		}
	}

}

// 删除指定的元素
func (tree *BinarySearchTree) Delete(value int64) {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}
	// 查找该值是否存在
	node := tree.Root.Find(value)
	if node == nil {
		// 不存在该值，直接返回
		return
	}
	// 查找该值的父亲节点
	parent := tree.Root.FindParent(value)
	// 第一种情况，删除的是根节点，且根节点没有儿子
	if parent == nil && node.Left == nil && node.Right == nil {
		// 置空后直接返回
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		// 第二种情况，删除的节点有父亲节点，但没有子树
		// 如果删除的是节点是父亲的左儿子，直接将该值删除即可
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			// 删除的原来是父亲的右儿子，直接将该值删除即可
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		// 第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点。
		// 右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到，替换后二叉查找树的性质又满足了。
		// 找右子树中最小的值，一直往右子树的左边找
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// 把最小的节点删掉
		tree.Delete(minNode.Value)
		// 最小值的节点替换被删除节点
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		// 第四种情况，只有一个子树，那么该子树直接替换被删除的节点即可
		// 父亲为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		// 左子树不为空
		if node.Left != nil {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的左子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}

// 中序遍历（使用二叉查找树可以实现排序，只需要对树进行中序遍历即可）
func (tree *BinarySearchTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *BinarySearchTreeNode) MidOrder() {
	if node == nil {
		return
	}
	// 先打印左子树
	node.Left.MidOrder()

	// 按照次数打印根节点
	for i := 0; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// 打印右子树
	node.Right.MidOrder()
}

func main() {
	values := []int64{3, 6, 8, 20, 9, 2, 6, 8, 9, 3, 5, 40, 7, 9, 13, 6, 8}
	// 初始化二叉查找树并添加元素
	tree := NewBinarySearchTree()
	for _, v := range values {
		tree.Add(v)
	}
	// 找到最大值或最小值的节点
	fmt.Println("find min value:", tree.FindMinValue())
	fmt.Println("find max value:", tree.FindMaxValue())
	// 查找不存在的99
	node := tree.Find(99)
	if node != nil {
		fmt.Println("find it 99!")
	} else {
		fmt.Println("not find it 99!")
	}
	// 查找存在的9
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}
	// 删除存在的9后，再查找9
	tree.Delete(9)
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}
	// 中序遍历，实现排序
	tree.MidOrder()
}
