package nats

import (
	nats "github.com/nats-io/nats.go"
)

// NewJSONConnection gives you back a json encoded nats connection
func NewJSONConnection() (*nats.EncodedConn, error) {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		return nil, err
	}

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	return c, err
}
