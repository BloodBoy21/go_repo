package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Println("Calculate Expensive Fibonacci for ", n)
	time.Sleep(time.Second * 5)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	exists := s.InProgress[job]
	if exists {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)
		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Println("Waiting for response from job ", job)
		res := <-response
		fmt.Println("Got response from job ", job, ":", res)
		return
	}
	s.Lock.RUnlock()
	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()
	fmt.Println("Starting job ", job)
	res := ExpensiveFibonacci(job)
	s.Lock.RLock()
	pendingWorkers,exists := s.IsPending[job]
	s.Lock.RUnlock()
	if exists {
		for _, worker := range pendingWorkers {
			worker <- res
		}
		fmt.Println("Result sent -All pending Workers ready job:" , job)

	}
	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([] chan int, 0)
	s.Lock.Unlock()

}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}
func main() {
	service := NewService()
	jobs := []int{1, 2, 3, 4,5,4,8,8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, job := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(job)}
	wg.Wait()
	}
