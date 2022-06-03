package repositories

import (
	"errors"
	"fmt"
	"instalite/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(request models.Comment) (*models.Comment, error)
	GetCommentsByUserID(userID int) ([]models.Comment, error)
	UpdateCommentByID(ID int, userID int, request models.Comment) (*models.Comment, error)
	DeleteCommentByID(ID, userID int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	db.AutoMigrate(models.Comment{})
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) CreateComment(comment models.Comment) (*models.Comment, error) {
	fmt.Println(comment)
	err := cr.db.Create(&comment).Error
	return &comment, err
}

func (cr *commentRepository) GetCommentsByUserID(userID int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `
	comments.id,
	comments.message,
	comments.user_id,
	comments.updated_at,
	comments.created_at,
	photos.id as photos_id,
	photos.caption as photos_caption,
	photos.title as photos_title,
	photos.user_id as photos_user_id,
	photos.photo_url as photos_url,
	users.id as users_id,
	users.username as users_name,
	users.email as users_email`

	err := cr.db.Model(&models.Comment{}).Select(query).Joins(
		"INNER JOIN photos ON comments.photo_id = photos.id INNER JOIN users ON comments.user_id = users.id;", userID,
	).Scan(&comments).Error

	return comments, err
}

func (cr *commentRepository) UpdateCommentByID(ID int, userID int, request models.Comment) (*models.Comment, error) {
	var result models.Comment
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			fmt.Println("1")
			return err
		}

		if result.UserID != userID {
			err := errors.New("you're forbidden to update this data")
			return err
		}

		err = tx.Where("id=?", ID).Updates(&request).Error
		if err != nil {
			fmt.Println("2")
			return err
		}

		query := `
		comments.id,
		comments.message,
		comments.user_id,
		comments.updated_at,
		photos.photo_url as photos_url		
		`
		err = tx.Model(&models.Comment{}).Select(query).Joins("INNER JOIN photos on comments.photo_id = photos.id;", ID).Scan(&result).Error
		if err != nil {
			fmt.Println("3")
			return err
		}

		return nil
	})

	return &result, err
}

func (cr *commentRepository) DeleteCommentByID(ID, userID int) error {
	var result models.Comment
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result).Error
		if err != nil {
			return err
		}

		if result.UserID != userID {
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
