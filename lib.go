package lib

import "time"

type Client struct {
	// The amount of time a request can take before it will be
	// terminated. Millisecond precision is supported.
	Timeout time.Duration
}

var DefaultClient = Client{}
