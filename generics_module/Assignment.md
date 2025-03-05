# Assignment: Building a Distributed Task Queue with Go Generics

## **Objective**

The objective of this assignment is to build a **distributed task queue** using Go generics that can handle various types of tasks in a scalable manner. The system should efficiently distribute tasks across multiple workers, allow for task prioritization, and be capable of processing multiple task types concurrently. Using Go generics, the system should be flexible and easy to extend for new task types while maintaining high performance.

---

## **Problem Statement**

### Scenario:
Your company is developing a **cloud-based distributed system** that requires processing a wide range of tasks, such as sending notifications, processing payments, generating reports, and executing machine learning models. Each task type has different execution requirements, and the system needs to handle thousands of tasks concurrently. The company needs a **scalable task queue** that can handle:
- Multiple types of tasks.
- Task prioritization (high-priority tasks should be processed first).
- Concurrent processing with multiple workers.
- Support for dynamically adding new task types without rewriting the core system.

You need to design and implement a **task queue system** that distributes tasks across workers efficiently and ensures that the system can scale based on workload demands.

---

## **Requirements**

### 1. **Generic Task Interface**:
- Define a generic `Task[T any]` interface that represents the task to be executed.
  - Each task should have an `Execute()` method to perform its specific operation.
  - Example tasks:
    - **NotificationTask**: Sends email or SMS notifications.
    - **PaymentTask**: Processes a payment transaction.
    - **ReportTask**: Generates and sends analytical reports.

### 2. **Task Queue**:
- Implement a generic task queue `TaskQueue[T any]` that holds tasks of any type.
  - The queue should support task prioritization (e.g., high-priority tasks should be processed first).
  - Implement functions to:
    - Add tasks to the queue (`Enqueue()`).
    - Remove and process tasks (`Dequeue()`).
    - Peek at the next task (`Peek()`).

### 3. **Worker Pool**:
- Implement a **worker pool** where workers pick up tasks from the queue and process them concurrently.
  - Each worker should run in its own goroutine.
  - Limit the number of concurrent workers to avoid resource exhaustion.
  - Allow dynamic scaling of workers based on the number of pending tasks in the queue.

### 4. **Task Scheduling and Retry Mechanism**:
- Implement task scheduling where tasks can be scheduled to run at specific times or intervals.
- Add support for retrying failed tasks. If a task fails, it should be retried a limited number of times with a delay between attempts.

### 5. **Task Prioritization**:
- Implement a priority system for tasks where high-priority tasks are processed before low-priority tasks.
  - For example, critical notification tasks should be processed before report generation tasks.

### 6. **Dynamic Task Type Support**:
- Ensure that new task types (e.g., data processing, machine learning tasks) can be easily added without modifying the core system.
  - Using Go generics, the task queue should be flexible enough to handle any new task type.

### 7. **Logging and Monitoring**:
- Implement logging for each task's lifecycle (task added, task started, task completed, task failed).
- Implement a monitoring system to track the number of tasks in the queue, number of workers, and task success/failure rates.

---

## **Deliverables**

1. **Source Code**:
   - Implement the task queue, worker pool, and task scheduling system using Go generics.
   - Ensure the system can handle multiple task types and support concurrency.

2. **Example Use Case**:
   - Demonstrate the system with three task types:
     1. **NotificationTask** for sending email/SMS.
     2. **PaymentTask** for processing payments.
     3. **ReportTask** for generating reports.
   - Add tasks of each type to the queue, process them using the worker pool, and demonstrate prioritization.

3. **Unit Tests**:
   - Write unit tests to validate the task queue, worker pool, task prioritization, and retry mechanism.

4. **Documentation**:
   - Write a README explaining how the task queue works, how to add new task types, and how to run the example use case.
   - Include instructions on configuring the worker pool and task scheduling.

---

## **Evaluation Criteria**

- **Correctness**:
  - Does the implementation correctly enqueue, prioritize, and process tasks?
  - Does the worker pool handle concurrency and dynamic scaling efficiently?

- **Use of Go Generics**:
  - Are generics used appropriately to support different task types? Does the system allow easy extension for new task types?

- **Concurrency Handling**:
  - Does the system handle multiple tasks concurrently without resource contention? Is the worker pool implemented efficiently?

- **Task Prioritization and Scheduling**:
  - Is task prioritization implemented correctly? Are high-priority tasks processed before lower-priority tasks?
  - Is task scheduling and retry functionality implemented correctly?

- **Code Structure and Documentation**:
  - Is the code modular, easy to follow, and well-documented? Is the process for adding new task types clear?

---

### **Bonus**:
- Implement **distributed queue** support where tasks are distributed across multiple servers.
- Add **persistence** for the task queue so that tasks are not lost if the system crashes.
- Implement a **dashboard** to visualize the queue, worker status, and task metrics in real time.

---

## **Example Use Case**

A distributed e-commerce system requires processing multiple tasks concurrently:
1. **NotificationTask**: Sends transactional emails or SMS to customers.
2. **PaymentTask**: Processes credit card payments.
3. **ReportTask**: Generates weekly sales reports.

The task queue should handle thousands of tasks per second, prioritize payment processing over report generation, and ensure that failed tasks (e.g., payment failures) are retried.

---

This assignment reflects a **real-world distributed system** challenge, where task queues are crucial for handling different workloads efficiently. Go generics allow flexibility, scalability, and extensibility for handling various task types with minimal code duplication.
