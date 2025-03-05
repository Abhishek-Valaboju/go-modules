package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"func_closers/pkg/middleware"
	"func_closers/pkg/ssh"
	"func_closers/pkg/types"
)

func main() {
	// Generate SSH RSA key pair
	privateKey, publicKey, err := ssh.GenerateRSAKeyPair()
	if err != nil {
		log.Fatal(err)
	}

	var port = "localhost:8080"
	var mux = http.NewServeMux()
	var encryptMiddleware = middleware.NewEncryptMiddleware(publicKey)

	// Add home page handler to serve mux
	mux.HandleFunc("/home",
		encryptMiddleware.Encrypt(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// Parse request
					resp := types.Internal{}
					err := json.NewDecoder(r.Body).Decode(&resp)
					if err != nil {
						log.Printf("AccountHandler: Error decoding the input message from request: %v", err)

						return
					}

					// Decrypt (the encrypted) input message using the private key
					// Note: the handler function is referencing privateKey from outside func body. Hence, function closures.
					message, err := ssh.DecryptData(resp.Message, privateKey)
					if err != nil {
						log.Printf("AccountHandler: Error decrypting the input message from request: %v", err)

						return
					}

					// Write the message to response writer
					fmt.Fprintf(w, "Welcome to Home Page. Got the input message: %q\n", message)
				},
			),
		).ServeHTTP,
	)

	// Run the HTTP server
	log.Printf("HTTP serving on address: %q", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
