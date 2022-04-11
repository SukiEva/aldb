package main

import (
	"fmt"
	"github.com/SukiEva/aldb/server/mdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	obj := &mdb.Algae{
		Name:          "拉拉",
		Position:      mdb.Position{},
		Url:           nil,
		RiverID:       primitive.ObjectID{},
		AnnotationIds: nil,
	}
	err := obj.Insert()
	if err != nil {
		fmt.Println("错误")
	}
}
