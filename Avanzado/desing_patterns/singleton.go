package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
	connection string
}
func (Database)CreateSingleConnection() {
	fmt.Println("Creating single connection")
	time.Sleep(time.Second*2)
	fmt.Println("DB connection created")
}
var db *Database
var lock sync.Mutex

func getDatabase() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		db = &Database{
			connection: "connected",
		}
		db.CreateSingleConnection()
	}else{
		fmt.Println("DB connection already exists")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabase()
		}()
	}
	wg.Wait()
}