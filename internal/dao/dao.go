package dao

import (
	pb "bsn_backend/api"
	"bsn_backend/internal/model"
	"context"
	"time"

	"github.com/go-kratos/kratos/pkg/cache/memcache"
	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC, NewBC, NewMongo)

// Dao dao interface
// go:generate kratos tool genbts
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	//bts: -nullcache=&model.Product{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	InitializeContracts(c context.Context) (err error)
	Price(ctx context.Context, authName string, name string, isMaterial bool) *pb.RspPrice
	SetPrice(ctx context.Context, req *pb.SetPriceRequest) *pb.RspPrice
	RegisterMaterial(ctx context.Context, authName string, materialType string, total int64, batchID string) *pb.RegisterMaterialResponse
	GetMaterial(ctx context.Context, authName string, materialType string) *pb.GetMaterialResponse
	ConsumeMaterial(ctx context.Context, authName string, materialType string, count int64) *pb.ConsumeMaterialResponse
	GetProducts(ctx context.Context, authName string, pdtType string) (*pb.RspProducts, error)
	RegisterProduct(ctx context.Context, req *pb.RegisterProductRequest) (*pb.RespResult, error)

	// Payment
	ProductDetails(ctx context.Context, req *pb.ProductDetailsRequest) (*pb.ProductResponse, error)
	MakeOrder(ctx context.Context, req *pb.ReqMakeOrder) (*pb.RespMakeOrder, error)
	ConfirmOrder(ctx context.Context, req *pb.ConfirmOrderRequest) (*pb.RespResult, error)
	CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.RespResult, error)
	Mint(ctx context.Context, req *pb.MintRequest) (*pb.RespResult, error)
	BalanceOf(ctx context.Context, req *pb.ReqAccount) (*pb.RespBalance, error)

	// Order
	GetOrders(from string, orderId int64) []*model.Order
	GetOrder(ctx context.Context, req *pb.ReqOrderID) (*pb.RespOrder, error)

	Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error)
}

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	mgodb      *mongo.Client
	demoExpire int32
	Bcos       *model.Bcos // blockchain client
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, db *sql.DB, bc *model.Bcos, mgodb *mongo.Client) (d Dao, cf func(), err error) {
	return newDao(r, mc, db, bc, mgodb)
}

func newDao(r *redis.Redis, mc *memcache.Memcache, db *sql.DB, bc *model.Bcos, mgodb *mongo.Client) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		Bcos:       bc,
		db:         db,
		redis:      r,
		mc:         mc,
		mgodb:      mgodb,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	go d.WatchEvent()
	go d.initContracts()
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
