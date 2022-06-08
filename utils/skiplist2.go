package utils

import "math/rand"

const (
	MaxLevel = 32
)
type SkipListKV struct {
	Key   string
	Value string
}

type SkipList2 struct {
	Head, Tail *SkipListElement
}

type SkipListElement struct {
	Data   SkipListKV
	levels []*SkipListElement
}

func (s *SkipList2) add(k, v string) {
	ele := &SkipListElement{Data: SkipListKV{Key: k, Value: v}, levels: make([]*SkipListElement, randLevel())}
	prevEleList := make([]*SkipListElement, len(ele.levels))
	l := len(s.Head.levels)
	for i := l-1; i >= 0; i-- {
		var prev, curr *SkipListElement
		for curr = s.Head.levels[i]; curr != nil; {
			if curr.Data.Key == k {
				return
			}
			if curr.Data.Key > k {
				break
			}
			curr = prev.levels[i]
		}
		if i < len(prevEleList) {
			prevEleList[i] = prev
		}
	}

	for i := 0; i < len(prevEleList); i++ {
		ele.levels[i] = prevEleList[i].levels[i]
		prevEleList[i].levels[i] = ele
	}
}

func (s *SkipList2) get(k string) *SkipListElement {
	l := len(s.Head.levels)
	var prev, curr *SkipListElement

	for i := l - 1; i >= 0; i-- {
		for curr = s.Head.levels[i]; curr != nil; curr = prev.levels[i] {
			if curr.Data.Key == k {
				return curr.levels[0]
			}
			if curr.Data.Key > k {
				break
			}
			prev = curr
		}
	}
	return nil
}

func (s *SkipList2) del(k string) {

}

func randLevel() int {
	i := 1
	for ; i < MaxLevel;i++ {
		if rand.Intn(2) == 0 {
			 break
		}
	}
	return i
}