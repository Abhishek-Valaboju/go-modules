## Assignment

### Implementing Synchronization with Mutex in Go (Bank Account Simulation)

#### Objective

The objective of this assignment is to demonstrate your understanding of mutex-based synchronization in Golang by solving various concurrency problems. You will use mutexes to ensure data integrity, avoid race conditions, and efficiently manage shared resources between multiple goroutines.

#### Problem Statement

You are tasked with designing and implementing solutions to several advanced-level concurrency problems in Go. Each problem requires careful use of mutexes (`sync.Mutex` and `sync.RWMutex`) to manage synchronization between goroutines and ensure safe access to shared resources.


#### Requirements for Bank Account Simulation

1. Implement a <code>program</code> simulating a <code>shared bank account</code> where multiple users <code>(goroutines)</code> can <code>deposit and withdraw</code> money.
2. Use a <code>mutex</code> to protect the balance and prevent <code>concurrent</code> modifications.
3. Print the <code>final balance</code> after all operations are completed.
4. <strong>Challenge</strong>:
    - Use sync.RWMutex to allow users to check the balance without locking the account for modifications. Prime Number Calculation with Concurrent Workers