package handler

import (
	"github.com/google/uuid"
	"net/http"
)

type Handler interface {
	Listen() error
}

type WebHandler interface {
	Handler
	Socket() Socket
}

type MQTTHandler interface {
	Handler
	Publish(topic string, qos byte, payload string) error
}

type Message struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Action func(sessionId uuid.UUID, value string) error

type Socket interface {
	Register(key string, action Action)
	SendAll(message Message) error
	Send(sessionId uuid.UUID, message Message) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
