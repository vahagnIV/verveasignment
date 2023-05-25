package shardeddatabase

import (
	"someurl.com/datarepository"
	"sync"
)

type IShardingStrategy interface {
	GetShardIndex(id string) uint16
}

type ShardingContext struct {
	repos            []datarepository.Repo
	shardingStrategy IShardingStrategy
}

func Initialize(repos []datarepository.Repo, shardingFunction IShardingStrategy) ShardingContext {
	return ShardingContext{repos: repos, shardingStrategy: shardingFunction}
}

func (ctx ShardingContext) Get(id string) datarepository.DataRow {
	idx := ctx.shardingStrategy.GetShardIndex(id)
	return ctx.repos[idx].Get(id)
}

func (ctx ShardingContext) Add(object datarepository.DataRow) {
	idx := ctx.shardingStrategy.GetShardIndex(object.Id)
	ctx.repos[idx].Add(object)
}

func (ctx ShardingContext) BatchInsert(object []datarepository.DataRow) {
	// There should be a more elegant way for doing this, but
	// no time for finding it
	arr := [][]datarepository.DataRow{}
	for i := 0; i < len(ctx.repos); i++ {
		arr = append(arr, []datarepository.DataRow{})
	}

	for i := 0; i < len(object); i++ {
		idx := ctx.shardingStrategy.GetShardIndex(object[i].Id)
		arr[idx] = append(arr[idx], object[i])
	}
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for idx, elem := range arr {
		lidx := idx
		lelem := elem
		go func() {
			ctx.repos[lidx].BatchInsert(lelem)
			wg.Done()
		}()
	}
	wg.Wait()
}
