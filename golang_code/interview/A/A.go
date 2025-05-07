package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	for i := 1; i <= 100; i++ {
		l.PushBack(i)
	}
	ptr := l.Front()
	i := 0
	for ptr != nil {
		next := ptr.Next()
		if next == nil {
			next = l.Front()
		}
		if i%9 == 8 {
			fmt.Print(ptr.Value.(int), " ")
			l.Remove(ptr)
		}
		ptr, i = next, i+1
	}
	fmt.Println("")
}
