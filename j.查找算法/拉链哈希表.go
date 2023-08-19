/*
1. 初始化：新建一个 2^x 个长度的数组，一开始 x 较小。
2. 添加键值：进行 hash(key) & (2^x-1)，定位到数组下标，查找数组下标对应的链表，如果链表有该键，更新其值，否则追加元素。
3. 获取键值：进行 hash(key) & (2^x-1)，定位到数组下标，查找数组下标对应的链表，如果链表不存在该键，返回 false，否则返回该值以及 true。
4. 删除键值：进行 hash(key) & (2^x-1)，定位到数组下标，查找数组下标对应的链表，如果链表不存在该键，直接返回，否则删除该键。
5. 进行键值增删时如果数组容量太大或者太小，需要相应缩容或扩容。

哈希表查找，是一种用空间换时间的查找算法，时间复杂度能达到：O(1)，最坏情况下退化到查找链表：O(n)。但均匀性很好的哈希算法以及合适空间大小的数组，在很大概率避免了最坏情况。
哈希表在添加元素时会进行伸缩，会造成较大的性能消耗，所以有时候会用到其他的查找算法：树查找算法。
(这里的算法可以模拟map映射的实现，大致差不太多)
*/

package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/OneOfOne/xxhash"
)

const (
	expandFactor = 0.75 // 扩容因子
)

// 键值对，连成一个链表
type keyPairs struct {
	key   string      // 键
	value interface{} // 值
	next  *keyPairs   // 下一个键值对
}

// 哈希表
type HashMap struct {
	array        []*keyPairs // 哈希表数组，每个元素是一个键值对
	capacity     int         // 数组容量
	len          int         // 已添加键值对元素数量
	capacityMask int         // 掩码，等于 capacity - 1
	lock         sync.Mutex
}

// -------------------------------------------------------------
// 创建大小为 capacity 的哈希表
func NewHashMap(capacity int) *HashMap {
	defaultCapacity := 1 << 4 // 默认大小为 16
	if capacity <= defaultCapacity {
		// 传入的大小小于默认大小，则使用默认大小
		capacity = defaultCapacity
	} else {
		// 否则，实际大小为大于 capacity 的第一个 2^k
		capacity = 1 << (int(math.Ceil(math.Log2(float64(capacity)))))
	}

	// 新建一个哈希表
	m := new(HashMap)
	m.array = make([]*keyPairs, capacity)
	m.capacity = capacity
	m.capacityMask = capacity - 1
	return m
}

// 返回哈希表已添加元素数量
func (m *HashMap) Len() int {
	return m.len
}

// 返回哈希表的容量
func (m *HashMap) Capacity() int {
	return m.capacity
}

// 求 key 的哈希值 (将函数赋值给变量的好处是可以将函数作为值进行传递和操作，类似于其他变量。)
var hashAlgorithm = func(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

// 对键进行哈希求值，并计算下标
func (m *HashMap) hashIndex(key string) int {
	hash := hashAlgorithm([]byte(key))     // 求哈希
	index := hash & uint64(m.capacityMask) // 求下标
	return int(index)
}

// 添加键值对（扩容需要一个个从新赋值，会计算量比较大）
func (m *HashMap) Put(key string, value interface{}) {
	// 实现并发安全
	m.lock.Lock()
	defer m.lock.Unlock()

	// 键值对要放的哈希表数组下标
	index := m.hashIndex(key)

	// 哈希表数组下标的元素
	element := m.array[index]

	// 元素为空，没有哈希冲突，直接赋值
	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		var lastPairs *keyPairs
		// 存在哈希冲突，遍历链表查看元素是否存在，存在则替换值
		for element != nil {
			if element.key == key {
				element.value = value
				return
			}
			lastPairs = element
			element = element.next
		}
		// 要是运行到这里， lastPairs 就是链表中的最后一个节点
		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}

	// 新的哈希表数量
	newLen := m.len + 1

	// 如果超出扩容因子，需要扩建
	if float64(newLen)/float64(m.capacity) >= expandFactor {
		// 新建一个原来两倍大小的哈希表
		newM := new(HashMap)
		newM.array = make([]*keyPairs, 2*m.capacity)
		newM.capacity = 2 * m.capacity
		newM.capacityMask = 2*m.capacity - 1

		// 遍历老的哈希表，将键值对重新哈希到新哈希表
		for _, pairs := range m.array { // pairs 类型为keyPairs的指针
			for pairs != nil {
				// 直接递归
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		// 替换老的哈希表
		m.array = newM.array
		m.capacity = newM.capacity
		m.capacityMask = newM.capacityMask
	}
	m.len = newLen
}

// 获取键值对（只是读取的话，可以不加锁）
func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	index := m.hashIndex(key)
	element := m.array[index]

	// 因为可能存在哈希冲突，所以还需要再用 key 确定
	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.next
	}
	return // 因为函数返回值中包含变量名，所以这时会返回默认初始值 nil 和 false
}

// 删除键值对
func (m *HashMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key)
	element := m.array[index]

	// 情况1： 空链表，不用删除，直接返回
	if element == nil {
		return
	}

	// 情况2： 链表的第一个元素就是要删除的元素
	if element.key == key {
		m.array[index] = element.next // 让链表的头部变为下一个元素
		m.len = m.len - 1
		return
	}

	// 情况3： 删除的不是链表的第一个元素
	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next // 这样就跳过了 nextElement 节点
			m.len = m.len - 1
			return
		}
		// 没有匹配上，向后推进
		element = nextElement
		nextElement = nextElement.next
	}
}

// 遍历打印哈希表
func (m *HashMap) Range() {
	for _, pairs := range m.array {
		for pairs != nil { // 当这个地方没有存入元素或者是链表的最后一个节点时，pairs 为 nil
			fmt.Printf("%v : %v  |  ", pairs.key, pairs.value)
			pairs = pairs.next
		}
	}
	fmt.Println()
}
/*
func main() { 
	// 新建一个哈希表
	hashMap := NewHashMap(16)

	// 放35个值
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i)) // Sprintf()也提供了一种数字转字符串的方法，而且还可以再修改
	}
	fmt.Println("cap:", hashMap.Capacity(), "len:", hashMap.Len())

	// 打印全部键值对
	hashMap.Range() // 因为一开始通过 hash 函数存的时候，就是无需。所以当你从头开始遍历的时候，打印出来的也是无需的。

	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}

	// 删除键
	hashMap.Delete(key)
	fmt.Println("after delete cap:", hashMap.Capacity(), "len:", hashMap.Len())
	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
}
 */