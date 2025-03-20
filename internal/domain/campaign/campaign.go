package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

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

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("Name is required")
	} else if content == "" {
		return nil, errors.New("Content is required")
	} else if len(emails) == 0 {
		return nil, errors.New("Emails is required")
	}

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		Id:          xid.New().String(),
		Name:        name,
		CreatedDate: time.Now(),
		Content:     content,
		Contacts:    contacts,
	}, nil
}
