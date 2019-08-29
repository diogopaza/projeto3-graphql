package models

type Item struct {
	ID         string   `json:"id"`
	Quantidade int      `json:"quantidade"`
	Valor      float64  `json:"valor"`
	Produto    *Produto `json:"produto"`
}
