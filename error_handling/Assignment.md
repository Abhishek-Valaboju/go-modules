## Assignment

### Custom Error Handling and Panic Recovery

#### Objective

The goal of this assignment is to help you learn how to:

- Implement custom error types by extending the `error` interface.
- Handle errors using proper error checks.
- Use `panic`, `recover`, and `defer` for managing critical failures and ensuring resource cleanup.

#### Problem Statement

You are tasked with creating a user authentication system for an application. In this system, you need to handle errors for invalid input and manage critical issues such as system failures (simulated using panics).

The system should:

1. Validate user credentials.
2. Return specific errors for invalid inputs (like missing username or password).
3. Ensure that the system handles unexpected failures using `panic` and `recover`.
4. Always clean up resources, like closing database connections (simulated), regardless of whether the function succeeds or fails.

You are not required to build a fully functional authentication system with external APIs. The focus is on error handling.

#### Requirements

1. **Create a custom error** type `ValidationError` that implements the `error` interface. This error will be used for invalid inputs (e.g., missing username or password).
2. **Create a function** `Authenticate(username, password string) error` that:
   - Returns a `ValidationError` if the username or password is empty.
   - Simulates a critical failure by using `panic` if the username is "panicUser".
3. \*\*Use <code>recover</code></strong> to handle panics gracefully and allow the program to continue execution.
4. <strong>Use <code>defer</code></strong> to ensure that cleanup actions (such as closing a database connection, simulated with a simple log message) always occur, even in the event of a panic or error.
