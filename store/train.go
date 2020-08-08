package store

import "sync"

type Train string

type TrainStore struct {
	activTrainsMutex sync.Mutex
	activeTrains     map[string]Train
}

func NewTrainStore() TrainStore {
	return TrainStore{
		activTrainsMutex: sync.Mutex{},
		activeTrains:     make(map[string]Train),
	}
}

func (t *TrainStore) AddTrain(name string) {
	t.activTrainsMutex.Lock()
	t.activeTrains[name] = Train(name)
	t.activTrainsMutex.Unlock()
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
