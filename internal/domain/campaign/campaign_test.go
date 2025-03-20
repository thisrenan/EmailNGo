package campaign

import "testing"

func TestNewcampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}
	campaign := NewCampaign(name, content, contacts)

	if campaign.Id != "1" {
		t.Error("Expected 1", campaign.Id)
	} else if campaign.Name != name {
		t.Error("Expected correct name", campaign.Name)
	} else if campaign.Content != content {
		t.Error("Expected correct content", campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Error("Expected correct contacts", campaign.Contacts)
	}

}
