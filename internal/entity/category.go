package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	SubCategories []SubCategory      `bson:"sub_categories"`
}

type SubCategory struct {
	ID   primitive.ObjectID `bson:"id"`
	Name string             `bson:"name"`
}
