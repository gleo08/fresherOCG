package models

type Order struct {
	Id         int         `json:"id"`
	UserId     int         `json:"user_id"`
	OrderItems []OrderItem `json:"order_items"`
}

type OrderItem struct {
	Id        int    `json:"id"`
	OrderId   int    `json:"order_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (order *Order) TableName() string {
	return "order"
}

func (orderItem *OrderItem) TableName() string {
	return "order_item"
}
