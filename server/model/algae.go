package model

import (
	"errors"
	"github.com/SukiEva/aldb/server/model/database"
)

func GetData() []Alga {
	res := make([]Alga, 0)
	algae, err := mgo.QueryAlgae()
	if err != nil {
		return res
	}
	for _, obj := range algae {
		river := mgo.QueryRiverById(obj.River)
		res = append(res, Alga{
			Name:  obj.Name,
			Src:   obj.Src,
			River: river.Name,
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
