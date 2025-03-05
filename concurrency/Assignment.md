**Assignment: Concurrent Chat Server**

**Objective**

The objective of this assignment is to design and implement a concurrent chat server in Go that allows multiple clients to connect and communicate with each other in real time. This project will challenge your understanding of goroutines, channels, and synchronization techniques, providing a practical application of Go’s concurrency model.

**Problem Statement**

You are tasked with creating a chat server that can handle multiple user connections simultaneously. Each user connects to the server via a TCP socket and can send and receive messages in real time. The server should echo messages sent by any client to all other connected clients, ensuring that all participants in the chat receive updates. Additionally, the server must handle client disconnections gracefully and ensure that resources are freed correctly.

**Requirements**

1. **Server Implementation:**
* Create a TCP server that listens on a specified port.
* Use goroutines to handle each client connection concurrently.
* Maintain a list of connected clients and manage their connections.
2. **Client Connections:**
* Each client should be able to send messages to the server.
* The server must broadcast incoming messages to all connected clients.
* Implement functionality for clients to disconnect from the server, which should also notify other clients.
3. **Message Handling:**
* Use channels to manage message passing between the server and connected clients.
* Implement a mechanism for clients to receive messages in real time, utilizing select statements to handle multiple communication channels.
4. **Graceful Shutdown:**
* Ensure that when a client disconnects, their resources are cleaned up, and other clients are informed of the disconnection.
* Implement a way for the server to shut down gracefully, allowing all clients to receive a final message before closing.
5. **Concurrency and Synchronization:**
* Utilize sync.WaitGroup to manage the completion of goroutines handling client connections.
* Consider using a mutex (from the sync package) to safely manage shared data structures (like the list of connected clients) to prevent race conditions.
6. **Testing:**
* Write test cases to verify the functionality of the server. Consider edge cases like rapid connect/disconnect and message overflow.
* Use Go’s built-in testing framework to create unit tests for your code.
7. **Documentation:**
* Provide clear documentation for your code, including comments and a README file explaining how to run the server and connect clients.
* Include instructions for testing the server.
8. **Bonus Features (Optional):**
* Implement user authentication, allowing users to log in with a username.
* Add functionality for private messaging between clients.
* Create a simple web interface where users can connect to the chat server via a web browser using WebSockets.

**Deliverables**



* Complete source code for the chat server and any client implementations.
* A README file with instructions on how to set up, run, and test the application.
* Documentation of any additional features implemented, along with a brief explanation of the design choices made.

**Evaluation Criteria**



* Correctness of the implementation: Does the server function as specified? Are messages correctly broadcasted?
* Code quality: Is the code well-organized, readable, and maintainable? Are appropriate Go idioms used?
* Concurrency management: Are goroutines and channels used effectively? Are race conditions handled properly?
* Documentation: Is the code adequately documented? Are the instructions clear and easy to follow?
* Bonus features: If applicable, are the bonus features implemented effectively and seamlessly integrated into the main application?