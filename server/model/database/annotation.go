package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Annotation struct {
	field.DefaultField `bson:",inline"`
	Description        string `json:"description" bson:"description"`
	Format             string `json:"format" bson:"format"`
	Url                string `json:"url" bson:"url"`
}

func (m *Mgo) QueryAnnotationById(obj primitive.ObjectID) *Annotation {
	var one *Annotation
	if err := annotation.Find(ctx, bson.M{"_id": obj}).One(&one); err != nil {
		return nil
	}
	return one
}

func (m *Mgo) UpsertAnnotation(obj *Annotation) error {
	_, err := annotation.Upsert(ctx, bson.M{"_id": obj.Id}, obj)
	return err
}

func (m *Mgo) InsertAnnotation(obj *Annotation) (interface{}, error) {
	res, err := annotation.InsertOne(ctx, obj)
	return res.InsertedID, err
}
