package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sum(4, 2,3,1,222))
}