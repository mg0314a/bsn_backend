package http

import bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

func HandleInit(c *bm.Context) {
	svc.InitContracts(c, nil)
	c.JSON(nil, nil)
}
