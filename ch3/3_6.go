package ch3

import (
	"github.com/viodid/ctci-solutions-go/queue"
)

const (
	DOG = "DOG"
	CAT = "CAT"
)

type AnimalType string

type AnimalShelter struct {
	q *queue.Queue[AnimalType]
}

func NewAnimalShelter() *AnimalShelter {
	return &AnimalShelter{queue.NewQueue[AnimalType]()}
}

func (as *AnimalShelter) Enqueue(t AnimalType) {
	as.q.Add(t)
}

func (as *AnimalShelter) DequeueAny() (AnimalType, error) {
	return as.q.Remove()
}

func (as *AnimalShelter) DequeueCat() (AnimalType, error) {
	bufferQueue := queue.NewQueue[AnimalType]()

	for !as.q.IsEmpty() {
		animal, err := as.q.Peek()
		if err != nil {
			var zero AnimalType
			return zero, err
		}
		if animal == CAT {
			break
		}
		animal, err = as.q.Remove()
		if err != nil {
			var zero AnimalType
			return zero, err
		}
		bufferQueue.Add(animal)
	}

	cat, err := as.q.Remove()
	if err != nil {
		var zero AnimalType
		return zero, err
	}

	if err := pushQueue(bufferQueue, as.q); err != nil {
		var zero AnimalType
		return zero, err
	}

	return cat, nil
}

func (as *AnimalShelter) DequeueDog() (AnimalType, error) {
	bufferQueue := queue.NewQueue[AnimalType]()

	for !as.q.IsEmpty() {
		animal, err := as.q.Peek()
		if err != nil {
			var zero AnimalType
			return zero, err
		}
		if animal == DOG {
			break
		}
		animal, err = as.q.Remove()
		if err != nil {
			var zero AnimalType
			return zero, err
		}
		bufferQueue.Add(animal)
	}

	dog, err := as.q.Remove()
	if err != nil {
		var zero AnimalType
		return zero, err
	}

	if err := pushQueue(bufferQueue, as.q); err != nil {
		var zero AnimalType
		return zero, err
	}

	return dog, nil
}

func pushQueue(src, dst *queue.Queue[AnimalType]) error {
	for !src.IsEmpty() {
		val, err := src.Remove()
		if err != nil {
			return err
		}
		dst.Add(val)
	}
	return nil
}
