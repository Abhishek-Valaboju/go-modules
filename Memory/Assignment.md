### Problem 1:
### GC Pressure & Memory Leak Detection
**Scenario:**

You are developing a high-performance web server that serves thousands of requests per second. Users report that memory usage keeps increasing over time, even when traffic stays constant. You suspect a memory leak due to objects not being garbage collected.

**Task:**

Write a Go program that simulates a memory leak by continuously allocating memory without proper cleanup.
Use runtime.ReadMemStats to monitor heap allocation and identify the issue.
Modify the program to fix the memory leak and ensure Go's garbage collector can reclaim memory.

**Hints:**

- Use `runtime.MemStats` to track `HeapAlloc`, `HeapInUse`, `NumGC`.
- Identify objects that are escaping to the heap unnecessarily.
- Implement sync.Pool or manual dereferencing to reduce GC pressure.