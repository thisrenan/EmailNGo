package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_Newcampaign_CreateCampaign(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	campaign, _ := NewCampaign(name, content, contacts)

	//assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_IdIsNotNill(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	campaign, _ := NewCampaign(name, content, contacts)

	//assert
	assert.NotNil(campaign.Id)
}

func Test_NewCampaign_CreatedDateNotNill(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	campaign, _ := NewCampaign(name, content, contacts)

	//assert
	assert.NotNil(campaign.CreatedDate)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign("", content, contacts)

	//assert
	assert.Equal("Name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign(name, "", contacts)

	//assert
	assert.Equal("Content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign(name, content, []string{})

	//assert
	assert.Equal("Emails is required", err.Error())
}
