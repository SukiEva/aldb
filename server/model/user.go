package model

import (
	"errors"
	"github.com/SukiEva/aldb/server/model/database"
)

func AddUser(obj Operator) error {
	if mgo.ExistsOperator(obj.Email) {
		return errors.New("email exists")
	}
	err := mgo.InsertOperator(&database.Operator{
		Name:        obj.Name,
		Password:    obj.Password,
		Email:       obj.Email,
		Access:      obj.Access,
		Annotations: nil,
	})
	return err
}

func GetUser(email string) Operator {
	user, err := mgo.QueryOperatorByEmail(email)
	if err != nil {
		return Operator{}
	}
	return Operator{
		Name:     user.Name,
		Password: "",
		Email:    user.Email,
		Access:   user.Access,
	}
}
