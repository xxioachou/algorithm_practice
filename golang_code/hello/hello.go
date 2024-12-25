package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// for {
	// 	var a, b int
	// 	_, err := fmt.Scan(&a, &b)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(a + b)
	// }
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		nums := strings.Split(scan.Text(), " ")
		for i := 0; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			fmt.Printf("%d ", num)
		}
		fmt.Println("")
		fmt.Println(strings.Join(nums, " "))
	}
}
