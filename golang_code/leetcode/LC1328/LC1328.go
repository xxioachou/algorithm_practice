package main

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	b := []byte(palindrome)
	ans := ""

	for i := 0; i < n; i ++ {
		if n % 2 == 1 && i == n / 2 {
			continue
		}

		t := b[i]
		b[i] = 'a'
		if b[i] == b[n - 1 - i] {
			b[i] = 'b'
		}
		if len(ans) != 0 {
			ans = min(ans, string(b))
		} else {
			ans = string(b)
		}
		b[i] = t
	}
	return ans
}
