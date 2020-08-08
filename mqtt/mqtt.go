package mqtt

import (
	"OSRailwayControl/app"
	"OSRailwayControl/handler"
	"errors"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strconv"
)

const notConnectedError = "not connected yet"

type Message interface {
	Duplicate() bool
	Qos() byte
	Retained() bool
	Topic() string
	MessageID() uint16
	Payload() []byte
	Ack()
}

type mqtt struct {
	app      *app.App
	port     int
	host     string
	user     string
	password string

	client MQTT.Client
}

func NewMQTT(app *app.App) handler.MQTTHandler {
	w := mqtt{
		app:      app,
		port:     app.Config.MQTT.Port,
		host:     app.Config.MQTT.Host,
		user:     app.Config.MQTT.User,
		password: app.Config.MQTT.Password,
	}
	return &w
}

func (m *mqtt) Listen() error {
	opts := MQTT.NewClientOptions()
	broker := m.host + ":" + strconv.Itoa(m.port)
	opts.AddBroker(broker)
	if m.user != "" {
		opts.SetUsername(m.user)
	}

	if m.password != "" {
		opts.SetPassword(m.password)
	}

	opts.SetCleanSession(true)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	m.client = client

	err := m.registerTopics()
	if err != nil {
		return err
	}

	for m.client.IsConnected() {
		// do nothing
	}

	return nil
}

func (m *mqtt) Register(topic string, qos byte, callback func(message Message)) error {
	if !m.isConnected() {
		return errors.New(notConnectedError)
	}

	m.client.Subscribe(topic, qos, func(_ MQTT.Client, message MQTT.Message) {
		callback(message)
	})

	return nil
}

func (m *mqtt) Publish(topic string, qos byte, payload string) error {
	if !m.isConnected() {
		return errors.New(notConnectedError)
	}

	m.client.Publish(topic, qos, false, payload)
	return nil
}

func (m *mqtt) isConnected() bool {
	return m.client != nil && m.client.IsConnected()
}
