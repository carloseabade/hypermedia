package model

import (
	"math/rand"
)

type ContactsSet []Contact

func NewContactsSet() *ContactsSet {
	contactList := make(ContactsSet, 0)
	contactList.Add(NewContact("Carlos", "Abade", "+5542984392626", "carloseabade@gmail.com"))
	contactList.Add(NewContact("Clara", "Batista", "+5542999487560", "clarabatista@gmail.com"))
	contactList.Add(NewContact("Camily", "Abade", "+5542999713956", "camilyabade@gmail.com"))
	contactList.Add(NewContact("Daniel", "Lattre", "+5542123456789", "danieldelattre@gmail.com"))
	return &contactList
}

func (c ContactsSet) SearchByName(name string) ContactsSet {
	var result []Contact
	for _, contact := range c {
		if contact.First == name {
			result = append(result, contact)
		}
	}
	return result
}

func (c ContactsSet) All() ContactsSet {
	return c
}

func (c *ContactsSet) Add(contact *Contact) *ContactError {
	var ce ContactError
	if len(contact.First) < 1 {
		ce.First = "First name is required!"
	}
	if len(contact.Last) < 1 {
		ce.Last = "Last name is required!"
	}
	if len(contact.Phone) < 1 {
		ce.Phone = "Phone is required!"
	}
	if len(contact.Email) < 1 {
		ce.Email = "Email is required!"
	}
	if ce.First != "" || ce.Last != "" || ce.Phone != "" || ce.Email != "" {
		return &ce
	}

	*c = append(*c, *contact)
	return nil
}

type ContactError struct {
	First string
	Last  string
	Phone string
	Email string
}

type Contact struct {
	Id    string
	First string
	Last  string
	Phone string
	Email string
}

func NewContact(first, last, phone, email string) *Contact {
	return &Contact{
		Id:    getRandomId(),
		First: first,
		Last:  last,
		Phone: phone,
		Email: email,
	}
}

func getRandomId() string {
	numbers := []rune("0123456789")
	b := make([]rune, 12)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}
