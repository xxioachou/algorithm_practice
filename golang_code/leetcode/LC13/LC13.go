package main

func romanToInt(s string) int {
	ans := 0
	for i := 0; i < len(s); {
		if s[i] == 'I' {
			if i + 1 < len(s) && s[i + 1] == 'V' {
				ans += 4
				i += 2
				continue
			}
			if i + 1 < len(s) && s[i + 1] == 'X' {
				ans += 9
				i += 2
				continue
			}
			ans += 1
		} else if s[i] == 'V' {
			ans += 5
		} else if s[i] == 'X' {
			if i + 1 < len(s) && s[i + 1] == 'L' {
				ans += 40
				i += 2
				continue
			}
			if i + 1 < len(s) && s[i + 1] == 'C' {
				ans += 90
				i += 2
				continue
			}
			ans += 10
		} else if s[i] == 'L' {
			ans += 50
		} else if s[i] == 'C' {
			if i + 1 < len(s) && s[i + 1] == 'D' {
				ans += 400
				i += 2
				continue
			}
			if i + 1 < len(s) && s[i + 1] == 'M' {
				ans += 900
				i += 2
				continue
			}
			ans += 100
		} else if s[i] == 'D' {
			ans += 500
		} else if s[i] == 'M' {
			ans += 1000
		}
		i ++
	}
	return ans
}