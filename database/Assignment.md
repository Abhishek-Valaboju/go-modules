## Assignment

### Implementing Complex Database Operations in Go

#### Objective

This assignment will guide you through the practical implementation of complex database operations in Go, focusing on efficient transaction handling, concurrency management with optimistic locking, and error handling. By the end of this assignment, you’ll be able to execute robust database transactions, manage data consistency using versioning, and efficiently handle database errors and connection pooling in Go.

#### Problem Statement

You are developing a backend service for a financial application that manages user accounts. Each account has a balance and a version field to help manage concurrent access. Your goal is to design database operations that update user balances while maintaining data consistency, handling concurrent access, and managing potential errors from database connections.


#### Instructions

1. Set up your Go project and dependencies:
   - Initialize a new Go module in a project folder.
   - Install the necessary PostgreSQL driver by running `go get github.com/lib/pq`.
   - Create a PostgreSQL database and set up a `users` table with the following schema:
   
   ```go
    CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    balance NUMERIC(10, 2) DEFAULT 0.00,
    version INT DEFAULT 1
    );
    ```

2. Transaction Handling and Optimization (Avoid Long-Running Transactions):
   - Implement a function `UpdateBalance(db *sql.DB, userID int, newBalance float64) error` that updates the user’s balance.
   - Avoid long-running transactions:
     - First, retrieve the user’s current balance outside of the transaction.
     - Then, initiate a transaction, perform only the update within it, and commit the transaction.
     - Ensure that your function minimizes lock contention by keeping the transaction scope small.

3. Concurrency Management with Optimistic Locking:
   - Modify the `UpdateBalance` function to use an optimistic locking mechanism.
   - Add a version check in the update statement. Use the `version` column to ensure that the balance update only occurs if the version hasn’t changed.
   - Increment the `version` field by 1 after each update.
   - If the version does not match, return a custom error (e.g., `ErrVersionConflict`) indicating a version conflict.

4. Error Handling and Logging:

   - Explicitly handle the `sql.ErrNoRows` and `sql.ErrConnDone` errors.
   - If `sql.ErrNoRows` is encountered when retrieving the user’s balance, log the error and return an informative message indicating that the user was not found.

5. Database Connection Pooling:

   - Configure a connection pool with the `sql.DB` instance by:
     - Setting the maximum number of open connections (`SetMaxOpenConns(10)`).
     - Setting the maximum number of idle connections (`SetMaxIdleConns(5)`).
     - Setting the maximum connection lifetime (`SetConnMaxLifetime(time.Hour)`).
   - Test your connection pool configuration by running a simple query that retrieves the total user count and logs the result.

6. Testing and Validation:

   - Write test cases to verify:
     - That the balance update occurs only if the version matches (simulate a version conflict).
     - Handling of `sql.ErrNoRows` when querying a non-existent user.
     - Handling of `sql.ErrConnDone` by closing the database connection manually and retrying the operation.
   - Confirm that transaction times are minimized, versioning works correctly, and errors are handled as expected.

#### Expected Outcome

After completing this assignment:

   - You will have implemented a robust database transaction in Go with optimal transaction handling, optimistic locking, and error management.
   - You will understand how to utilize connection pooling for efficient database operations.
   - You will have built test cases to validate your implementations.
