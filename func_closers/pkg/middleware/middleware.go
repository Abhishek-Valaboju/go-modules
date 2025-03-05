package middleware

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"func_closers/pkg/ssh"
	"func_closers/pkg/types"
)

// middleware is the encrypt middleware
type middleware struct {
	publicKey *rsa.PublicKey
}

// NewEncryptMiddleware constructs new encrypt middleware
func NewEncryptMiddleware(key *rsa.PublicKey) *middleware {
	return &middleware{
		publicKey: key,
	}
}

// Encrypt is the middleware function which encrypts the request message
func (m middleware) Encrypt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse client request
		message := types.External{}
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Printf("EncryptRequestMiddleware: Error reading the input message from request: %v", err)

			return
		}

		// Encrypt input message using the public key
		// Note: the handler function is referencing e.publicKey from outside func body. Hence, function closures.
		encryptedMsg, err := ssh.EncryptData([]byte(message.Message), m.publicKey)
		if err != nil {
			log.Printf("EncryptRequestMiddleware: Error encrypting input message using public key: %v", err)

			return
		}

		// Write the encrypted input message back to request for inner handler(s) to use
		requestWriter := bytes.NewBuffer([]byte{})
		if err = json.NewEncoder(requestWriter).Encode(types.Internal{
			Message: encryptedMsg,
		}); err != nil {
			log.Printf("EncryptRequestMiddleware: Error writing to request to forward the encrypting input message: %v", err)

			return
		}
		r.Body = io.NopCloser(requestWriter)

		// Call the inner handler(s)
		// Note: the handler function is referencing next from outside func body. Hence, function closures.
		next.ServeHTTP(w, r)
	})
}
