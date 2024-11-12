package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Lesson is goroutines
// var m = sync.Mutex{}
// var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

//var results = []string{}

func main() {
	//waiting
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	//fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	// Simulate Database Call Delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Println("The result from the database is:", dbData[i])
	//save(dbData[i])
	//results = append(results, dbData[i])
	//log()
	wg.Done()
}

/*func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
*/
