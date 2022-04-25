package model

type Position struct {
	Longitude int `json:"longitude" bson:"longitude"`
	Latitude  int `json:"latitude" bson:"latitude"`
}

type River struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Position `json:"position"`
}

type Alga struct {
	Name        string `json:"name"`
	Position    `json:"position"`
	Src         string `json:"src"`
	River       `json:"river"`
	Annotations []Annotation `json:"annotations"`
}

type Annotation struct {
	Description string `json:"description"`
	Format      string `json:"format"`
	Url         string `json:"url"`
}

type Operator struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Access   int    `json:"access" binding:"-"`
}
