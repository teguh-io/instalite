package controllers

import (
	"instalite/services"

	"github.com/gin-gonic/gin"
)

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	UpdateSocialMediaByID(ctx *gin.Context)
	DeleteSocialMediaByID(ctx *gin.Context)
}

type socialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(sms services.SocialMediaService) SocialMediaController {
	return &socialMediaController{
		socialMediaService: sms,
	}
}

func (smc *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	panic("implement me")
}

func (smc *socialMediaController) GetSocialMedias(ctx *gin.Context) {
	panic("implement me")
}

func (smc *socialMediaController) UpdateSocialMediaByID(ctx *gin.Context) {
	panic("implement me")
}

func (smc *socialMediaController) DeleteSocialMediaByID(ctx *gin.Context) {
	panic("implement me")
}
