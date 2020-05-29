package model

import (
	"bsn_backend/contracts/access"
	"bsn_backend/contracts/material"
	"bsn_backend/contracts/payment"
	"bsn_backend/contracts/produce"
	"github.com/chislab/go-fiscobcos/channel"
	"github.com/chislab/go-fiscobcos/ethclient"
)

type Bcos struct {
	HttpClient *ethclient.Client
	ChanClient *channel.Client

	*material.Material
	*payment.Payment
	*produce.Produce
	*access.Access
}

const (
	ORDER_IN_PORGRESS = iota
	ORDER_DONE
	ORDER_CANCERED
)
