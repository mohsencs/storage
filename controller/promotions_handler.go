package controller

import (
	"errors"
	"fmt"
	"net/http"
	"storage/model"
	"storage/service"
	"storage/utils"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.PromotionService
}

func NewPromotionController(rg *gin.RouterGroup, ps service.PromotionService) {
	h := &handler{service: ps}

	rg.POST("upload", h.upload)
	rg.GET(":id", h.getPromotin)
}

func (h *handler) upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		utils.PrintError(err)
		return
	}

	h.service.UploadPromotion(c, f)

	c.Status(http.StatusOK)
}

func (h *handler) getPromotin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, errors.New("id not found"))
		return
	}

	promotion, err := h.service.GetPromotionById(c, id)
	if err != nil {
		utils.PrintError(err)
		return
	}

	fmt.Printf("aaaaaaaaaaaaaaaa %v", promotion)

	c.JSON(http.StatusOK, toResponse(promotion))
}

type response struct {
	Id             string  `json:"id"`
	Price          float32 `json:"price"`
	ExpirationDate string  `json:"expiration_date"`
}

func toResponse(p *model.Promotion) *response {
	return &response{
		Id:             p.Id,
		Price:          p.Price,
		ExpirationDate: p.ExpirationDate,
	}
}
