package repositories

import (
	"errors"
	"instalite/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(photo models.Photo) (*models.Photo, error)
	GetPhotosByUserID(UserID int) ([]models.Photo, error)
	UpdatePhotoByID(ID int, photo models.Photo) (*models.Photo, error)
	DeletePhotoByID(ID int, userID int) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	db.AutoMigrate(&models.Photo{})
	return &photoRepository{
		db: db,
	}
}

func (pr *photoRepository) CreatePhoto(photo models.Photo) (*models.Photo, error) {
	err := pr.db.Create(&photo).Error
	return &photo, err
}

func (pr *photoRepository) GetPhotosByUserID(userID int) ([]models.Photo, error) {
	var photos []models.Photo
	err := pr.db.Model(&models.Photo{}).Select(
		"photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at, users.username, users.email",
	).Joins(
		"INNER JOIN users ON users.id = photos.user_id;", userID,
	).Scan(&photos).Error

	return photos, err

}

func (pr *photoRepository) UpdatePhotoByID(ID int, photo models.Photo) (*models.Photo, error) {
	var result models.Photo
	err := pr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(photo.UserID) {
			err := errors.New("you're forbidden to update this data")
			return err
		}

		err = tx.Where("id=?", ID).Updates(&photo).Find(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return &result, err
}

func (pr *photoRepository) DeletePhotoByID(ID int, userID int) error {
	var result models.Photo
	err := pr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(userID) {
			err := errors.New("you're forbidden to delete this data")
			return err
		}
		err = tx.Where("id=?", ID).Delete(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
