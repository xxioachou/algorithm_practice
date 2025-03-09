package main

import (
	"log"
	"time"
)


func solve() {

}

func main() { 
	ch := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- true
	}()
	select {
	case ok := <- ch:
		log.Printf("%+v", ok)
	case <- time.After(3 * time.Second):
		log.Printf("2")
	}
}