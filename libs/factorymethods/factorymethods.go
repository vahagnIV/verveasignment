package factorymethods

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hash/fnv"
	"os"
	"path"
	"someurl.com/datarepository"
	"someurl.com/shardeddatabase"
	"someurl.com/sqlrepository"
)

const ServerUpdateUrl string = "http://127.0.0.1:3333/updateDatabase/"
const DatabasePath string = "./data/db/"
const shardCount uint16 = 100

func GetDatabaseFileTemplateFromTimestamp(timestamp string) string {
	db_dir := path.Join(path.Dir(DatabasePath), timestamp)
	return path.Join(db_dir, "gorm")
}

type HashShardingStrategy struct {
	shardCount uint16
}

func (str HashShardingStrategy) GetShardIndex(id string) uint16 {
	h := fnv.New32a()
	h.Write([]byte(id))
	return uint16(h.Sum32() % uint32(str.shardCount))
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func InitShardedDb(filename string, createIfNotExist bool) (shardeddatabase.ShardingContext, error) {
	repos := []datarepository.Repo{}
	hashingStrategy := HashShardingStrategy{shardCount: shardCount}
	for i := uint16(0); i < shardCount; i++ {
		f := fmt.Sprintf("%s%d.db", filename, i)

		if !createIfNotExist && !FileExists(f) {
			return shardeddatabase.ShardingContext{}, errors.New("database does not exist or incomplete")
		}
		sqlrepo, err := sqlrepository.Initialize(sqlite.Open(f), &gorm.Config{})
		if err != nil {
			return shardeddatabase.ShardingContext{}, err
		}
		repos = append(repos, sqlrepo)
	}
	return shardeddatabase.Initialize(repos, hashingStrategy), nil
}

func InitSqlShard(filename string) (datarepository.Repo, error) {
	return sqlrepository.Initialize(sqlite.Open(filename), &gorm.Config{})
}
