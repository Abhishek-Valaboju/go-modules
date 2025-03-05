## Assignment

### Task Management REST API

#### Overview

You are tasked with building a REST API for a task management system. Users will be able to:

1. **Create** new tasks
2. **View** all tasks or a specific task
3. **Update** a task (mark it complete or change its details)
4. **Delete** a task

Each task will have the following properties:

- ID (int): A unique identifier for the task.
- Title (string): A short description of the task.
- Description (string): A detailed description of the task.
- Completed (bool): Whether the task is completed or not.
- Due Date (string): The date by which the task should be completed.

#### Requirements

- Use Go’s standard `net/http` package to build the API.
- Store tasks in memory using a slice of structs (in-memory persistence).
- Proper error handling and status codes should be implemented.
- Use middleware to log requests.
- Implement validation for task data (title must not be empty, etc.).
- Test the API using tools like curl or Postman.

#### Assignment Instructions

##### 1. Project Structure

Organize the project files to make it easy to maintain and extend. The structure should look like following.

```
task-manager/
│
├── cmd
│   └── task-manager
│       └── main.go # Main entry point for the application
|
|── pkg
|   └── data-store
|         └── tasks.go
├── controller
│   ├── handler
│   │   └── handler.go # HTTP handler functions for each endpoint
│   ├── middleware
│   │   └── middleware.go # Middleware (e.g., logging)
└── go.mod # Go module file
```

##### 2. API Endpoints

<table>
  <tr>
   <td><strong>HTTP Method</strong>
   </td>
   <td><strong>Endpoint</strong>
   </td>
   <td><strong>Description</strong>
   </td>
  </tr>
  <tr>
   <td><code>GET</code>
   </td>
   <td><code>/tasks</code>
   </td>
   <td>Get a list of all tasks
   </td>
  </tr>
  <tr>
   <td><code>GET</code>
   </td>
   <td><code>/tasks/{id}</code>
   </td>
   <td>Get details of a specific task
   </td>
  </tr>
  <tr>
   <td><code>POST</code>
   </td>
   <td><code>/tasks</code>
   </td>
   <td>Create a new task
   </td>
  </tr>
  <tr>
   <td><code>PUT</code>
   </td>
   <td><code>/tasks/{id}</code>
   </td>
   <td>Update a task (title, description, or mark completed)
   </td>
  </tr>
  <tr>
   <td><code>DELETE</code>
   </td>
   <td><code>/tasks/{id}</code>
   </td>
   <td>Delete a task
   </td>
  </tr>
</table>

##### 3. Data Model

Each task will be represented by a Go struct:

```go
type Task struct {
   ID          int        `json:"id"`
   Title       string     `json:"title"`
   Description string     `json:"description"`
   Completed   bool       `json:"completed"`
   DueDate     time.Time  `json:"due_date"` // Format: YYYY-MM-DD
}
```
