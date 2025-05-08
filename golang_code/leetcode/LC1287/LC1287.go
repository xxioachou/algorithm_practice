package main

func findSpecialInteger(arr []int) int {
	n := len(arr)
   	for i := 0; i < n; {
		j := i + 1
		for j < n && arr[j] == arr[i] {
			j ++
		}
		cnt := j - i
		if cnt * 4 > n {
			return arr[i]
		}
		i = j
   	} 
	return -1
}

func main() {

}