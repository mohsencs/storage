package service

import (
	"mime/multipart"
	"storage/model"

	"github.com/gin-gonic/gin"
)

type PromotionService interface {
	UploadPromotion(ctx *gin.Context, file *multipart.FileHeader)
	GetPromotionById(ctx *gin.Context, id string) (*model.Promotion, error)
}
