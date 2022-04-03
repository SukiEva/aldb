package db

import (
	"context"
	"github.com/qiniu/qmgo"
	"sync"
)

type mdb struct {
	cli *qmgo.QmgoClient
	ctx context.Context
}

var (
	mgo  *mdb
	once = &sync.Once{}
)

func NewDB() *mdb {
	if mgo == nil {
		once.Do(func() {
			ctx := context.Background()
			client, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "algae"})
			if err != nil {
				panic("Fail to connect mongodb")
			}
			mgo = &mdb{
				cli: client,
				ctx: ctx,
			}
		})
	}
	return mgo
}
