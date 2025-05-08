package main

import "math/rand"

const P = 0.5
const MAX_LEVEL = 16
const INVALID_VAL = -1

// [0, MAX_LEVEL)
func randomLevel() int {
	lv := 0
	for rand.Float32() < P && lv + 1 < MAX_LEVEL {
		lv ++
	}
	return lv
}

type SkiplistNode struct {
	val 	int
	next	[]*SkiplistNode
}

type Skiplist struct {
	head 	*SkiplistNode
	level 	int
}

func Constructor() Skiplist {
	return Skiplist{head: &SkiplistNode{INVALID_VAL, make([]*SkiplistNode, MAX_LEVEL)}, level: 0}
}


func (s *Skiplist) Search(target int) bool {
	curr := s.head
	for i := s.level; i >= 0; i -- {
		for curr.next[i] != nil && curr.next[i].val < target {
			curr = curr.next[i]
		}
	}
	curr = curr.next[0]
	return curr != nil && curr.val == target
}	


func (s *Skiplist) Add(num int)  {
	curr := s.head
	last := make([]*SkiplistNode, MAX_LEVEL)
	for i := range last {
		last[i] = s.head
	}
	for i := s.level; i >= 0; i -- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		last[i] = curr
	}
	newNode := SkiplistNode{val: num, next: make([]*SkiplistNode, MAX_LEVEL)}
	lv := randomLevel()
	for i, node := range last[:lv + 1] {
		newNode.next[i] = node.next[i]
		node.next[i] = &newNode
	}
	s.level = max(s.level, lv)
}


func (s *Skiplist) Erase(num int) bool {
	curr := s.head
	last := make([]*SkiplistNode, MAX_LEVEL)
	for i := s.level; i >= 0; i -- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		last[i] = curr
	}
	curr = curr.next[0]
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i <= s.level && last[i].next[i] == curr; i ++ {
		last[i].next[i] = curr.next[i]
	}
	for s.level > 0 && s.head.next[s.level] == nil {
		s.level --
	}
	return true
}