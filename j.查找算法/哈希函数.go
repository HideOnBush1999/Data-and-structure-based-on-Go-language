package main

import (
	"fmt"

	"github.com/OneOfOne/xxhash"
)

// 将一个键进行Hash
func XXHash(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

func main() {
	keys := []string{"hi", "my", "friend", "I", "love", "you", "my", "apple"}
	for _, key := range keys {
		// []byte(key) 是将字符串 key 转换为字节切片的操作
		fmt.Printf("xxhash('%s')=%d\n", key, XXHash([]byte(key))) // 对于对应的字符串，每次计算出来的值是固定不变的
	}
}
