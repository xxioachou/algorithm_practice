package main

func isSubstringPresent(s string) bool {
    mp := make([]int, 26 * 26)
	for i := 0; i + 1 < len(s); i++ {
		a := int(s[i] - 'a')
		b := int(s[i + 1] - 'a')
		mp[a * 26 + b] ++
	}
	for i := 0; i + 1 < len(s); i++ {
		a := int(s[i] - 'a')
		b := int(s[i + 1] - 'a')
		if mp[b * 26 + a] > 0 {
			return true
		}
	}
	return false
} 

func main() {

}