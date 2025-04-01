package campaign

import (
	"EmailNGo/internal/contract"
	internalerrors "EmailNGo/internal/internal-errors"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
	Cancel(id string) error
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.Id, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {

	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.CampaignResponse{
		Id:      campaign.Id,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}

func (s *ServiceImp) Cancel(id string) error {

	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return internalerrors.ErrInternal
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Cancel()
	err = s.Repository.Save(campaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
