package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() (client *mongo.Client, cf func(), err error) {
	var (
		cfg struct {
			Uri string
		}
		ct paladin.TOML
	)
	if err = paladin.Get("mongo.toml").Unmarshal(&ct); err != nil {
		return
	}
	client = &mongo.Client{}
	if err = ct.Get("mongo").UnmarshalTOML(&cfg); err != nil {
		return
	}
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		return
	}
	client.Database("bcos").Collection("orders").Drop(context.Background())
	cf = func() { client.Disconnect(context.Background()) }
	return
}
