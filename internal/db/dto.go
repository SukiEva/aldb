package db

type Position struct {
	Longitude int `bson:"longitude"`
	Latitude  int `bson:"latitude"`
}
