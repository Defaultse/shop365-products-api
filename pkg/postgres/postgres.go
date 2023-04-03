package postgres

import (
	"context"
	"fmt"
	"hash/fnv"
	"shop365-products-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ShardNum int
type shardMap map[ShardNum]*gorm.DB

const ShardQuantity = 2

const (
	Shard1 ShardNum = iota + 1
	Shard2
)

type Postgres struct {
	ShardMap shardMap
}

func NewPostgres(ctx context.Context, configs config.PG) (*Postgres, error) {
	postgres := &Postgres{
		ShardMap: initShardMap(ctx, configs),
	}

	return postgres, nil
}

func initShardMap(ctx context.Context, configs config.PG) shardMap {
	var m = shardMap{
		Shard1: discoveryShard(ctx, configs.URL),
		Shard2: discoveryShard(ctx, configs.URL2),
	}

	return m
}

func discoveryShard(ctx context.Context, dsn string) *gorm.DB {
	pgClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := pgClient.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(5)

	return pgClient
}

func GetShardIDFromHash(s string) int64 {
	h := fnv.New64()
	h.Write([]byte(s))

	selectedShard := h.Sum64()%ShardQuantity + 1

	fmt.Println("----Selected shard num:", selectedShard)

	return int64(selectedShard)
}
