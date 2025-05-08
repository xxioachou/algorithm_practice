package main

func evenOddBit(n int) []int {
	even, odd := 0, 0
	for i := 0; n > 0; n >>= 1 {
		if n % 2 == 1 {
			if i % 2 == 0 {
				even ++
			} else {
				odd ++
			}
		}
		i ++
	}
	return []int{even, odd}
}
