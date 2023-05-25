module someurl.com/factorymethods

go 1.20

replace someurl.com/datarepository => ../datarepository

replace someurl.com/sqlrepository => ../sqlrepository

replace someurl.com/shardeddatabase => ../shardeddatabase

require (
	gorm.io/driver/sqlite v1.5.1
	gorm.io/gorm v1.25.1
	someurl.com/datarepository v0.0.0-00010101000000-000000000000
	someurl.com/shardeddatabase v0.0.0-00010101000000-000000000000
	someurl.com/sqlrepository v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
)
