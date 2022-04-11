package mdb

import (
	"context"
	"fmt"
	"github.com/SukiEva/aldb/server/config"
	"github.com/SukiEva/aldb/server/log"
	"github.com/qiniu/qmgo"
)

var (
	cli *qmgo.Client
	mdb *qmgo.Database
	ctx context.Context
	cfg = config.GetEnv()

	algae *qmgo.Collection
	river *qmgo.Collection
)

func init() {
	var err error
	ctx = context.Background()
	cli, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: cfg.Mongo.Uri})
	if err != nil {
		panic("Fail to connect mongodb")
	}
	mdb = cli.Database(cfg.Mongo.DB)
	algae = mdb.Collection("algae")
	river = mdb.Collection("river")
}

func Close() {
	if err := cli.Close(ctx); err != nil {
		log.Logger.Info(err)
	}
}

func (obj *Algae) Insert() error {
	res, err := algae.InsertOne(ctx, obj)
	fmt.Println(res)
	return err
}
