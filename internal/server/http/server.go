package http

import (
	"net/http"

	pb "bsn_backend/api"
	"bsn_backend/internal/model"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svc pb.DemoServer

// New new a bm server.
func New(s pb.DemoServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	//engine = bm.DefaultServer(&cfg)
	engine = bm.NewServer(&cfg)
	engine.Use(bm.Recovery())
	cors := bm.CORS([]string{"127.0.0.1:5000", "localhost:5000", "127.0.0.1:8080", "localhost:8080", "192.168.31.227:5000"})
	engine.Use(cors)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	e.GET("/init", HandleInit)
	g := e.Group("/backend")
	{
		g.GET("/start", howToStart)
	}

}
