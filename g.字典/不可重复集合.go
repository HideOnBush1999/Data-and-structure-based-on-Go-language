package main

import (
	"fmt"
	"sync"
)

type Set struct {
	m map[int]struct{} // 用字典来实现，因为字段键不能重复
	// 字典的值用空的结构体来表示，空的结构体是不占地址的
	len          int // 集合大小
	sync.RWMutex     // 锁，实现并发安全
}

// 新建一个空集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// 添加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 实际往字典添加这个键
	// 元素作为字典的键，会自动去重。所以集合个数不一定是+1
	s.len = len(s.m) // 重新计算元素数量
}

// 删除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()

	if s.len == 0 {
		return
	}
	delete(s.m, item) // 实际从字典删除这个键
	s.len = len(s.m)  // 重新计算元素数量
}

// 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// 集合是否为空
func (s *Set) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

// 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{} // 字典重新赋值
	s.len = 0                // 大小归零
}

// 将集合转化为列表
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func main() {
	s := NewSet(5)
	s.Add(1)
	s.Add(1)

	s.Add(2)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 does exist")
	}
	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}

/*
s.Lock()获取写锁，用于互斥访问和修改Set的内部状态。
s.RLock()获取读锁，用于支持并发读取Set的内部状态，不涉及修改操作。
*/
