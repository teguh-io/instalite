package services

import (
	"instalite/models"
	"instalite/params"
	"instalite/repositories"
)

type CommentService interface {
	CreateComment(userID int, request params.CreateCommentRequest) (*params.CreateCommentResponse, error)
	GetCommentsByUserID(userID int) ([]params.GetCommentResponse, error)
	UpdateCommentByID(ID int, userID int, request params.UpdateCommentRequest) (*params.UpdateCommentResponse, error)
	DeleteCommentByID(ID, userID int) (*params.DeleteCommentResponse, error)
}

type commentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentRepository(cr repositories.CommentRepository) CommentService {
	return &commentService{
		commentRepo: cr,
	}
}

func toGetCommentResponse(commentModel models.Comment) params.GetCommentResponse {
	user := params.User{
		ID:       &commentModel.UsersID,
		Username: commentModel.UsersName,
		Email:    commentModel.UsersEmail,
	}

	photo := params.Photo{
		ID:       commentModel.PhotosID,
		Title:    commentModel.PhotosTitle,
		Caption:  commentModel.PhotosCaption,
		PhotoUrl: commentModel.PhotosUrl,
		UserID:   commentModel.PhotosUserID,
	}

	return params.GetCommentResponse{
		ID:        commentModel.ID,
		Message:   commentModel.Message,
		PhotoID:   commentModel.PhotoID,
		UserID:    commentModel.UserID,
		CreatedAt: commentModel.CreatedAt,
		UpdatedAt: commentModel.UpdatedAt,
		User:      user,
		Photo:     photo,
	}
}

func toGetCommentResponses(commentModels []models.Comment) []params.GetCommentResponse {
	getCommentResponses := make([]params.GetCommentResponse, len(commentModels))
	for idx, commentModel := range commentModels {
		getCommentResponses[idx] = toGetCommentResponse(commentModel)
	}

	return getCommentResponses
}

func (cs *commentService) CreateComment(userID int, request params.CreateCommentRequest) (*params.CreateCommentResponse, error) {
	commentModel := models.Comment{
		Message: request.Message,
		PhotoID: request.PhotoID,
		UserID:  userID,
	}

	res, err := cs.commentRepo.CreateComment(commentModel)
	if err != nil {
		return nil, err
	}

	createCommentResponse := params.CreateCommentResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}

	return &createCommentResponse, nil

}

func (cs *commentService) GetCommentsByUserID(userID int) ([]params.GetCommentResponse, error) {
	res, err := cs.commentRepo.GetCommentsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return toGetCommentResponses(res), nil
}

func (cs *commentService) UpdateCommentByID(ID int, userID int, request params.UpdateCommentRequest) (*params.UpdateCommentResponse, error) {
	commentModel := models.Comment{
		Message: request.Message,
	}

	res, err := cs.commentRepo.UpdateCommentByID(ID, userID, commentModel)
	if err != nil {
		return nil, err
	}

	updateCommentResponse := params.UpdateCommentResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoUrl:  res.PhotosUrl,
		UserID:    res.UserID,
		UpdatedAt: res.UpdatedAt,
	}

	return &updateCommentResponse, nil
}

func (cs *commentService) DeleteCommentByID(ID, userID int) (*params.DeleteCommentResponse, error) {
	err := cs.commentRepo.DeleteCommentByID(ID, userID)
	if err != nil {
		return nil, err
	}

	deleteCommentResponse := params.DeleteCommentResponse{
		Message: "Your comment has been successfully deleted",
	}

	return &deleteCommentResponse, err
}
