package models

type Pedido struct {
	ID        string `json:"id"`
	ClienteId int    `json:clienteId`
}
