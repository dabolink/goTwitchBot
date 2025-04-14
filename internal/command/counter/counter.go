package counter

import (
	"errors"
)

var (
	ErrInvalidNumber  = errors.New("number not next in sequence")
	ErrKeyAlreadyUsed = errors.New("key already in use")
	ErrCountRestarted = errors.New("number 1 used. restart count")
)

type Counter struct {
	currentCountMap map[string]int
	current         int
}

func (cm *Counter) Play(playingUserId string, val int) error {
	err := cm.increment(playingUserId, val)
	if err != nil {
		cm.Reset()
		if err == ErrCountRestarted {
			cm.increment(playingUserId, val)
		}
	}
	return err
}

func (cm *Counter) increment(key string, value int) error {
	if value != cm.current {
		if value == 1 {
			return ErrCountRestarted
		}
		return ErrInvalidNumber
	}
	if _, ok := cm.currentCountMap[key]; ok {
		return ErrKeyAlreadyUsed
	}
	cm.current++
	cm.currentCountMap[key] = value
	return nil

}

func (cm *Counter) Reset() {
	cm.current = 1
	cm.currentCountMap = make(map[string]int, 100)
}

func NewCounter() *Counter {
	cm := &Counter{}
	cm.Reset()
	return cm
}
