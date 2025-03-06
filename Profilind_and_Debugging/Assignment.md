## Problem: 
Performance Bottleneck Analysis and Debugging in a Concurrent Go Program

#### Scenario
You are building a high-performance concurrent Go application that processes millions of data points in real-time. However, your program exhibits high CPU usage, excessive memory allocation, and deadlocks under load.
Your task is to:
- Profile the CPU and memory usage using pprof
- Identify and fix goroutine blocking using go tool trace
- Debug deadlocks and unexpected crashes using Delve (dlv)

#### Task 1: Implement a Sample Go Program with Performance Issues
Write a concurrent Go program that:
Spawns multiple goroutines to process a large dataset.
Uses an inefficient memory allocation approach.
Contains an intentional deadlock.
#### Task 2: CPU & Memory Profiling using pprof
Start the pprof server and analyze performance bottlenecks.
#### Task 3: Tracing Goroutines using go tool trace
#### Task 4: Debugging Deadlocks & Crashes using Delve
#### Task 5: Fixing Performance Issues












