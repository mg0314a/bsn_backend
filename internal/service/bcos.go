package service

import (
	pb "bsn_backend/api"
	"bsn_backend/internal/model"
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Service) InitContracts(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return nil, s.dao.InitializeContracts(ctx)
}

func (s *Service) GetPrice(ctx context.Context, req *pb.GetPriceRequest) (*pb.RspPrice, error) {
	authName := req.AuthName
	name := req.Type
	isMaterial := req.IsMaterial
	return s.dao.Price(ctx, authName, name, isMaterial), nil
}

func (s *Service) SetPrice(ctx context.Context, req *pb.SetPriceRequest) (*pb.RspPrice, error) {
	return s.dao.SetPrice(ctx, req), nil
}

func (s *Service) RegisterMaterial(ctx context.Context, req *pb.RegisterMaterialRequest) (*pb.RegisterMaterialResponse, error) {
	materialType := req.MaterialType
	batchID := req.BatchID
	total := req.TotalNum
	authName := req.AuthName
	return s.dao.RegisterMaterial(ctx, authName, materialType, total, batchID), nil
}

func (s *Service) GetMaterial(ctx context.Context, req *pb.GetMaterialRequest) (*pb.GetMaterialResponse, error) {
	materialType := req.MaterialType
	authName := req.AuthName
	return s.dao.GetMaterial(ctx, authName, materialType), nil
}

func (s *Service) ConsumeMaterial(ctx context.Context, req *pb.ConsumeMaterialRequest) (*pb.ConsumeMaterialResponse, error) {
	materialType := req.MaterialType
	count := req.Count
	authName := req.AuthName
	return s.dao.ConsumeMaterial(ctx, authName, materialType, count), nil
}

func (s *Service) GetProducts(ctx context.Context, req *pb.GetProductRequest) (*pb.RspProducts, error) {
	authName := req.AuthName
	productType := ctx.(*bm.Context).Request.Form.Get("product_type")
	return s.dao.GetProducts(ctx, authName, productType)
}

func (s *Service) RegisterProduct(ctx context.Context, req *pb.RegisterProductRequest) (*pb.RespResult, error) {
	return s.dao.RegisterProduct(ctx, req)
}

func (s *Service) ProductDetails(ctx context.Context, req *pb.ProductDetailsRequest) (*pb.ProductResponse, error) {
	return s.dao.ProductDetails(ctx, req)
}

func (s *Service) MakeOrder(ctx context.Context, req *pb.ReqMakeOrder) (*pb.RespMakeOrder, error) {
	return s.dao.MakeOrder(ctx, req)
}

func (s *Service) ConfirmOrder(ctx context.Context, req *pb.ConfirmOrderRequest) (*pb.RespResult, error) {
	return s.dao.ConfirmOrder(ctx, req)
}

func (s *Service) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.RespResult, error) {
	return s.dao.CancelOrder(ctx, req)
}

func (s *Service) Mint(ctx context.Context, req *pb.MintRequest) (*pb.RespResult, error) {
	return s.dao.Mint(ctx, req)
}

func (s *Service) BalanceOf(ctx context.Context, req *pb.ReqAccount) (*pb.RespBalance, error) {
	return s.dao.BalanceOf(ctx, req)
}

func (s *Service) GetOrders(ctx context.Context, req *pb.ReqOrders) (*pb.RespOrders, error) {
	var ret = pb.RespOrders{}
	myOrders := s.dao.GetOrders(req.From, model.ORDER_IN_PORGRESS)
	for _, v := range myOrders {
		var order = pb.Order{
			OrderId:   v.OId,
			Payer:     v.Payer,
			Producer:  v.Producer,
			Amount:    v.Amount,
			OrderType: v.OrderType,
			CreatedAt: v.CreatedAt,
			Status:    v.Status,
		}
		ret.Orders = append(ret.Orders, &order)
	}
	return &ret, nil
}

func (s *Service) GetOrder(ctx context.Context, req *pb.ReqOrderID) (*pb.RespOrder, error) {
	return s.dao.GetOrder(ctx, req)
}

func (s *Service) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	return s.dao.Trace(ctx, req)
}
