package mqtt

import (
	"OSRailwayControl/handler"
	"fmt"
)

func (m *mqtt) onNewTrain(message Message) {
	ok := m.app.TrainStore.AddTrain(string(message.Payload()))
	if !ok {
		fmt.Println("could not add train")
		return
	}

	err := m.app.Webserver.Socket().SendAll(handler.Message{
		Key: "addTrain", Value: string(message.Payload()),
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (m *mqtt) onRemoveTrain(message Message) {
	m.app.TrainStore.RemoveTrain(string(message.Payload()))
	err := m.app.Webserver.Socket().SendAll(handler.Message{
		Key: "removeTrain", Value: string(message.Payload()),
	})

	if err != nil {
		fmt.Println(err)
	}
}

func (m *mqtt) registerTopics() error {
	err := m.Register("/OSRailway/register", 0, m.onNewTrain)
	if err != nil {
		return err
	}
	err = m.Register("/OSRailway/remove", 0, m.onRemoveTrain)
	if err != nil {
		return err
	}

	return nil
}
