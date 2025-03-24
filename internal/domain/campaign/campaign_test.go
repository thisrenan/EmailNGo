package campaign

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body Hi!"
	contacts = []string{"email1@e.com", "email2@e.com"}
	fake     = faker.New()
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

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign("", content, contacts)

	//assert
	assert.Equal("Name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	//assert
	assert.Equal("Name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin5(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign(name, "", contacts)

	//assert
	assert.Equal("Content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	//assert
	assert.Equal("Content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	_, err := NewCampaign(name, content, []string{"email_invalid"})

	//assert
	assert.Equal("Email is invalid", err.Error())
}
