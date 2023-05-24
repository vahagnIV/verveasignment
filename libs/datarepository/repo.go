package datarepository

import "time"

type MyObject struct {
	Id             string `gorm: "primaryKey"`
	Price          float64
	ExpirationDate time.Time
}

type Repo interface {
	get(id string) MyObject
	add(object MyObject)
}
