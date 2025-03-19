package main

import (
	"math"
	"sync"
)

type lru struct {
	mu       sync.Mutex
	timestmp int
	size     int
	mp       map[string]string
	vis      map[string]int
}

func constructor(size int) lru {
	return lru{
		timestmp: 0,
		size:     size,
		mp:       make(map[string]string),
		vis:      make(map[string]int),
	}
}

func (u lru) get(k string) (string, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.mp[k]; !ok {
		return "", false
	}

	u.timestmp++
	u.vis[k] = u.timestmp
	return u.mp[k], true
}

func (u lru) set(k, v string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.mp[k]; !ok && len(u.mp) >= u.size {
		minv := math.MaxInt
		var key string
		for k, v := range u.vis {
			if minv > v {
				minv = v
				key = k
			}
		}
		delete(u.mp, key)
	}
	u.timestmp++
	u.vis[k] = u.timestmp
	u.mp[k] = v
}
