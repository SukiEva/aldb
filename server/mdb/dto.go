package mdb

import "go.mongodb.org/mongo-driver/bson/primitive"

type Position struct {
	Longitude int `bson:"longitude"`
	Latitude  int `bson:"latitude"`
}

type River struct {
	Name     string `bson:"name"`
	Address  string `bson:"address"`
	Position `bson:"position"`
	AlgaeIds []primitive.ObjectID `bson:"algae-ids"`
}

type Algae struct {
	Name          string `bson:"name"`
	Position      `bson:"position"`
	Url           []string             `bson:"url"`
	RiverID       primitive.ObjectID   `bson:"river-id"`
	AnnotationIds []primitive.ObjectID `bson:"annotation-ids"`
}

type Operator struct {
	Name         string               `json:"name"`
	Password     string               `json:"password"`
	Email        string               `json:"email"`
	Access       int                  `json:"access"`
	OperationIds []primitive.ObjectID `json:"operation-ids"`
}

type Operation struct {
	CreateTime       int                  `json:"create-time"`
	UpdateTime       int                  `json:"update-time"`
	OperatorIds      []primitive.ObjectID `json:"operator-ids"`
	AnnotationID     primitive.ObjectID   `json:"annotation-id"`
	OldAnnotationIds []primitive.ObjectID `json:"old_annotation-ids"`
}

type Annotation struct {
	Description string             `json:"description"`
	Format      string             `json:"format"`
	Detail      string             `json:"detail"`
	AlgaeId     primitive.ObjectID `json:"algae-id"`
	OperationId primitive.ObjectID `json:"operation-id"`
}
