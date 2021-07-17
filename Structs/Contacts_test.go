package structs

import (
	"testing"
	"time"
)

var (
	Persons = []Contact{
		Contact{
			ID:        0,
			Name:      "Clown",
			Phone:     "+998937853311",
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        1,
			Name:      "Kale",
			Phone:     "+998937853311",
			Firstname: "Archi",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        2,
			Name:      "Charli",
			Phone:     "+009813239120",
			Firstname: "Charles",
			Lastname:  "Dickens",
			Email:     "IG@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
	}
	TestsForContactUpdate = []Contact{
		Contact{
			ID:        0,
			Name:      "Wick",
			Phone:     "+998937853311",
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        1,
			Name:      "Ocean",
			Phone:     "+998937853311",
			Firstname: "Archi",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        2,
			Name:      "Charli",
			Phone:     "+009813239120",
			Firstname: "Charles",
			Lastname:  "Dickens",
			Email:     "IG@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
	}

	p = NewContactManager()
)

func TestCreate(t *testing.T) {

	for _, ppl := range Persons {

		Passed := true

		temp := &Contact{}

		output := p.create(ppl)

		if output == temp {
			Passed = false
		} else if *output != ppl {
			Passed = false
		}

		if !Passed { // If "create" function returned the empty pointer or pointer to element that not equal to the our Contact we will return the error
			t.Error("Failed to Create object: ", ppl)
		}
	}

}

func TestUpdate(t *testing.T) {

	//TestCreate(t) //We will be able to test each function separately, without the need to check the entire test

	for _, ppl := range TestsForContactUpdate {

		output := p.update(ppl)

		if output == &ppl {
			t.Error("Failed to Update object cause we cant find the contact with this ID: ", ppl)
		} else if *output != ppl {
			t.Error("Failed to Update object or slice was changed", ppl)
		}

	}

}
func TestDelete(t *testing.T) {

	//TestCreate(t) //We will be able to test each function separately, without the need to check the entire test

	for _, ppl := range TestsForContactUpdate {

		output := p.delete(&ppl)

		if output != &ppl {
			t.Error("Failed to Delete object: ", ppl)
		}
	}

}

func TestGet(t *testing.T) {

	//TestCreate(t) //We will be able to test each function separately, without the need to check the entire test

	output := p.get(1)
	emptyPointer := &Contact{}
	if output == emptyPointer {
		t.Error("Failed to Get contact with id: 1. Because object was not found")
	} else if output.ID != 1 {
		t.Error("Failes to Get contact with id: 1.")
	}

}
func TestGetAll(t *testing.T) {

	//TestCreate(t) //We will be able to test each function separately, without the need to check the entire test

	if p.GetAll() != &p.Contacts {
		t.Error("failed to get the list of contacts")
	}
}
