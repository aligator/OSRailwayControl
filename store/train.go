package store

import (
	"regexp"
	"sync"
)

type Train string

type TrainStore struct {
	activTrainsMutex sync.Mutex
	activeTrains     map[string]Train

	trainNameRegexp *regexp.Regexp
}

func NewTrainStore() TrainStore {
	r, err := regexp.Compile("^[a-zA-Z0-9]+$")
	if err != nil {
		panic(err)
	}
	return TrainStore{
		activTrainsMutex: sync.Mutex{},
		activeTrains:     make(map[string]Train),
		trainNameRegexp:  r,
	}
}

func (t *TrainStore) AddTrain(name string) bool {
	t.activTrainsMutex.Lock()
	ok := t.trainNameRegexp.MatchString(name)
	if ok {
		t.activeTrains[name] = Train(name)
	}
	t.activTrainsMutex.Unlock()
	return ok
}

func (t *TrainStore) RemoveTrain(name string) {
	t.activTrainsMutex.Lock()
	delete(t.activeTrains, name)
	t.activTrainsMutex.Unlock()
}

func (t *TrainStore) GetTrains() []Train {
	t.activTrainsMutex.Lock()
	trains := make([]Train, 0)
	for _, train := range t.activeTrains {
		trains = append(trains, train)
	}
	t.activTrainsMutex.Unlock()

	return trains
}
