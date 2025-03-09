package main

func minAnagramLength(s string) int {
	n := len(s)
	cmp := func(a, b []int) int {
		eq := true
		for i := range(a) {
			if a[i] > b[i] {
				return 1
			}
			if a[i] != b[i] {
				eq = false
			}
		}
		if eq {
			return 0
		}
		return -1
	}
	check := func(Len int) bool {
		a := make([]int, 26)
		b := make([]int, 26)
		for i := 0; i < Len; i++ {
			a[int(s[i] - 'a')] ++
		}
		for i := Len; i < n; i += Len {
			for k := range(b) {
				b[k] = 0
			}
			for j := i; j < i + Len; j++ {
				b[int(s[j] - 'a')] ++
			}
			if cmp(b, a) != 0 {
				return false
			}
		}
		return true
	}
	for l := 1; l <= n; l++ {
		if n % l != 0 {
			continue
		}
		if check(l) {
			return l
		}
	}
	return n
} 

func main() {

}