package main

import "fmt"

func main() {
	c := make(chan int, 2)//channel con buffer
	c2 := make(chan int)//sin buffer
	c <- 1
	c <- 199
	c2<-200
	fmt.Println(<-c)
	fmt.Println(<-c)

}