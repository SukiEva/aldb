package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Algae struct {
	Name          string `bson:"name"`
	Position      `bson:"position"`
	Url           string               `bson:"url"`
	RiverID       primitive.ObjectID   `bson:"river_id"`
	AnnotationIds []primitive.ObjectID `bson:"annotation_ids"`
}
