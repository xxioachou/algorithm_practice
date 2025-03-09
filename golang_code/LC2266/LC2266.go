package main

func countTexts(pressedKeys string) int {
	const mod = 1e9 + 7
    n := len(pressedKeys)
	f := make([]int, n + 1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = f[i - 1]
		if i - 1 > 0 && pressedKeys[i - 1 - 1] == pressedKeys[i - 1] {
			f[i] += f[i - 2]
			f[i] %= mod
			if i - 2 > 0 && pressedKeys[i - 2 - 1] == pressedKeys[i - 1 - 1] {
				f[i] += f[i - 3]
				f[i] %= mod
				if (pressedKeys[i - 1] == '7' || pressedKeys[i - 1] == '9') && 
					i - 3 > 0 && pressedKeys[i - 3 - 1] == pressedKeys[i - 2 - 1] {
					f[i] += f[i - 4];
					f[i] %= mod
				}
			}
		}
	}
	return f[n]
}

func main() {

}
