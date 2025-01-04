package main

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	strs := strings.Split(date, "-")
	a := make([]string, 0)
	for _, str := range strs {
		x, _ := strconv.Atoi(str)
		a = append(a, strconv.FormatInt(int64(x), 2))
	}
	return strings.Join(a, "-")
}

func main() {

}