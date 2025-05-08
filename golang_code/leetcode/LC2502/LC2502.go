package main

type Allocator struct {
	vis []bool
    mp 	map[int][][]int
}


func Constructor(n int) Allocator {
    return Allocator{
		vis: make([]bool, n),
		mp: make(map[int][][]int),
	}
}


func (a *Allocator) Allocate(size int, mID int) int {
    for i := 0; i < len(a.vis); {
		if a.vis[i] {
			i ++
			continue
		}

		j := i + 1
		for j < len(a.vis) && j - i < size && !a.vis[j] {
			j ++
		}
		if j - i >= size {
			if _, ok := a.mp[mID]; !ok {
				a.mp[mID] = make([][]int, 0)
			}
			a.mp[mID] = append(a.mp[mID], []int{i, j})

			for k := i; k < j; k ++ {
				a.vis[k] = true
			}
			return i
		}
		i = j
	}
	return -1
}


func (a *Allocator) FreeMemory(mID int) int {
    cnt := 0
	if v, ok := a.mp[mID]; ok {
		for _, p := range v {
			for i := p[0]; i < p[1]; i ++ {
				a.vis[i] = false
				cnt ++
			}
		}
		delete(a.mp, mID)
	}
	return cnt
}


/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
