package store

import (
	"errors"
	"regexp"
	"sync"
)

type Train struct {
	Name      string `json:"name"`
	Status    bool   `json:"status"`
	Speed     int    `json:"speed"`
	Direction int    `json:"direction"`
}

type TrainStore struct {
	trainsMutex sync.Mutex
	trains      map[string]Train

	trainNameRegexp *regexp.Regexp
}

func NewTrainStore() TrainStore {
	r, err := regexp.Compile("^[a-zA-Z0-9]+$")
	if err != nil {
		panic(err)
	}
	return TrainStore{
		trainsMutex:     sync.Mutex{},
		trains:          make(map[string]Train),
		trainNameRegexp: r,
	}
}

func (t *TrainStore) checkTrain(train Train) error {
	ok := t.trainNameRegexp.MatchString(train.Name)
	if !ok {
		return errors.New("invalid train name - should only contain a-z, A-Z, 0-9")
	}

	if train.Speed < 0 || train.Speed > 1023 {
		return errors.New("the train speed has to be from 0 to 1023")
	}

	if train.Direction != -1 && train.Direction != 1 {
		return errors.New("the train diraction can either be '-1' or '1'")
	}

	return nil
}

func (t *TrainStore) SetTrain(train Train) error {
	t.trainsMutex.Lock()
	defer t.trainsMutex.Unlock()

	if err := t.checkTrain(train); err != nil {
		return err
	}

	t.trains[train.Name] = train
	return nil
}

func (t *TrainStore) GetTrain(name string) (Train, bool) {
	train, ok := t.trains[name]
	return train, ok
}

func (t *TrainStore) RemoveTrain(name string) {
	t.trainsMutex.Lock()
	delete(t.trains, name)
	t.trainsMutex.Unlock()
}

func (t *TrainStore) GetTrains() []Train {
	t.trainsMutex.Lock()
	trains := make([]Train, 0)
	for _, train := range t.trains {
		trains = append(trains, train)
	}
	t.trainsMutex.Unlock()

	return trains
}
