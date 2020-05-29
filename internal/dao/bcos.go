package dao

import (
	pb "bsn_backend/api"
	"bsn_backend/contracts/access"
	"bsn_backend/contracts/material"
	"bsn_backend/contracts/payment"
	"bsn_backend/contracts/produce"
	"bsn_backend/internal/model"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/chislab/go-fiscobcos/accounts/abi/bind"
	"github.com/chislab/go-fiscobcos/channel"
	"github.com/chislab/go-fiscobcos/common"
	"github.com/chislab/go-fiscobcos/core/types"
	"github.com/go-kratos/kratos/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	err          error
	tx           *types.Transaction
	accessAddr   common.Address
	produceAddr  common.Address
	materialAddr common.Address
	paymentAddr  common.Address

	receipt *types.Receipt

	callOpts = &bind.CallOpts{From: common.HexToAddress("0x1"), GroupId: 1}

	orderIDCache = make(map[int64]string)
)

func (d *dao) Price(ctx context.Context, authName string, name string, isMaterial bool) *pb.RspPrice {
	var prices = &pb.RspPrice{Prices: []*pb.Price{}}
	if isMaterial {
		price, err := d.Bcos.GetPrice(callOpts, Auths[fmt.Sprintf("%s_PRODUCER", name)].From, str2Big(name))
		if err != nil {
			return nil
		}
		prices.Prices = append(prices.Prices, &pb.Price{
			Name:     name,
			Producer: Auths[fmt.Sprintf("%s_PRODUCER", name)].From.String(),
			Price:    price.Int64(),
		})
		return prices
	}
	price, err := d.Bcos.GetProductPrice(callOpts, Auths[fmt.Sprintf("%s_PRODUCER", name)].From, str2Big(name))
	if err != nil {
		return nil
	}
	prices.Prices = append(prices.Prices, &pb.Price{
		Name:     name,
		Producer: Auths[fmt.Sprintf("%s_PRODUCER", name)].From.String(),
		Price:    price.Int64(),
	})
	return prices
}

func (d *dao) SetPrice(ctx context.Context, req *pb.SetPriceRequest) *pb.RspPrice {
	authName := req.AuthName
	if req.IsMaterial {
		tx, err = d.Bcos.SetPrice(Auths[authName], str2Big(req.Type), big.NewInt(req.Price))
	} else {
		tx, err = d.Bcos.UpdateProductPrice(Auths[authName], str2Big(req.Type), big.NewInt(req.Price))
	}
	d.checkTx(tx, err)
	return d.Price(ctx, authName, req.Type, req.IsMaterial)
}

func (d *dao) InitializeContracts(c context.Context) (err error) {
	return d.initContracts()
}

func (d *dao) deloyContracts() {
	fmt.Println("deploay contracts...")
	// 方便bsn_backend对该模块进行复用
	accessAddr, tx, d.Bcos.Access, err = access.DeployAccess(Auths["RootAdmin"], d.Bcos.HttpClient)
	d.checkTx(tx, err)
	fmt.Println("accessAddr", accessAddr.String())

	produceAddr, tx, d.Bcos.Produce, err = produce.DeployProduce(Auths["PdtAdmin"], d.Bcos.HttpClient, accessAddr, big.NewInt(3))
	d.checkTx(tx, err)
	fmt.Println("produceAddr", produceAddr.String())

	materialAddr, tx, d.Bcos.Material, err = material.DeployMaterial(Auths["MtrAdmin"], d.Bcos.HttpClient, accessAddr)
	d.checkTx(tx, err)
	fmt.Println("materialAddr", materialAddr.String())

	paymentAddr, tx, d.Bcos.Payment, err = payment.DeployPayment(Auths["PayAdmin"], d.Bcos.HttpClient, big.NewInt(50))
	d.checkTx(tx, err)
	fmt.Println("paymentAddr", paymentAddr.String())

	// 设置权限
	tx, err = d.Bcos.Access.GrantPayment(Auths["RootAdmin"], paymentAddr)
	d.checkTx(tx, err)
	tx, _ = d.Bcos.Access.GrantProductProducer(Auths["RootAdmin"], Auths["TV_PRODUCER"].From)
	tx, _ = d.Bcos.Access.GrantMaterialProducer(Auths["RootAdmin"], Auths["CPU_PRODUCER"].From)
	tx, _ = d.Bcos.Access.GrantMaterialProducer(Auths["RootAdmin"], Auths["LCD_PRODUCER"].From)
	tx, _ = d.Bcos.Access.GrantMaterialProducer(Auths["RootAdmin"], Auths["AUDIO_PRODUCER"].From)

	//init 合约
	tx, err = d.Bcos.Produce.SetMaterialContract(Auths["PdtAdmin"], materialAddr)
	d.checkTx(tx, err)
	tx, err = d.Bcos.Produce.SetPaymentContract(Auths["PdtAdmin"], paymentAddr)
	d.checkTx(tx, err)
	tx, err = d.Bcos.Payment.SetProductProducer(Auths["PayAdmin"], produceAddr)
	d.checkTx(tx, err)
	tx, err = d.Bcos.Payment.SetMaterialProducer(Auths["PayAdmin"], materialAddr)
	d.checkTx(tx, err)

	// 充值
	mintAmount := big.NewInt(100000)
	tx, err = d.Bcos.Payment.Mint(Auths["PayAdmin"], Auths["CPU_PRODUCER"].From, mintAmount)
	tx, err = d.Bcos.Payment.Mint(Auths["PayAdmin"], Auths["LCD_PRODUCER"].From, mintAmount)
	tx, err = d.Bcos.Payment.Mint(Auths["PayAdmin"], Auths["AUDIO_PRODUCER"].From, mintAmount)
	tx, err = d.Bcos.Payment.Mint(Auths["PayAdmin"], Auths["Custom"].From, mintAmount)
	tx, err = d.Bcos.Payment.Mint(Auths["PayAdmin"], Auths["TV_PRODUCER"].From, mintAmount)
}

func (d *dao) initContracts() error {
	fmt.Println("Init contracts, please be patient...")
	d.deloyContracts()
	fmt.Println("================== check full procedures ==================")

	tx, err = d.Bcos.Material.SetPrice(Auths["LCD_PRODUCER"], str2Big("LCD"), big.NewInt(int64(100)))
	d.checkTx(tx, err)
	tx, err = d.Bcos.Material.SetPrice(Auths["AUDIO_PRODUCER"], str2Big("Audio"), big.NewInt(int64(50)))
	d.checkTx(tx, err)
	tx, err = d.Bcos.Material.SetPrice(Auths["CPU_PRODUCER"], str2Big("CPU"), big.NewInt(int64(200)))
	d.checkTx(tx, err)
	tx, err = d.Bcos.Produce.UpdateProductPrice(Auths["TV_PRODUCER"], str2Big("TV"), big.NewInt(3000))
	d.checkTx(tx, err)
	for k, v := range Auths {
		log.Info("%s: %s", k, v.From.String())
	}
	// 普通用户下产品订单
	//tx, err = d.Bcos.Payment.MakeOrder(Auths["Custom"], false, Auths["TV_PRODUCER"].From, str2Big("TV"), big.NewInt(5), big.NewInt(3000))
	//d.checkTx(tx, err)
	//// 产品厂家下原材料订单
	//tx, err = d.Bcos.Payment.MakeOrder(Auths["TV_PRODUCER"], true, Auths["CPU_PRODUCER"].From, str2Big("LCD"), big.NewInt(100), big.NewInt(100))
	//d.checkTx(tx, err)
	//tx, err = d.Bcos.Payment.MakeOrder(Auths["TV_PRODUCER"], true, Auths["LCD_PRODUCER"].From, str2Big("Audio"), big.NewInt(100), big.NewInt(50))
	//d.checkTx(tx, err)
	//tx, err = d.Bcos.Payment.MakeOrder(Auths["TV_PRODUCER"], true, Auths["AUDIO_PRODUCER"].From, str2Big("CPU"), big.NewInt(100), big.NewInt(200))
	//d.checkTx(tx, err)
	//
	//// 检查原料, 如果数量充足, 就进行订单确认，但是这里假定订单数量都不充足。进行生产后交付。理应检查原件
	//// 线下准备原料, 在这里进行注册
	//tx, err = d.Bcos.Material.NewMaterial(Auths["CPU_PRODUCER"], str2Big("LCD"), big.NewInt(300), str2Big("LCD_1"))
	//d.checkTx(tx, err)
	//tx, err = d.Bcos.Material.NewMaterial(Auths["LCD_PRODUCER"], str2Big("Audio"), big.NewInt(300), str2Big("Audio_1"))
	//d.checkTx(tx, err)
	//tx, err = d.Bcos.Material.NewMaterial(Auths["AUDIO_PRODUCER"], str2Big("CPU"), big.NewInt(300), str2Big("CPU_1"))
	//d.checkTx(tx, err)

	// // 生产厂家确认原料订单
	// tx, err = d.Bcos.Payment.ConfirmOrder(Auths["TV_PRODUCER"], big.NewInt(0)) // LCD=
	// d.checkTx(tx, err)
	// tx, err = d.Bcos.Payment.ConfirmOrder(Auths["TV_PRODUCER"], big.NewInt(2)) // Audio
	// d.checkTx(tx, err)
	// tx, err = d.Bcos.Payment.ConfirmOrder(Auths["TV_PRODUCER"], big.NewInt(4)) // CPU
	// d.checkTx(tx, err)

	// // 进行产品生产
	// // 1. 首先消耗零件
	// tx, err = d.Bcos.Material.ConsumeMaterial(Auths["TV_PRODUCER"], str2Big("LCD"), big.NewInt(20))
	// d.checkTx(tx, err)
	// tx, err = d.Bcos.Material.ConsumeMaterial(Auths["TV_PRODUCER"], str2Big("Audio"), big.NewInt(20))
	// d.checkTx(tx, err)
	// tx, err = d.Bcos.Material.ConsumeMaterial(Auths["TV_PRODUCER"], str2Big("CPU"), big.NewInt(20))
	// d.checkTx(tx, err)

	// // 2. 产品进行注册, 注册不需要具体知道零件消耗了多少
	// for i := 0; i < 10; i++ {
	// 	productID := fmt.Sprintf("ProductID_%d", i)
	// 	tx, err = d.Bcos.Produce.RegisterProduct(Auths["TV_PRODUCER"], str2Big("TV"), str2Big(productID), str2Big("2020-05-20"),
	// 		[]*big.Int{str2Big("LCD_1"), str2Big("Audio_1"), str2Big("CPU_1")})
	// 	d.checkTx(tx, err)
	// }

	// tx, err = d.Bcos.Payment.ConfirmOrder(Auths["Custom"], big.NewInt(1))
	// d.checkTx(tx, err)

	return nil
}

func (d *dao) parseProductEvt(msg *types.Log) {
	switch msg.Topics[0] {
	case d.Bcos.Produce.ABI.Events["EvtProductCreated"].ID:
		evt, _ := d.Bcos.Produce.ParseEvtProductCreated(*msg)
		log.Info("produce ProductCreated %s, id: %s, creator: %s", string(evt.ProductType.Bytes()), string(evt.ProductID.Bytes()), evt.Creator.String())
	case d.Bcos.Produce.ABI.Events["EvtProductOwnerChanged"].ID:
		evt, _ := d.Bcos.Produce.ParseEvtProductOwnerChanged(*msg)
		log.Info("produce ProductOwnerChanged, id: %s, new: %s, old: %s", string(evt.Id.Bytes()), evt.NewOwner.String(), evt.OldOwner.String())
	}
}

func (d *dao) parseMaterialEvt(msg *types.Log) {
	switch msg.Topics[0] {
	case d.Bcos.Material.ABI.Events["EvtMaterialCreated"].ID:
		evt, _ := d.Bcos.Material.ParseEvtMaterialCreated(*msg)
		log.Info("material EvtMaterialCreated from: %s, materialType: %s, amount: %s", evt.From.String(), string(evt.MaterialType.Bytes()), evt.Num.String())
	case d.Bcos.Material.ABI.Events["EvtMaterialTransferred"].ID:
		evt, _ := d.Bcos.Material.ParseEvtMaterialTransferred(*msg)
		log.Info("material EvtMaterialTransferred from: %s, to: %s, materialType: %s, amount: %s",
			evt.From.String(), evt.To.String(), string(evt.MaterialType.Bytes()), evt.Num.String())
	case d.Bcos.Material.ABI.Events["EvtMaterialConsumed"].ID:
		evt, _ := d.Bcos.Material.ParseEvtMaterialConsumed(*msg)
		log.Info("material EvtMaterialConsumed from: %s, materialType: %s, amount: %s", evt.From.String(), string(evt.MaterialType.Bytes()), evt.Num.String())
	case d.Bcos.Material.ABI.Events["EvtPriceUpdated"].ID:
		evt, _ := d.Bcos.Material.ParseEvtPriceUpdated(*msg)
		log.Info("material EvtPriceUpdated from: %s, materialType: %s, Price: %s", evt.From.String(), string(evt.MaterialType.Bytes()), evt.Price.String())
	}
}

func (d *dao) parsePaymentEvt(msg *types.Log) {
	switch msg.Topics[0] {
	case d.Bcos.Payment.ABI.Events["EvtMakeOrder"].ID:
		evt, _ := d.Bcos.Payment.ParseEvtMakeOrder(*msg)
		mgoTx := d.mgodb.Database("bcos").Collection("orders")
		orderLog := model.Order{
			OId:       evt.Id.Int64(),
			Payer:     evt.Payer.String(),
			Producer:  evt.Producer.String(),
			Amount:    evt.Cnt.Int64(),
			OrderType: string(evt.OrderType.Bytes()),
			CreatedAt: time.Now().Unix(), // TODO 应该设置为交易时间
			Status:    0,
		}
		_, err = mgoTx.InsertOne(context.TODO(), orderLog)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Info("payment MakeOrder %s, id: %s, payer: %s, producer: %s, amount: %s, price: %s",
			string(evt.OrderType.Bytes()), evt.Id.Text(16), evt.Payer.String(), evt.Producer.String(), evt.Cnt.String(), evt.Price.String())
	case d.Bcos.Payment.ABI.Events["EvtConfirmOrder"].ID:
		evt, _ := d.Bcos.Payment.ParseEvtConfirmOrder(*msg)
		log.Info("payment ConfirmOrder %s, %s, id: %s", evt.From.String(), string(evt.OrderType.Bytes()), evt.Id.Text(16))
	case d.Bcos.Payment.ABI.Events["EvtCancelOrder"].ID:
		evt, _ := d.Bcos.Payment.ParseEvtCancelOrder(*msg)
		log.Info("payment CancelOrder, id: %s", evt.Id.Text(16))
	case d.Bcos.Payment.ABI.Events["EvtMint"].ID:
		evt, _ := d.Bcos.Payment.ParseEvtMint(*msg)
		log.Info("payment EvtMint account %s, amount: %s", evt.Account.String(), evt.Amount.String())
	}
}

func (d *dao) checkTx(tx *types.Transaction, err error) *types.Receipt {
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	receipt, err = func(tx *types.Transaction, err error) (*types.Receipt, error) {
		receipt := d.WaitMinedByHash(tx.Hash())
		if receipt.Status != "0x0" {
			return receipt, fmt.Errorf("receipt.Status = %s TxHash = %s, Output: %s", receipt.Status, receipt.TxHash.String(), getReceiptOutput(receipt.Output))
		}
		return receipt, nil
	}(tx, err)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return receipt
}

func getReceiptOutput(output string) string {
	if strings.HasPrefix(output, "0x") {
		output = output[2:]
	}
	b, err := hex.DecodeString(output)
	if err != nil || len(b) < 36 {
		return output
	}
	b = b[36:]
	tail := len(b) - 1
	for ; tail >= 0; tail-- {
		if b[tail] != 0 {
			break
		}
	}
	return string(b[:tail+1])
}

func (d *dao) WatchEvent() {
	chanArg := channel.RegisterEventLogRequest{
		FromBlock: "latest",
		ToBlock:   "latest",
		Addresses: []string{},
		Topics:    []string{},
		GroupID:   "1",
		FilterID:  channel.NewFilterID(),
	}
	ch, err := d.Bcos.ChanClient.SubEventLogs(chanArg)
	if err != nil {
		panic(err)
	}
	for {
		msg := <-ch
		switch msg.Address {
		case produceAddr:
			d.parseProductEvt(msg)
		case materialAddr:
			d.parseMaterialEvt(msg)
		case paymentAddr:
			d.parsePaymentEvt(msg)
		}
	}
}

func str2Big(str string) *big.Int {
	return new(big.Int).SetBytes([]byte(str))
}

func (d *dao) WaitMinedByHash(txHash common.Hash) *types.Receipt {
	ctx := context.Background()
	queryTicker := time.NewTicker(time.Millisecond * 200)
	defer queryTicker.Stop()
	for {
		receipt, _ := d.Bcos.HttpClient.TransactionReceipt(ctx, txHash)
		if receipt != nil {
			return receipt
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil
		case <-queryTicker.C:
		}
	}
}

func (d *dao) RegisterMaterial(ctx context.Context, authName string, materialType string, total int64, batchID string) *pb.RegisterMaterialResponse {
	tx, err = d.Bcos.Material.NewMaterial(Auths[authName], str2Big(materialType), big.NewInt(total), str2Big(batchID))
	receipt = d.checkTx(tx, err)
	m, err := d.Bcos.ReadNewMaterial(receipt.Output)
	if err != nil {
		var response = &pb.RegisterMaterialResponse{
			MaterialType: "",
			BatchID:      "",
			Producer:     "",
			CreatedAt:    "",
		}
		return response
	}
	mt := string(m.MaterialType.Bytes())
	bid := string(m.BatchID.Bytes())
	producer := m.Producer.String()
	createdAt := time.Unix(m.CreatedAt.Int64()/1000, 0).Format("2006-01-02 15:04:05")

	var response = &pb.RegisterMaterialResponse{
		MaterialType: mt,
		BatchID:      bid,
		Producer:     producer,
		CreatedAt:    createdAt,
	}
	return response
}

func (d *dao) GetMaterial(ctx context.Context, authName string, materialType string) *pb.GetMaterialResponse {
	opts := &bind.CallOpts{From: Auths[authName].From, GroupId: 1}
	if opts.From == common.HexToAddress("0x0") {
		return nil
	}
	num, err := d.Bcos.Material.GetMyMaterial(opts, str2Big(materialType))
	if err != nil {
		return nil
	}

	var response = &pb.GetMaterialResponse{
		Number: num.Int64(),
	}
	return response
}

func (d *dao) ConsumeMaterial(ctx context.Context, authName string, materialType string, count int64) *pb.ConsumeMaterialResponse {
	tx, err = d.Bcos.Material.ConsumeMaterial(Auths[authName], str2Big(materialType), big.NewInt(count))
	receipt = d.checkTx(tx, err)
	var response = &pb.ConsumeMaterialResponse{
		RcptStatus: "",
		RcptOutput: "",
	}
	return response
}

func (d *dao) GetProducts(ctx context.Context, authName string, pdtType string) (*pb.RspProducts, error) {
	addr := Auths[authName].From
	opts := &bind.CallOpts{From: addr, GroupId: 1}
	ids, err := d.Bcos.GetMyProducts(opts, str2Big(pdtType))
	if err != nil {
		return nil, err
	}
	var myPdts = &pb.RspProducts{Products: make(map[string]*pb.ProductResponse)}
	for _, id := range ids {
		pdt, err := d.Bcos.Details(opts, id)
		if err != nil {
			continue
		}
		var mtrInfos []string
		for _, v := range pdt.MaterialBatches {
			mtrInfos = append(mtrInfos, string(v.Bytes()))
		}
		myPdts.Products[string(id.Bytes())] = &pb.ProductResponse{
			Owner:           addr.String(),
			Producer:        pdt.Producer.String(),
			CreatedAt:       pdt.CreatedAt.Int64(),
			Batch:           string(pdt.Batch.Bytes()),
			MaterialBatches: mtrInfos,
			IsSold:          pdt.Sold,
		}
	}
	return myPdts, nil
}

func (d *dao) ProductDetails(ctx context.Context, req *pb.ProductDetailsRequest) (*pb.ProductResponse, error) {
	//opts := &bind.CallOpts{From: Auths[req.AuthName].From, GroupId: 1}
	product, err := d.Bcos.Details(callOpts, str2Big(req.ProductID))
	if err != nil {
		return nil, err
	}
	mbatches := make([]string, len(product.MaterialBatches))
	for i, mb := range product.MaterialBatches {
		mbatches[i] = string(mb.Bytes())
	}
	return &pb.ProductResponse{
		Owner:           product.Owner.String(),
		Producer:        product.Producer.String(),
		CreatedAt:       product.CreatedAt.Int64(),
		Batch:           string(product.Batch.Bytes()),
		MaterialBatches: mbatches,
		IsSold:          product.Sold,
	}, nil
}

func (d *dao) RegisterProduct(ctx context.Context, req *pb.RegisterProductRequest) (*pb.RespResult, error) {
	mbs := make([]*big.Int, len(req.MaterialBatches))
	for i, b := range req.MaterialBatches {
		mbs[i] = str2Big(b)
	}
	tx, err := d.Bcos.RegisterProduct(Auths[req.AuthName], str2Big(req.Type), str2Big(req.ProductID), str2Big(req.Batch), mbs)
	r := d.checkTx(tx, err)
	if r == nil {
		return &pb.RespResult{
			IsSucceed: false,
		}, nil
	}
	return &pb.RespResult{
		IsSucceed: true,
	}, nil
}

func (d *dao) MakeOrder(ctx context.Context, req *pb.ReqMakeOrder) (*pb.RespMakeOrder, error) {
	opts := new(bind.TransactOpts)
	var to common.Address
	price := new(big.Int)
	var err error
	var orderpayer string
	if req.IsMaterial {
		opts = Auths[req.Payer]
		to = Auths[req.Producer].From
		orderpayer = req.Payer
		price, err = d.Bcos.GetPrice(callOpts, to, str2Big(req.Type))
	} else {
		opts = Auths[req.Payer]
		to = Auths[req.Producer].From
		orderpayer = req.Producer
		price, err = d.Bcos.GetProductPrice(callOpts, to, str2Big(req.Type))
	}
	if err != nil {
		return nil, err
	}
	tx, err := d.Bcos.MakeOrder(opts, req.IsMaterial, to, str2Big(req.Type), big.NewInt(req.Count), price)
	receipt := d.checkTx(tx, err)
	if receipt == nil {
		return nil, fmt.Errorf("failed to make order")
	}
	id, err := d.Bcos.ReadMakeOrder(receipt.Output)
	if err != nil {
		return nil, err
	}
	orderIDCache[id.Int64()] = orderpayer
	return &pb.RespMakeOrder{
		OrderId: id.Text(10),
	}, nil
}

func (d *dao) ConfirmOrder(ctx context.Context, req *pb.ConfirmOrderRequest) (*pb.RespResult, error) {
	//authName := req.AuthName
	// payer, ok := orderIDCache[req.OrderId]
	// if !ok {
	// 	return nil, fmt.Errorf("order not found")
	// }
	opts := Auths[req.AuthName]
	tx, err := d.Bcos.ConfirmOrder(opts, big.NewInt(req.OrderId))
	r := d.checkTx(tx, err)
	if r == nil {
		return nil, fmt.Errorf("failed to confirm order")
	}
	go d.UpdateOrder(req.OrderId, model.ORDER_DONE)
	return &pb.RespResult{
		IsSucceed: true,
	}, nil
}

func (d *dao) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.RespResult, error) {
	opts, ok := Auths[req.AuthName]
	if !ok {
		return nil, fmt.Errorf("canceler not found")
	}
	tx, err := d.Bcos.CancelOrder(opts, big.NewInt(req.OrderId))
	r := d.checkTx(tx, err)
	if r == nil {
		return nil, fmt.Errorf("failed to cancel order")
	}
	go d.UpdateOrder(req.OrderId, model.ORDER_CANCERED)

	return &pb.RespResult{
		IsSucceed: true,
	}, nil
}

func (d *dao) Mint(ctx context.Context, req *pb.MintRequest) (*pb.RespResult, error) {
	//authName := req.AuthName
	opts := Auths["PayAdmin"]
	to, ok := Auths[req.Account]
	if !ok {
		return nil, fmt.Errorf("account not found")
	}
	tx, err := d.Bcos.Mint(opts, to.From, big.NewInt(req.Amount))
	r := d.checkTx(tx, err)
	if r == nil {
		return nil, fmt.Errorf("failed to mint")
	}
	return &pb.RespResult{
		IsSucceed: true,
	}, nil
}

func (d *dao) BalanceOf(ctx context.Context, req *pb.ReqAccount) (*pb.RespBalance, error) {
	auth, ok := Auths[req.Account]
	if !ok {
		return nil, fmt.Errorf("account not found, got '%s'", req.Account)
	}

	balance, err := d.Bcos.BalanceOf(callOpts, auth.From)
	if err != nil {
		return nil, err
	}
	return &pb.RespBalance{
		Account: req.Account,
		Balance: balance.Int64(),
	}, nil
}

func (d *dao) GetOrders(from string, orderId int64) []*model.Order {
	var myOrders []*model.Order
	mgoTx := d.mgodb.Database("bcos").Collection("orders")
	auth, ok := Auths[from]
	if !ok || auth == nil {
		return []*model.Order{}
	}
	filter := bson.M{"$and": []bson.M{{"$or": []bson.M{{"payer": auth.From.String()}, {"producer": auth.From.String()}}}, {"status": orderId}}}

	cur, err := mgoTx.Find(context.TODO(), filter)
	if err != nil {
		return nil
	}
	for cur.Next(context.TODO()) {
		var result = &model.Order{}
		cur.Decode(result)
		myOrders = append(myOrders, result)
	}
	return myOrders
}

func (d *dao) UpdateOrder(oid int64, next int8) error {
	mgoTx := d.mgodb.Database("bcos").Collection("orders")
	//filter := bson.M{"$and": []bson.M{{"oid": oid}, {"payer": Auths[from].From.String()}}}
	filter := bson.M{"oid": oid}
	update := bson.D{
		{"$set", bson.D{
			{"status", next},
		}},
	}
	result, err := mgoTx.UpdateMany(context.Background(), filter, update)
	if err != nil {
		log.Error("mongo update error %v", err)
		return err
	}
	if result.ModifiedCount < 1 {
		return errors.New("order " + string(oid) + " not affected.")
	}
	return nil
}

func (d *dao) GetOrder(ctx context.Context, req *pb.ReqOrderID) (*pb.RespOrder, error) {
	order, err := d.Bcos.GetOrder(callOpts, big.NewInt(req.OrderId))
	if err != nil {
		return nil, err
	}
	return &pb.RespOrder{
		Payer:    order.Payer.String(),
		Producer: order.Producer.String(),
		Amount:   order.Amount.Int64(),
		Count:    order.Count.Int64(),
		Type:     string(order.OrderType.Bytes()),
		Status:   int32(order.Status),
	}, nil
}

func (d *dao) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	ids, err := d.Bcos.Trace(callOpts, str2Big(req.MaterialBatch))
	if err != nil {
		return nil, err
	}
	pids := make([]string, len(ids))
	for i, id := range ids {
		pids[i] = string(id.Bytes())
	}
	return &pb.TraceResponse{
		ProductIds: pids,
	}, nil
}
