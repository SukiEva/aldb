package model

import (
	"errors"
	"github.com/SukiEva/aldb/server/model/database"
)

func GetData() []Alga {
	algae, err := mgo.QueryAlgae()
	if err != nil {
		return nil
	}
	res := make([]Alga, 0)
	for _, obj := range algae {
		river := mgo.QueryRiverById(obj.River)
		annotations := make([]Annotation, 0)
		for _, aid := range obj.Annotations {
			anno := mgo.QueryAnnotation(aid)
			annotations = append(annotations, Annotation{
				Description: anno.Description,
				Format:      anno.Format,
				Url:         anno.Url,
			})
		}
		res = append(res, Alga{
			Name:        obj.Name,
			Src:         obj.Src,
			River:       river.Name,
			Annotations: annotations,
		})
	}
	return res
}

func AddAlga(obj Alga) error {
	river := mgo.QueryRiverByName(obj.River)
	err := mgo.InsertAlga(&database.Alga{
		Name:        obj.Name,
		Src:         obj.Src,
		River:       river.Id,
		Annotations: nil,
	})
	return err
}

func AddRiver(obj River) error {
	if mgo.ExistsRiver(obj.Name) {
		return errors.New("river exists")
	}
	err := mgo.InsertRiver(&database.River{
		Name:    obj.Name,
		Address: obj.Address,
	})
	return err
}

func GetRivers() []River {
	rivers, err := mgo.QueryRivers()
	if err != nil {
		return nil
	}
	res := make([]River, 0)
	for _, obj := range rivers {
		res = append(res, River{
			Name:    obj.Name,
			Address: obj.Address,
		})
	}
	return res
}
