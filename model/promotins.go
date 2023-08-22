package model

type Promotion struct {
	Id             string  `bson:"promotion_id,omitempty"`
	Price          float32 `bson:"price,omitempty"`
	ExpirationDate string  `bson:"expiration_date,omitempty"`
}
