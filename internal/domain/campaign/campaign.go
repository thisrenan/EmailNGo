package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	Id          string
	Name        string
	CreatedDate time.Time
	Content     string
	Contacts    []Contact
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		Id:          "1",
		Name:        name,
		CreatedDate: time.Now(),
		Content:     content,
		Contacts:    contacts,
	}
}
