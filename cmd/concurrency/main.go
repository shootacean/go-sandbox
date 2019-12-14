package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go count(10, c)
	for n := range c {
		fmt.Printf("received: %d\n", n)
	}
}

func count(n int, c chan int) {
	for i := 0; i < n; i++ {
		fmt.Printf("send: %d\n", i)
		c <- i
	}
	close(c)
}
