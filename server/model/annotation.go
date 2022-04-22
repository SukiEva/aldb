package model

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Annotation struct {
	field.DefaultField `bson:",inline"`
	Description        string             `json:"description" bson:"description"`
	Format             string             `json:"format" bson:"format"`
	Detail             string             `json:"detail" bson:"detail"`
	Alga               primitive.ObjectID `json:"alga" bson:"alga"`
}

func (m *Mgo) QueryAnnotations(obj primitive.ObjectID) ([]Annotation, error) {
	var batch []Annotation
	err := annotation.Find(ctx, bson.M{"alga": obj}).All(&batch)
	return batch, err
}

func (m *Mgo) QueryAnnotation(obj primitive.ObjectID) (Annotation, error) {
	var one Annotation
	err := annotation.Find(ctx, bson.M{"_id": obj}).One(&one)
	return one, err
}

func (m *Mgo) UpsertAnnotation(obj Annotation) error {
	_, err := annotation.Upsert(ctx, bson.M{"_id": obj.Id}, obj)
	return err
}

func (m *Mgo) InsertAnnotation(obj Annotation) error {
	_, err := annotation.InsertOne(ctx, obj)
	return err
}
