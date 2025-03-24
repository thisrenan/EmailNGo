package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	Id          string    `validate:"required"`
	Name        string    `validate:"min=5,max=24"`
	CreatedDate time.Time `validate:"required"`
	Content     string    `validate:"min=5,max=1024"`
	Contacts    []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

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
