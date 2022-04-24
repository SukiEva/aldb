package model

func GetData() []Alga {
	algae, err := mgo.QueryAlgae()
	if err != nil {
		return nil
	}
	res := make([]Alga, 0)
	for _, obj := range algae {
		river := mgo.QueryRiver(obj.River)
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
			Name: obj.Name,
			Position: Position{
				Longitude: obj.Longitude,
				Latitude:  obj.Latitude,
			},
			Src: obj.Src,
			River: River{
				Name:    river.Name,
				Address: river.Address,
				Position: Position{
					Longitude: river.Longitude,
					Latitude:  river.Latitude,
				},
			},
			Annotations: annotations,
		})
	}
	return res
}
