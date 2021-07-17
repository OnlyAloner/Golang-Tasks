package structs

type Contact struct {
	ID int
	Name,
	Phone,
	Firstname,
	Lastname,
	Email,
	CreatedAt string
}
type ContactManager struct {
	Contacts []Contact
}

func NewContactManager() *ContactManager {
	Temp := new(ContactManager)
	return Temp
}

func (This *ContactManager) create(newContact Contact) *Contact {

	for ind, elem := range This.Contacts {
		if elem.ID == newContact.ID {
			return &This.Contacts[ind]
		}
	}

	This.Contacts = append(This.Contacts, newContact)

	return &This.Contacts[len(This.Contacts)-1]

	//func returning the pointer to element that equal to newContact in Contacts
}

func (This *ContactManager) update(updContact Contact) *Contact {
	for ind, elem := range This.Contacts {
		if elem.ID == updContact.ID {
			This.Contacts[ind] = updContact
			return &This.Contacts[ind]
		}
	}

	return &updContact // If contact not finded we will return the pointer to updContact
}

func (This *ContactManager) delete(delContact Contact) *Contact {
	for ind, elem := range This.Contacts {
		if elem.ID == delContact.ID {

			This.Contacts = append(This.Contacts[:ind], This.Contacts[ind+1:]...)
			return &delContact // If contact not finded or he was deleted we will return the pointer to delContact

		}
	}

	return &delContact // If contact not finded or he was deleted we will return the pointer to delContact
}

func (This *ContactManager) get(ID int) *Contact {
	for ind, elem := range This.Contacts {
		if elem.ID == ID {
			return &This.Contacts[ind]
		}
	}

	return &Contact{} // If we find we return the pointer to element, otherwise pointer to empty Contact
}

func (This *ContactManager) GetAll() *[]Contact {
	return &This.Contacts
}
