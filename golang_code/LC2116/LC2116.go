package main

func canBeValid(s string, locked string) bool {
    a := make([]int, 0)
	l := make([]int, 0)
    for i, x := range s {
		if locked[i] == '1' {
			if x == '(' {
				l = append(l, i)
			} else {
				if len(l) != 0 {
					l = l[:len(l) - 1]
				} else if len(a) != 0 {
					a = a[:len(a) - 1]
				} else {
					return false
				}
			}
		} else {
			a = append(a, i)
		}
    }

	for len(l) > 0 {
		if len(a) == 0 || a[len(a) - 1] < l[len(l) - 1] {
			return false
		}

		a = a[:len(a) - 1]
		l = l[:len(l) - 1]
	}
	return len(a) % 2 == 0
}
