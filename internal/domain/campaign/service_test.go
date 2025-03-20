package campaign

import (
	"EmailNGo/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test X",
		Content: "Body",
		Emails:  []string{"test1@test.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test X",
		Content: "Body",
		Emails:  []string{"test1@test.com"},
	}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}
