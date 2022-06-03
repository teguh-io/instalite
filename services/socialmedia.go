package services

import (
	"instalite/params"
	"instalite/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(request params.CreateSocialMediaRequest) (*params.CreateCommentResponse, error)
	GetSocialMedias() (*params.GetSocialMediaResponses, error)
	UpdateSocialMediaByID(ID int, userID int, request params.UpdateSocialMediaRequest) (*params.UpdateSocialMediaResponse, error)
	DeleteSocialMediaByID(ID, userID int) (*params.DeleteSocialMediaResponse, error)
}

type socialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepository
}

func NewSocialMediaService(smr repositories.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: smr,
	}
}

func (sms *socialMediaService) CreateSocialMedia(request params.CreateSocialMediaRequest) (*params.CreateCommentResponse, error) {
	panic("implement me")
}

func (sms *socialMediaService) GetSocialMedias() (*params.GetSocialMediaResponses, error) {
	panic("implement me")
}

func (sms *socialMediaService) UpdateSocialMediaByID(ID int, userID int, request params.UpdateSocialMediaRequest) (*params.UpdateSocialMediaResponse, error) {
	panic("implement me")
}

func (sms *socialMediaService) DeleteSocialMediaByID(ID, userID int) (*params.DeleteSocialMediaResponse, error) {
	panic("implement me")
}
