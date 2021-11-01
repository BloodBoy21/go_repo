package main

import "fmt"

func main() {
	task := []int{2, 4, 6, 10, 12}
	nWorkers := 4
	jobs := make(chan int, len(task))
	results := make(chan int, len(task))
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}
	for _, j := range task {
		jobs <- j
	}
	close(jobs)
	for i := 0; i < len(task); i++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		fib := Fibonacci(j)
		fmt.Printf("worker %d finished job,result:%d\n", id, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}