package model

type Order struct {
	OId       int64
	Payer     string
	Producer  string
	Amount    int64
	OrderType string
	CreatedAt int64
	Status    int32
}
