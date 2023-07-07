package models

type Stock struct{
	StockId int64 `json:"stockId"`
	Name string `json:"name"`
	Price int64 `json:"price"`
	Company string `json:"company"`
}

