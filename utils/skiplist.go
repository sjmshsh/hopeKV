package utils

import (
	"github.com/sjmshsh/hopeKV/utils/codec"
	"math/rand"
	"sync"
)

const (
	defaultMaxHeight = 48
)

type SkipList struct {
	header *Element

	rand *rand.Rand

	maxLevel int
	length   int
	lock     sync.RWMutex
	size     int64
}

type Element struct {
	// levels[i] 存的是这个节点的第 i 个 level 的下一个节点
	levels []*Element
	entry  *codec.Entry
	score  float64
}

func (list *SkipList) randLevel() int {
	// 有 1/2 的几率返回 1
	// 有 1/4 的几率返回 2
	// 有 1/8 的几率返回 3
	// 直到最大层
	for i := 0; i < list.maxLevel; i++ {
		if list.rand.Intn(2) == 0 {
			return i
		}
	}

	return list.maxLevel
}

func (list *SkipList) Size() int64 {
	return list.size
}
