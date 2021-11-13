package model

import "time"

type Response struct {
	Data []Data `json:"Data"`
}
type Data struct {
	Date  time.Time `json:"Date"`
	Value []Funds   `json:"Value"`
}
type Funds struct {
	Name        string    `json:"Name"`
	RankOfFund  float32   `json:"Rank of fund"`
	UpdatedDate time.Time `json:"Updated date"`
	Performance float32   `json:"Performance"`
	Price       float32   `json:"Price"`
}
type FinomenaResponse struct {
	Status bool            `json:"status"`
	Error  string          `json:"error"`
	Data   []FinnomenaData `json:"data"`
}
type FinnomenaData struct {
	MstarID          string    `json:"mstar_id"`
	ThailandFundCode string    `json:"thailand_fund_code"`
	NavReturn        float32   `json:"nav_return"`
	Nav              float32   `json:"nav"`
	NavDate          time.Time `json:"nav_date"`
	AvgReturn        float32   `json:"avg_return"`
}
