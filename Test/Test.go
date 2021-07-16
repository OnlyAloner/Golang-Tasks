package main

import "fmt"

type Contact struct {
	ID int
	Name,
	Phone,
	Firstname,
	Lastname,
	Email,
	CreateadAt string
}

type ContactManager struct {
	Contacts []Contact
}

func NewContactManager() *ContactManager {
	Temp := new(ContactManager)
	return Temp
}
func (This *ContactManager) get(ID int) *Contact {
	for ind, elem := range This.Contacts {
		if elem.ID == ID {
			return &This.Contacts[ind]
		}
	}

	return &Contact{}
}

func main() {
	A := NewContactManager()
	fmt.Print(A.get(13))

}
