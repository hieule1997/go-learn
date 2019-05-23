package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}
func main() {
	fooval := make(chan int)
	go foo(fooval, 5)
	go foo(fooval, 3)
	v1 := <-fooval
	v2 := <-fooval
	fmt.Println(v1, v2)
}
