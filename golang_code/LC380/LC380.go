package main

import "math/rand"

type RandomizedSet struct {
	mp		map[int]int
	nums	[]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
		mp: make(map[int]int),
		nums: make([]int, 0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.mp[val]; !ok {
		this.mp[val] = len(this.nums)
		this.nums = append(this.nums, val)
		return true
	}
	return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if idx1, ok := this.mp[val]; ok {
		idx2 := len(this.nums) - 1
		x, y := this.nums[idx1], this.nums[idx2]
		this.nums[idx1], this.nums[idx2] = y, x
		this.nums = this.nums[:idx2]
		this.mp[y] = idx1 
		delete(this.mp, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {

}