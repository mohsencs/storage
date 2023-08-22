package repo

import (
	"context"
	"storage/model"
	"storage/utils"
)

type repository struct {
	storage model.Storage
}

func NewPromotionRepository(strg model.Storage) PromotionRepository {
	r := repository{
		storage: strg,
	}

	return &r
}

func (r *repository) UploadPromotion(ctx context.Context, promotions []model.Promotion) error {
	err := r.storage.SaveBatch(ctx, promotions)
	if err != nil {
		utils.PrintError(err)
		return err
	}

	return nil
}

func (r *repository) RemoveAllPromotions() error {
	if err := r.storage.RemoveAll(); err != nil {
		return err
	}
	return nil
}

func (r *repository) GetById(ctx context.Context, id string) (model *model.Promotion, err error) {

	promotin, err := r.storage.GetPromotion(ctx, id)
	if err != nil {
		utils.PrintError(err)
		return nil, err
	}

	return promotin, nil
}
