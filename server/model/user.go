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
		Name:       obj.Name,
		Password:   obj.Password,
		Email:      obj.Email,
		Access:     obj.Access,
		Operations: nil,
	})
	if err != nil {
		return err
	}
	return nil
}
