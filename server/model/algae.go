package model

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alga struct {
	field.DefaultField `bson:",inline"`
	Name               string `json:"name" bson:"name"`
	Position           `json:"position" bson:"position"`
	Src                string               `json:"src" bson:"src"`
	River              primitive.ObjectID   `json:"river" bson:"river"`
	Annotations        []primitive.ObjectID `json:"annotations" bson:"annotations"`
}

func (m *Mgo) QueryAlgae() ([]Alga, error) {
	var batch []Alga
	err := algae.Find(ctx, bson.M{}).All(&batch)
	return batch, err
}

func (m *Mgo) QueryAlga(name string) (Alga, error) {
	var one Alga
	err := algae.Find(ctx, bson.M{"name": name}).One(&one)
	return one, err
}

func (m *Mgo) InsertAlga(obj Alga) error {
	_, err := algae.InsertOne(ctx, obj)
	return err
}

func (m *Mgo) UpsertAlga(obj Alga) error {
	_, err := algae.Upsert(ctx, bson.M{"_id": obj.Id}, obj)
	return err
}

func (m *Mgo) DropAlga(name string) error {
	err := algae.Remove(ctx, bson.M{"name": name})
	return err
}
