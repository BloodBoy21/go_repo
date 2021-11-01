package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name  string
	Delay time.Duration
	Num   int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	Quit       chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		WorkerPool: workerPool,
		JobQueue:   make(chan Job),
		Quit:       make(chan bool),
	}
}
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
			fmt.Printf("worker %d: started\n", w.Id)
			fib := Fibonacci(job.Num)
			time.Sleep(job.Delay)
			fmt.Printf("worker %d: completed %d\n", w.Id, fib)
			case <-w.Quit:
				fmt.Printf("worker %d stopped\n",w.Id)
				return
				}
			}	
	}()
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}


func Fibonacci(i int) int {
	if i <= 1 {
		return i
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}

func newDispatcher(jobQueue chan Job,maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: workerPool,
		MaxWorkers: maxWorkers,
		JobQueue:   jobQueue,
	}
}
func(d *Dispatcher)Dispatch(){
	for {
		select {
			case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}
func (d *Dispatcher)Run(){
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request,jobQueue chan Job) {
	if r.Method != "POST"{
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value", http.StatusBadRequest)
		return
	}
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Bad num value", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}
	if delay < 0 || value < 0 {
		http.Error(w, "Negative values not allowed", http.StatusBadRequest)
		return
	}
	job := Job{Name: name, Delay: delay, Num : value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)

}

func main(){
	const (
	maxWorkers = 4
	maxQueueSize = 20
	port = "8081"
	)
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := newDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))

}