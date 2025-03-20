package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewcampaign(t *testing.T) {
	//arrange
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	//act
	campaign := NewCampaign(name, content, contacts)

	//assert
	assert.Equal(campaign.Id, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}
