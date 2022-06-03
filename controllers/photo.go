package controllers

import (
	"instalite/params"
	"instalite/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	UpdatePhotoByID(ctx *gin.Context)
	DeletePhotoByID(ctx *gin.Context)
}

type photoController struct {
	photoService services.PhotoService
}

func NewPhotoController(ps services.PhotoService) PhotoController {
	return &photoController{
		photoService: ps,
	}
}

func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	photo := params.CreatePhotoRequest{}
	err := ctx.ShouldBind(&photo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.CreatePhoto(userID, photo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (pc *photoController) GetPhotos(ctx *gin.Context) {
	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.GetPhotosByUserID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (pc *photoController) UpdatePhotoByID(ctx *gin.Context) {
	IDStr := ctx.Param("photoid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request params.UpdatePhotoRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.UpdatePhotoByID(ID, userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)

}

func (pc *photoController) DeletePhotoByID(ctx *gin.Context) {
	IDStr := ctx.Param("photoid")
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.DeletePhotoByID(ID, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
