package main

func replaceElements(arr []int) []int {
    maxv := -1
	tmaxv := -1
	for i := len(arr) - 1; i >= 0; i -- {
		tmaxv = max(maxv, arr[i])
		arr[i] = maxv
		maxv = tmaxv
	}
	return arr
}

func main() {

}