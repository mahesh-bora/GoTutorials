package main

import (
	"fmt"
	"sync"
	"time"
)


// Go Routines are basically used for parallel execution, also parallel execution doesn't necessaarily mean concurrency
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
var mu sync.RWMutex // Added mutex to safely access the results slice
// var mu sync.RWMutex 

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal Execution Time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v\n", results) // Corrected format specifier
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	// mu.Lock() // Lock before modifying results slice
	// results = append(results, dbData[i])
	// mu.Unlock() // Unlock after modification
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	mu.Lock()
	results = append(results, result)
	mu.Unlock()
}

func log(){
	mu.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	mu.RUnlock()
}
