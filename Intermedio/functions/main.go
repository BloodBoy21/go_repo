package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	getDouble := func(z int) int {
		return z * 2
	}
	y:= func()int{
		return x*10
	}()
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(getDouble(10))
	c:=make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(time.Second*5)
		fmt.Println("End")
		c<-1
	}()
	fmt.Println(<-c)
}