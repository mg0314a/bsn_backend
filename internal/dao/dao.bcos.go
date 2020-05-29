package dao

import (
	"bsn_backend/internal/model"
	"context"
	"crypto/ecdsa"
	"github.com/chislab/go-fiscobcos/accounts/abi/bind"
	"github.com/chislab/go-fiscobcos/channel"
	"github.com/chislab/go-fiscobcos/crypto"
	"github.com/chislab/go-fiscobcos/ethclient"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"math/big"
)

var Auths map[string]*bind.TransactOpts

func NewBC() (bc *model.Bcos, cf func(), err error) {
	var (
		cfg struct {
			Endpoint string
		}
		chanCfg  channel.Config
		ctrAddrs struct {
			Product  string
			Payment  string
			Material string
		}
		ct      paladin.TOML
		priKeys = make(map[string]string)
	)
	if err = paladin.Get("bcos.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Http").UnmarshalTOML(&cfg); err != nil {
		return
	}
	if err = ct.Get("Channel").UnmarshalTOML(&chanCfg); err != nil {
		return
	}

	if err = ct.Get("CtrAddrs").UnmarshalTOML(&ctrAddrs); err != nil {
		return
	}
	if err = ct.Get("PriKeys").UnmarshalTOML(&priKeys); err != nil {
		return
	}
	bc = new(model.Bcos)
	if bc.HttpClient, err = ethclient.Dial(cfg.Endpoint); err != nil {
		return
	}
	if bc.ChanClient, err = channel.NewClient(chanCfg); err != nil {
		return
	}

	initAuths(bc.HttpClient, priKeys)
	//bc.Material, _ = material.NewMaterial(common.HexToAddress(ctrAddrs.Material), bc.HttpClient)
	//bc.Payment, _ = payment.NewPayment(common.HexToAddress(ctrAddrs.Payment), bc.HttpClient)
	//bc.Produce, _ = produce.NewProduce(common.HexToAddress(ctrAddrs.Product), bc.HttpClient)

	return
}

func NewAuthFromPriKey(height *big.Int, priKey ...string) *bind.TransactOpts {
	var priv *ecdsa.PrivateKey
	if len(priKey) == 0 {
		priv, _ = crypto.GenerateKey()
	} else {
		priv, _ = crypto.HexToECDSA(priKey[0])
	}
	auth := bind.NewKeyedTransactor(priv, 1, 1)

	auth.BlockLimit = height.Uint64() + 100
	auth.Context = context.Background()
	return auth
}

func initAuths(gethCli *ethclient.Client, priKeys map[string]string) {
	Auths = make(map[string]*bind.TransactOpts)
	height, _ := gethCli.BlockNumber(context.Background())
	for k, v := range priKeys {
		Auths[k] = NewAuthFromPriKey(height, v)
	}
}
