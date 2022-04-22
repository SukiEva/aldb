package model

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Position struct {
	Longitude int `json:"longitude" bson:"longitude"`
	Latitude  int `json:"latitude" bson:"latitude"`
}

type River struct {
	field.DefaultField `bson:",inline"`
	Name               string `json:"name" bson:"name"`
	Address            string `json:"address" bson:"address"`
	Position           `json:"position" bson:"position"`
	Algae              []primitive.ObjectID `json:"algae" bson:"algae"`
}

func (m *Mgo) QueryRivers() ([]River, error) {
	var batch []River
	err := river.Find(ctx, bson.M{}).All(&batch)
	return batch, err
}

func (m *Mgo) QueryRiver(name string) (River, error) {
	var one River
	err := river.Find(ctx, bson.M{"name": name}).One(&one)
	return one, err
}

func (m *Mgo) UpsertRiver(r River) error {
	_, err := river.Upsert(ctx, bson.M{"name": r.Name}, r)
	return err
}

func (m *Mgo) DropRiver(name string) error {
	err := river.Remove(ctx, bson.M{"name": name})
	return err
}
