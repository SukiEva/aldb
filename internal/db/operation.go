package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Operator struct {
	Name         string               `json:"name"`
	OperationIds []primitive.ObjectID `json:"operation_ids"`
}

type Operation struct {
	CreateTime       int                  `json:"create_time"`
	UpdateTime       int                  `json:"update_time"`
	OperatorIds      []primitive.ObjectID `json:"operator_ids"`
	AnnotationID     primitive.ObjectID   `json:"annotation_id"`
	OldAnnotationIds []primitive.ObjectID `json:"old_annotation_ids"`
}
