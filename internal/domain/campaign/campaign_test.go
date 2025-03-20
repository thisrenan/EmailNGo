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
	campaign := NewCampaign(name, content, contacts)

	//assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_IdIsNotNill(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	campaign := NewCampaign(name, content, contacts)

	//assert
	assert.NotNil(campaign.Id)
}

func Test_NewCampaign_CreatedDateNotNill(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	campaign := NewCampaign(name, content, contacts)

	//assert
	assert.NotNil(campaign.CreatedDate)
}
