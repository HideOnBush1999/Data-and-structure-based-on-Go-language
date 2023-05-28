package main

import "fmt"

// 定义节点结构体
type LinkNode struct {
	Data     int64     // 数据
	NextNode *LinkNode // 下一个节点的地址
}

//
func main() {
	// 新的节点
	node := new(LinkNode)
	node.Data = 2

	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1 // 将 node1 链接到 node 节点后面

	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2 // 将 node2 链接到 node1 节点后面

	// 按照顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
			continue
		}
		break
	}
}
