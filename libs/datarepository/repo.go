package datarepository

import "time"

type DataRow struct {
	Id             string `gorm: "primaryKey"`
	Price          float64
	ExpirationDate time.Time
}

type Repo interface {
	Get(id string) DataRow
	Add(object DataRow)
	BatchInsert(object []DataRow)
}
