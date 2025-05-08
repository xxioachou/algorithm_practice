package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func trap(height []int) int {

}

func main() {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	parts := strings.Split(str, " ")
	h := make([]int, len(parts))
	for i := range h {
		h[i], _ = strconv.Atoi(parts[i])
	}
	fmt.Println("h:", h)
	fmt.Println(trap(h))
}