package repo

import (
	"context"
	"storage/model"
)

type PromotionRepository interface {
	UploadPromotion(context.Context, []model.Promotion) error
	GetById(context.Context, string) (*model.Promotion, error)
	RemoveAllPromotions() error
}
