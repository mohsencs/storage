package model

import (
	"context"
	"log"
	"os"
	"storage/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoStorage() (Storage, error) {
	uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("MONGO_DATABASE")
	coll_name := os.Getenv("MONGO_COLLECTION")
	if uri == "" || db_name == "" || coll_name == "" {
		log.Fatal("Missed some mongo enviroment values. uri: ", uri, ". db: ", db_name, ". collection: ", coll_name)
	}

	client, error := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if error != nil {
		return nil, error
	}

	coll := client.Database(db_name).Collection(coll_name)

	return &mongoStorage{
		client: client,
		coll:   coll,
	}, nil
}

func (ms *mongoStorage) SaveBatch(ctx context.Context, promotions []Promotion) error {

	var interfaceSlice []interface{} = make([]interface{}, len(promotions))
	for i, d := range promotions {
		interfaceSlice[i] = d
	}

	result, err := ms.coll.InsertMany(ctx, interfaceSlice)
	if err != nil {
		utils.PrintError(err)
		return err
	}

	log.Printf("Saved batch of promotions in mongo. Inserted %v items to db, from %v inputs.", len(result.InsertedIDs), len(promotions))

	return nil
}

func (ms *mongoStorage) RemoveAll() error {

	result, err := ms.coll.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		utils.PrintError(err)
		return err
	}

	log.Printf("Remove all promotions. Count of deleted items: %v ", result.DeletedCount)
	return nil
}

func (ms *mongoStorage) GetPromotion(ctx context.Context, id string) (*Promotion, error) {
	var promotion Promotion

	if err := ms.coll.FindOne(ctx, bson.D{{"promotion_id", id}}).Decode(&promotion); err != nil {
		utils.PrintError(err)
		return nil, err
	}

	return &promotion, nil
}
