package sqlrepository

import (
	"gorm.io/gorm"
	"someurl.com/datarepository"
)

type SqlData struct {
	db *gorm.DB
}

func Initialize(dialector gorm.Dialector, opts ...gorm.Option) (SqlData, error) {

	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return SqlData{}, err
	}

	err = db.AutoMigrate(&datarepository.DataRow{})
	if err != nil {
		return SqlData{}, err
	}

	result := SqlData{db: db}
	return result, err
}

func (d SqlData) Get(id string) datarepository.DataRow {
	var b datarepository.DataRow
	d.db.Model(&datarepository.DataRow{}).First(&b, "id = ?", id)
	return b
}

func (d SqlData) Add(dataRow datarepository.DataRow) {
	d.db.Model(&datarepository.DataRow{}).Create(dataRow)
}

func (d SqlData) BatchInsert(dataRow []datarepository.DataRow) {
	for _, element := range dataRow {
		d.Add(element)
	}
}
