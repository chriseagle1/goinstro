package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go printHelloWorld(i, ch)
	}
	for {
		msg := <- ch
		fmt.Println(msg)
	}
}

func printHelloWorld(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("hello world, Gorouting num %d!\n", i)
	}
}
