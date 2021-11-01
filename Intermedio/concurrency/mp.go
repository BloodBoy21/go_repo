package main

import (
	"fmt"
	"time"
)


func main() {
	c := make(chan int)
	c2 := make(chan int)
	d1 := 4*time.Second
	d2 := 2*time.Second
	go DoSomething(d1,c,1)
	go DoSomething(d2,c2,2)
	for i:=0;i<2;i++{
		select {
		case v:= <-c:
			fmt.Println("received",v)
		case v:= <-c2:
			fmt.Println("received",v)
		}
	}
}

func DoSomething(i time.Duration,c chan<-int,param int) {
	time.Sleep(i)
	c <- param


	
}