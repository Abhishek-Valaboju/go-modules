package main

import (
	"fmt"
	"runtime"
)

type Data struct {
	Payload [1024]byte // 1 KB data
}

var leakSlice []Data

func memoryLeak() {
	for i := 0; i < 1000000; i++ {
		leakSlice = append(leakSlice, Data{})
		if i%100000 == 0 {
			runtime.GC()
			printMemUsage("Leak Simulation")
		}
	}
}
func printMemUsage(stage string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s HeapAlloc = %v KB, NumGC = %v\n", stage, m.HeapAlloc/1024, m.NumGC)
}
func memoryOptimized() {
	fmt.Println("Memory Optimized")
	leakSlice = nil
	runtime.GC()
	printMemUsage("After optimization")
}

func main() {
	fmt.Println("Starting memory leak simulation...")
	memoryLeak()
	memoryOptimized()
}
