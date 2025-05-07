package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadArray() []int {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.TrimRight(str, "\r\n")
	parts := strings.Split(str, ",")
	a := make([]int, len(parts))
	for i := range parts {
		a[i], _ = strconv.Atoi(parts[i])
	}
	return a
}

func main() {
	A := ReadArray()
	B := ReadArray()
	mp := make(map[int]struct{})
	ans := make(map[int]struct{})
	for _, x := range A {
		mp[x] = struct{}{}
	}
	for _, x := range B {
		if _, ok := mp[x]; ok {
			ans[x] = struct{}{}
		}
	}
	ids := make([]string, 0)
	for k := range ans {
		ids = append(ids, strconv.Itoa(k))
	}
	slices.Sort(ids)
	fmt.Println(strings.Join(ids, ","))
}
