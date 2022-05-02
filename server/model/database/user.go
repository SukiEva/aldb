package database

import (
	"github.com/qiniu/qmgo/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operator struct {
	field.DefaultField `bson:",inline"`
	Name               string               `json:"name" bson:"name"`
	Password           string               `json:"password" bson:"password"`
	Email              string               `json:"email" bson:"email"`
	Access             int                  `json:"access" bson:"access"`
	Operations         []primitive.ObjectID `json:"operations" bson:"operations"`
}

type Operation struct {
	field.DefaultField `bson:",inline"`
	OperatorEmail      string               `json:"operatorEmail" bson:"operatorEmail"`
	Annotation         primitive.ObjectID   `json:"annotation" bson:"annotation"`
	OldAnnotations     []primitive.ObjectID `json:"oldAnnotations" bson:"oldAnnotations"`
}

func (m *Mgo) CheckOperator(email, password string) (Operator, error) {
	var one Operator
	err := operator.Find(ctx, bson.M{"email": email, "password": password}).One(&one)
	return one, err
}

func (m *Mgo) InsertOperator(obj *Operator) error {
	_, err := operator.InsertOne(ctx, obj)
	return err
}

func (m *Mgo) ExistsOperator(email string) bool {
	var one Operator
	if err := operator.Find(ctx, bson.M{"email": email}).One(&one); err != nil {
		return false
	}
	return true
}

func (m *Mgo) UpsertOperator(obj *Operator) error {
	_, err := operator.Upsert(ctx, bson.M{"_id": obj.Id}, obj)
	return err
}

func (m *Mgo) QueryOperations(email string) ([]Operation, error) {
	var batch []Operation
	err := operation.Find(ctx, bson.M{"operatorEmail": email}).All(&batch)
	return batch, err
}
