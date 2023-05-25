module importer

go 1.20

require (
	someurl.com/datarepository v0.0.0-00010101000000-000000000000
	someurl.com/factorymethods v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	gorm.io/driver/sqlite v1.5.1 // indirect
	gorm.io/gorm v1.25.1 // indirect
	someurl.com/shardeddatabase v0.0.0-00010101000000-000000000000 // indirect
	someurl.com/sqlrepository v0.0.0-00010101000000-000000000000 // indirect

)

replace someurl.com/datarepository => ../libs/datarepository

replace someurl.com/factorymethods => ../libs/factorymethods

replace someurl.com/sqlrepository => ../libs/sqlrepository

replace someurl.com/shardeddatabase => ../libs/shardeddatabase
