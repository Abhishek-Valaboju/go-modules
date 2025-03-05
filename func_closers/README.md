# Function Closures
A closure is an anonymous function that references variables from outside its body.

## Prerequisites
- Golang basics (syntax, data types, etc.)
- Functions in Golang.
- Anonymous functions.
- _(Optional) Standard libraries and GoDoc._
- _(Optional) HTTP server, handler, and middleware._

## Example
This is a simple example to demonstrate how variables from outside function closure are references, and can also be updated as well.
```go
package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		// Referencing the variable from outside function body, and updates it. i.e., function closures.
		counter += 1
	}

	// Prints 0
	fmt.Println(counter)

	increment()
	// Prints 1
	fmt.Println(counter)

	counter += 10
	increment()
	// Prints 12
	fmt.Println(counter)
}

```

References:
- https://go.dev/tour/moretypes/25
- https://www.geeksforgeeks.org/closures-in-golang/
- https://gobyexample.com/closures
- _(Optional) https://code101.medium.com/understanding-closures-in-go-encapsulating-state-and-behaviour-558ac3617671_

## Some Practical Uses of Closures and Popular Examples
1. Isolating data
    - In a sense where you want some data to persist even after the anonymous function exits.
    <br>Eg. a counter variable. 
2. Middleware functions wrapping handler functions in HTTP server
    - The middleware functions are wrapped around the actual handler functions. The middleware functions call the actual handler function from outside its bounds.
    <br>Check this snippet: https://go.dev/play/p/dG1ciJCB3at
3. Some functions from the standard library that are required to use function closures:
    1. https://pkg.go.dev/sort#Search
        - The function in the second argument `f func(int) bool` makes use of a string or slice from outside, that needs to be searched. Check examples for the Search func in the godoc.
    2. https://pkg.go.dev/sort#Slice
        - The function in the second argument `less func(i, j int) bool` makes use of a string or slice from outside, that needs to be sorted. Check examples for the Search func in the godoc.
    3. https://pkg.go.dev/os#Expand
        - Same as above.
    4. https://pkg.go.dev/strings#TrimFunc
        - Same as above.

References:
- https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/

## Assignment
### Assignment-1: Encryption Middleware
#### Prerequisites
1. Basics of Go.
1. Running HTTP services in Go. Middlewares.
1. Basic understanding of how SSH keys are used for encryption-decryption.

#### Task
Build a Go application that runs an HTTP server with the following:
1. Generate SSH key pair in main func.
2. Encrypt-middleware:
    1. Create a EncryptMiddleware type that holds the public key.
    2. Define a method on EncryptMiddleware which acts as the middleware, i.e., take http.Handler and return http.Handler.
    3. The method should read the request from client, encrypt it using the public key, and write back the encrypted message to request body again so the home handler can read it.
3. Home-handler:
    1. Define and add an in-place /home HandlerFunc that
        1. Reads the encrypted input message from client (through EncryptMiddleware).
        1. Decrypt the message using the private key.
        1. Print a welcome message along with the decrypt input message to response. Eg., `Welcome to Home Page. Got the input message: "abc"`
            > Note: The input message printed in the response should be same as the input from client (user).

##### Things to note:
1. The home handleFunc is expected to use the private key from outside its function body (i.e., from main function) demonstrating function closures.
    ```go
    privateKey, publicKey, err := ssh.GenerateRSAKeyPair()
    ...
    mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        ...
            // Inside the handlerFunc body, referencing privateKey from outside
            message, err := ssh.DecryptData(message, privateKey)
        ...
    }
    ```

2. The handler function that is returned from encrypt middleware method is expected to use the public key (from the receiver of method) and home handler (passed as argument to method) from the method, demonstrating the function closures.
    ```go
    func (m middleware) Encrypt(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ...
            encryptedMsg, err := EncryptData([]byte(message.Message), e.publicKey)
            ...
            next.ServeHTTP(w, r)
        })
    )
    ```

#### Solution
Please use it only if you are unable to solve the problem.

https://go.dev/play/p/1aiWWsIWNP0

> Note: The program will not run in Go playground as it needs a port to run the HTTP server. Copy-paste the code to local, into multiple files and run it. The file structure is described using the tags `-- <filepath> --` in the same.
