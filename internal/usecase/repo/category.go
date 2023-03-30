package repo

import (
	"context"
	"fmt"

	"shop365-products-api/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepo struct {
	db *mongo.Client
}

func NewCategoryRepo(db *mongo.Client) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (cr *CategoryRepo) GetAll() (*[]entity.Category, error) {
	var categories []entity.Category

	coll := cr.db.Database("shop365").Collection("product_category")
	filter := bson.D{{}}

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = cursor.All(context.TODO(), &categories)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &categories, nil
}

func (cr *CategoryRepo) GetByID() (*entity.Category, error) {
	var category *entity.Category

	coll := cr.db.Database("shop365").Collection("product_category")

	// filter := mongo.

	err := coll.FindOne(context.TODO(), nil).Decode(&category)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return category, nil
}
