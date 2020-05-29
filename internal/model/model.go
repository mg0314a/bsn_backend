package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

//type Article struct {
//	ID int64
//	Content string
//	Author string
//}

type Product struct {
	ID              int       `json:"id"`
	Owner           string    `json:"owner"`
	BatchID         string    `json:"batch_id"`
	Materials       *[]string `json:"materials"`
	ProductionDate  string    `json:"production_date"`
	DeprecatedState bool      `json:"deprecated_state"`
	IsSold          bool      `json:"is_sold"`
}
