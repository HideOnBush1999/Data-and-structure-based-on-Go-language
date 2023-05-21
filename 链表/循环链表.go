package main

import "fmt"

// 循环链表
type Ring struct {
	prev, next *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

// 创建 N 个节点的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring) // 头节点
	p := r         // 当前节点
	for i := 1; i < n; i++ {
		// 创建一个前驱节点为当前节点的节点
		// 然后当前节点的后驱节点指向这个创建的节点
		p.next = &Ring{prev: p}
		p = p.next // 当前节点变成原本当前节点的后驱节点
	}
	p.next = r // 尾节点的后驱节点指向首节点
	r.prev = p // 头节点的前驱节点指向尾节点
	return r
}

// 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// 获取当前节点的下 n 个节点,当 n 为负数，表示从前面往前遍历，否则往后面遍历
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 添加节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// 获取链表长度
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: "hello"})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		node = node.Next()
		if node == r {
			return
		}
	}
}

func main() {
	r := new(Ring)
	r.init()
	linkNewTest()
}
