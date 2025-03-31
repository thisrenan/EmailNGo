package campaign

import (
	internalerrors "EmailNGo/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Id         string `gorm:"size:50"`
	Email      string `validate:"email"`
	CampaignId string `gorm:"size:50"`
}

const (
	Pending string = "Pending"
	Started string = "Started"
	Done    string = "Done"
)

type Campaign struct {
	Id          string    `validate:"required" gorm:"size:50"`
	Name        string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedDate time.Time `validate:"required"`
	Content     string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts    []Contact `validate:"min=1,dive"`
	Status      string    `gorm:"size:20"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].Id = xid.New().String()
	}

	campaign := &Campaign{
		Id:          xid.New().String(),
		Name:        name,
		CreatedDate: time.Now(),
		Content:     content,
		Contacts:    contacts,
		Status:      Pending,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err

}

func GetCampaign(ID int) (string, string, string, string, error) {
	return "", "", "", "", nil
}
