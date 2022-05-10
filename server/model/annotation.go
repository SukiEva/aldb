package model

import (
	"github.com/SukiEva/aldb/server/model/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAnnotationByAlga(algaName string) []Annotation {
	res := make([]Annotation, 0)
	alga, err := mgo.QueryAlgaByName(algaName)
	if err != nil {
		return res
	}
	for _, obj := range alga.Annotations {
		anno := mgo.QueryAnnotationById(obj)
		res = append(res, Annotation{
			Description: anno.Description,
			Format:      anno.Format,
			Url:         anno.Url,
			CreateAt:    anno.CreateAt.Format("2006-01-02 15:04"),
			UpdateAt:    anno.UpdateAt.Format("2006-01-02 15:04"),
		})
	}
	return res
}

func GetAnnotationByUser(userEmail string) []Annotation {
	res := make([]Annotation, 0)
	user, err := mgo.QueryOperatorByEmail(userEmail)
	if err != nil {
		return res
	}
	for _, obj := range user.Annotations {
		anno := mgo.QueryAnnotationById(obj)
		res = append(res, Annotation{
			Description: anno.Description,
			Format:      anno.Format,
			Url:         anno.Url,
			CreateAt:    anno.CreateAt.Format("2006-01-02 15:04"),
			UpdateAt:    anno.UpdateAt.Format("2006-01-02 15:04"),
			Id:          anno.Id.String(),
		})
	}
	return res
}

func AddAnnotation(obj Anno) (interface{}, error) {
	return mgo.InsertAnnotation(&database.Annotation{
		Description: obj.Description,
		Format:      obj.Format,
		Url:         obj.Url,
	})
}

func BindToAlga(algaName string, id primitive.ObjectID) error {
	alga, err := mgo.QueryAlgaByName(algaName)
	if err != nil {
		return err
	}
	var anno []primitive.ObjectID
	if alga.Annotations == nil {
		anno = []primitive.ObjectID{id}
	} else {
		anno = append(alga.Annotations, id)
	}
	err = mgo.UpdateAlga(alga.Id, anno)
	return err
}

func BindToUser(userEmail string, id primitive.ObjectID) error {
	user, err := mgo.QueryOperatorByEmail(userEmail)
	if err != nil {
		return err
	}
	var anno []primitive.ObjectID
	if user.Annotations == nil {
		anno = []primitive.ObjectID{id}
	} else {
		anno = append(user.Annotations, id)
	}
	err = mgo.UpdateOperator(user.Id, anno)
	return err
}
