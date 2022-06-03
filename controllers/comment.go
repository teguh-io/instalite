package controllers

import (
	"instalite/params"
	"instalite/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetCommentsByUserID(ctx *gin.Context)
	UpdateCommentByID(ctx *gin.Context)
	DeleteCommentByID(ctx *gin.Context)
}

type commentController struct {
	commentService services.CommentService
}

func NewCommentController(cs services.CommentService) CommentController {
	return &commentController{
		commentService: cs,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	var comment params.CreateCommentRequest
	err := ctx.ShouldBind(&comment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.CreateComment(userID, comment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (cc *commentController) GetCommentsByUserID(ctx *gin.Context) {
	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.GetCommentsByUserID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) UpdateCommentByID(ctx *gin.Context) {
	IDStr := ctx.Param("commentid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request params.UpdateCommentRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.UpdateCommentByID(ID, userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) DeleteCommentByID(ctx *gin.Context) {
	IDStr := ctx.Param("commentid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.DeleteCommentByID(ID, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
