package web

import (
	"OSRailwayControl/handler"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{} // use default options

type socket struct {
	conn    map[uuid.UUID]*websocket.Conn
	actions map[string]handler.Action
}

func newSocket() handler.Socket {
	s := socket{
		conn:    make(map[uuid.UUID]*websocket.Conn),
		actions: make(map[string]handler.Action),
	}

	return &s
}

func (s *socket) Register(key string, action handler.Action) {
	s.actions[key] = action
}

func (s *socket) receive(sessionId uuid.UUID, _ int, message []byte) error {
	m := handler.Message{}
	err := json.Unmarshal(message, &m)
	if err != nil {
		return err
	}

	action, ok := s.actions[m.Key]
	if !ok {
		return errors.New("could not find an action for the key: " + m.Key)
	}

	return action(sessionId, m.Value)
}

func (s *socket) SendAll(message handler.Message) error {
	for sessionId := range s.conn {
		err := s.Send(sessionId, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *socket) Send(sessionId uuid.UUID, message handler.Message) error {
	conn, ok := s.conn[sessionId]
	if !ok {
		return errors.New("session does not exist: " + sessionId.String())
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, jsonMessage)
}

func (s *socket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer func() {
		_ = c.Close()
	}()

	sessionId := uuid.New()

	closedSignal := make(chan bool)
	c.SetCloseHandler(func(code int, text string) error {
		fmt.Println(code, text)
		closedSignal <- true
		return nil
	})

	s.conn[sessionId] = c

	for {
		select {
		case closed := <-closedSignal:
			if closed {
				break
			}
		default:
		}

		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("receive:", err)
		}
		err = s.receive(sessionId, mt, message)
		if err != nil {
			log.Println("receive:", err)
		}
	}
}

func (w *web) setupSocketListeners() {
	w.socket.Register("getTrains", func(sessionId uuid.UUID, _ string) error {
		trains := w.app.TrainStore.GetTrains()

		jsonTrains, err := json.Marshal(trains)
		if err != nil {
			return err
		}

		return w.socket.Send(sessionId, handler.Message{
			Key:   "getTrains",
			Value: string(jsonTrains),
		})
	})

	w.socket.Register("setSpeed", func(sessionId uuid.UUID, value string) error {
		var message struct {
			Train string `json:"train"`
			Speed int    `json:"speed"`
		}

		err := json.Unmarshal([]byte(value), &message)
		if err != nil {
			return err
		}

		return w.app.Mqtt.Publish("/OSRailway/"+message.Train+"/drive", 0, strconv.Itoa(message.Speed))
	})
}
