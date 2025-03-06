package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func cpuIntensiveTaskFix(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1_000_000; i++ { // Reduce iterations
		_ = rand.Intn(100) * rand.Intn(100)
	}
}
func cpuIntensiveTask(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1_000_000_000; i++ {
		_ = rand.Intn(100) * rand.Intn(100)
	}
}

func memoryLeakTask(wg *sync.WaitGroup) {
	defer wg.Done()
	data := make([]int, 0) // Inefficient allocation
	for i := 0; i < 1000000; i++ {
		data = append(data, rand.Intn(1000)) // Growing slice in an unoptimized way
	}
}

func deadlockTaskFix(wg *sync.WaitGroup, m1, m2 *sync.Mutex) {
	defer wg.Done()
	m1.Lock()
	m2.Lock()

	fmt.Println("Locked both mutexes safely")
	m2.Unlock()
	m1.Unlock()
}

func deadlockTask(wg *sync.WaitGroup, m1, m2 *sync.Mutex) {
	defer wg.Done()

	m1.Lock()
	time.Sleep(1 * time.Second)
	m2.Lock() // This may cause a deadlock, Deadlock occurs if another goroutine locks in reverse order

	fmt.Println("Locked both mutexes")
	m2.Unlock()
	m1.Unlock()
}

func main() {
	// Start pprof profiling server
	go func() {
		fmt.Println("Starting pprof server at :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	var wg sync.WaitGroup
	var mu1, mu2 sync.Mutex

	// Spawn CPU-intensive goroutines
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go cpuIntensiveTask(&wg)
	}

	// Spawn memory leak goroutine
	wg.Add(1)
	go memoryLeakTask(&wg)

	// Create two goroutines leading to deadlock
	wg.Add(2)
	go deadlockTask(&wg, &mu1, &mu2)
	go deadlockTask(&wg, &mu2, &mu1)

	wg.Wait()
}
