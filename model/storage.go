package model

import "context"

type Storage interface {
	SaveBatch(context.Context, []Promotion) error
	RemoveAll() error
	GetPromotion(context.Context, string) (*Promotion, error)
}
