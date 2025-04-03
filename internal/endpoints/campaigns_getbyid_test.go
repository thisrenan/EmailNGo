package endpoints

import (
	"EmailNGo/internal/contract"
	internalMock "EmailNGo/internal/test/internal-mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignResponse{
		Id:      "323",
		Name:    "Test",
		Content: "Hello!",
		Status:  "Pending",
	}
	service := new(internalMock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(res, req)

	assert.Equal(200, status)
	assert.Nil(campaign.Id, response.(*contract.CampaignResponse).Id)
	assert.Nil(campaign.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignGetById_should_return_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalMock.CampaignServiceMock)
	errExpected := errors.New("Something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(res, req)

	assert.Equal(errReturned.Error(), errExpected.Error())
}
