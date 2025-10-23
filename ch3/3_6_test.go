package ch3

import (
	"testing"
)

func TestAnimalShelter(t *testing.T) {
	shelter := NewAnimalShelter()

	animals := []AnimalType{DOG, CAT, CAT, DOG, CAT}
	for _, a := range animals {
		shelter.Enqueue(a)
	}

	animal, err := shelter.DequeueAny()
	if err != nil {
		t.Fatal(err)
	}
	if animal != DOG {
		t.Errorf("DequeueAny wrong animal. expected=%s, got=%s",
			DOG, animal)
	}

	animal, err = shelter.DequeueDog()
	if err != nil {
		t.Fatal(err)
	}
	if animal != DOG {
		t.Errorf("DequeueAny wrong animal. expected=%s, got=%s",
			DOG, animal)
	}

	animal, err = shelter.DequeueAny()
	if err != nil {
		t.Fatal(err)
	}
	if animal != CAT {
		t.Errorf("DequeueAny wrong animal. expected=%s, got=%s",
			CAT, animal)
	}

	for i := 0; i < 2; i++ {
		animal, err = shelter.DequeueCat()
		if err != nil {
			t.Fatal(err)
		}
		if animal != CAT {
			t.Errorf("DequeueAny wrong animal. expected=%s, got=%s",
				CAT, animal)
		}
	}

}
