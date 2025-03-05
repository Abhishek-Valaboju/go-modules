package types

// External is the request type for client request
type External struct {
	Message string `json:"message"`
}

// Internal to be used between middlewares and handlers
type Internal struct {
	Message []byte `json:"message"`
}
