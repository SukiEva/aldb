package model

import (
	"context"
	"github.com/SukiEva/aldb/server/config"
	"github.com/SukiEva/aldb/server/utils/log"
	"github.com/qiniu/qmgo"
)

var (
	cli *qmgo.Client
	mdb *qmgo.Database
	ctx context.Context
	cfg = config.GetEnv()

	algae      *qmgo.Collection
	river      *qmgo.Collection
	operator   *qmgo.Collection
	operation  *qmgo.Collection
	annotation *qmgo.Collection
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
	operator = mdb.Collection("operator")
	operation = mdb.Collection("operation")
	annotation = mdb.Collection("annotation")
}

type Mgo struct{}

func Close() {
	if err := cli.Close(ctx); err != nil {
		log.Logger.Info(err)
	}
}
