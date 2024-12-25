package main

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	// 车后同行且象不在他们之间
	if a == e && (c != a || d < min(b, f) || d > max(b, f)) {
		return 1
	}
	// 车后同列且象不在他们之间
	if b == f && (d != b || c < min(a, e) || c > max(a, e)) {
		return 1
	}
	// 象后同一对角线且车不在他们之间 
	if c + d == e + f && (a + b != c + d || a < min(c, e) || a > max(c, e)) {
		return 1
	}
	if c - d == e - f && (a - b != c - d || a < min(c, e) || a > max(c, e)) {
		return 1
	}
	return 2
} 

func main() {

}