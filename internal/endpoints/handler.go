package endpoints

import "EmailNGo/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
