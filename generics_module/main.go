package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Task[T any] interface {
	Execute() error
	GetPriority() int
}
type TaskQueue[T Task[T]] struct {
	queue []T
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewTaskQueue[T Task[T]]() *TaskQueue[T] {
	q := &TaskQueue[T]{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *TaskQueue[T]) Enqueue(task T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, task)
	// Sort by priority (higher priority first)
	for i := len(q.queue) - 1; i > 0 && q.queue[i].GetPriority() > q.queue[i-1].GetPriority(); i-- {
		q.queue[i], q.queue[i-1] = q.queue[i-1], q.queue[i]
	}
	q.cond.Signal()
}

// Dequeue removes and returns the highest-priority task
func (q *TaskQueue[T]) Dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.queue) == 0 {
		q.cond.Wait()
	}
	task := q.queue[0]
	q.queue = q.queue[1:]
	return task
}

// Worker Pool
type WorkerPool[T Task[T]] struct {
	queue   *TaskQueue[T]
	workers int
	wg      sync.WaitGroup
}

// NewWorkerPool initializes workers
func NewWorkerPool[T Task[T]](queue *TaskQueue[T], workers int) *WorkerPool[T] {
	return &WorkerPool[T]{queue: queue, workers: workers}
}

// Start workers to process tasks concurrently
func (wp *WorkerPool[T]) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for {
				task := wp.queue.Dequeue()
				log.Printf("[Worker %d] Processing task: %+v\n", workerID, task)
				if err := task.Execute(); err != nil {
					log.Printf("[Worker %d] Task failed: %v\n", workerID, err)
				}
			}
		}(i + 1)
	}
}

// Wait for all workers to finish
func (wp *WorkerPool[T]) Wait() {
	wp.wg.Wait()
}

// NotificationTask (Example Task Type)
type NotificationTask struct {
	message  string
	priority int
}

// Implement Task interface
func (n NotificationTask) Execute() error {
	log.Println("Sending Notification:", n.message)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate work
	return nil
}
func (n NotificationTask) GetPriority() int {
	return n.priority
}

// PaymentTask (Example Task Type)
type PaymentTask struct {
	amount   float64
	priority int
}

// Implement Task interface
func (p PaymentTask) Execute() error {
	log.Println("Processing Payment:", p.amount)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate work
	return nil
}
func (p PaymentTask) GetPriority() int {
	return p.priority
}

// ReportTask (Example Task Type)
type ReportTask struct {
	reportType string
	priority   int
}

// Implement Task interface
func (r ReportTask) Execute() error {
	log.Println("Generating Report:", r.reportType)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate work
	return nil
}
func (r ReportTask) GetPriority() int {
	return r.priority
}

// Main function
func main() {
	rand.Seed(time.Now().UnixNano())
	// Create task queue
	queue := NewTaskQueue[Task[any]]()
	// Create and start worker pool
	workerPool := NewWorkerPool(queue, 3)
	go workerPool.Start()
	// Enqueue different tasks
	queue.Enqueue(NotificationTask{"Email to User1", 2})
	queue.Enqueue(PaymentTask{100.50, 5})
	queue.Enqueue(ReportTask{"Sales Report", 1})
	queue.Enqueue(NotificationTask{"SMS to User2", 3})
	queue.Enqueue(PaymentTask{250.00, 4})
	// Wait for tasks to be processed
	time.Sleep(3 * time.Second)
}
