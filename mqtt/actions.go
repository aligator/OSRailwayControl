package mqtt

import (
	"OSRailwayControl/handler"
	"OSRailwayControl/store"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (m *mqtt) updateTrain(train store.Train) error {
	trainJson, err := json.Marshal(train)
	if err != nil {
		return err
	}

	err = m.app.TrainStore.SetTrain(train)
	if err != nil {
		return err
	}

	err = m.app.Webserver.Socket().SendAll(handler.Message{
		Key: "updateTrain", Value: string(trainJson),
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *mqtt) extractTrainName(topic string) (string, error) {
	if !strings.HasPrefix(topic, m.topicPrefix) {
		return "", errors.New("the topic has the wrong prefix")
	}

	topic = strings.TrimPrefix(topic, m.topicPrefix+"/")

	split := strings.SplitN(topic, "/", 3)
	if len(split) < 2 {
		return "", errors.New("cannot split the topic to get the train name")
	}

	return split[0], nil
}

func (m *mqtt) onTrainStatus(message Message) {
	trainName, err := m.extractTrainName(message.Topic())
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(message.Payload()) == 0 || string(message.Payload()) == "" {
		// completely remove train
		m.app.TrainStore.RemoveTrain(trainName)
		err := m.app.Webserver.Socket().SendAll(handler.Message{
			Key: "removeTrain", Value: string(trainName),
		})

		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	status := string(message.Payload()) == "1"

	train, ok := m.app.TrainStore.GetTrain(trainName)
	if !ok {
		train = store.Train{
			Name:      trainName,
			Status:    status,
			Speed:     0,
			Direction: 1,
		}
	} else {
		train.Status = status
	}

	err = m.updateTrain(train)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (m *mqtt) onSpeedChange(message Message) {
	trainName, err := m.extractTrainName(message.Topic())
	if err != nil {
		fmt.Println(err)
		return
	}

	speed, err := strconv.Atoi(string(message.Payload()))
	if err != nil {
		fmt.Println(err)
		return
	}

	train, ok := m.app.TrainStore.GetTrain(trainName)
	if !ok {
		fmt.Println("train not found", trainName)
	}

	train.Speed = int(math.Round(math.Abs(float64(speed))))
	if speed < 0 {
		train.Direction = -1
	} else if speed > 0 {
		train.Direction = 1
	} // else leave as it is

	err = m.updateTrain(train)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (m *mqtt) registerTopics() error {
	err := m.Register("/OSRailway/+/status", 0, m.onTrainStatus)
	if err != nil {
		return err
	}
	err = m.Register("/OSRailway/+/drive", 0, m.onSpeedChange)
	if err != nil {
		return err
	}
	err = m.Register("/OSRailway/+/drive/force", 0, m.onSpeedChange)
	if err != nil {
		return err
	}

	return nil
}
