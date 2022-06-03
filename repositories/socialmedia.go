package repositories

import (
	"instalite/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error)
	GetSocialMedias() ([]models.SocialMedia, error)
	UpdateSocialMediaByID(ID int, userID int, socialMedia models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMediaByID(ID int, userID int) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (smr *socialMediaRepository) CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	err := smr.db.Create(&socialMedia).Error
	return &socialMedia, err
}

func (smr *socialMediaRepository) GetSocialMedias() ([]models.SocialMedia, error) {
	panic("implement me")
}

func (smr *socialMediaRepository) UpdateSocialMediaByID(ID int, userID int, socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	panic("implement me")
}

func (smr *socialMediaRepository) DeleteSocialMediaByID(ID int, userID int) error {
	panic("implement me")
}
