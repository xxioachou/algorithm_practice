package main

func minSwaps(s string) int {
	cnt := 0
	r := 0
	for _, ch := range s {
		if ch == '[' {
			cnt ++
		} else {
			if cnt > 0 {
				cnt --
			} else {
				r ++
			}
		}
	}
	return (r + 1) / 2
}
