//Memorize pattern
package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{f: f, cache: make(map[int]FunctionResult)}

}
func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	if result, ok := m.cache[key]; ok {
		m.lock.Unlock()
		return result.value, result.err
	}
	value, err := m.f(key)
	m.cache[key] = FunctionResult{value, err}
	m.lock.Unlock()
	return value, err
}

func GetFibonnaci(n int) (interface{}, error) {
	return Fibonnaci(n), nil
}
func main() {
	cache := NewCache(GetFibonnaci)
	fibo := []int{42,40,41,42,38}
	var wg sync.WaitGroup
	for _, v := range fibo {
		wg.Add(1)
		go func (index int)  {
		defer	wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%d: %d,T:%v\n", index, value, time.Since(start))
		}(v)
	}	
	wg.Wait()
}