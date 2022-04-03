package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type River struct {
	Name     string `bson:"name"`
	Position `bson:"position"`
	AlgaeIds []primitive.ObjectID `bson:"algae_ids"`
}
