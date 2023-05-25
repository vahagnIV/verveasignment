module someurl.com/sqlrepository

go 1.20

replace someurl.com/datarepository => ../datarepository

require (
	gorm.io/driver/sqlite v1.5.1
	gorm.io/gorm v1.25.1
	someurl.com/datarepository v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
)
