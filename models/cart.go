package models

import "time"

// Product model struct
type Cart struct {
	ID            	int                 `json:"id" gorm:"primary_key:auto_increment"`
	QTY           	int                 `json:"qty" gorm:"type: int"`
	SubTotal      	int                 `json:"subtotal" gorm:"type: int"`
	ProductId     	int                 `json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product       	ProductTransaction  `json:"product"`
	ToppingID     	[]int               `json:"-" gorm:"-"`
	Topping       	[]Topping           `json:"topping" gorm:"many2many:cart_toppings; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionID 	int 				`json:"-" gorm:"type: int"`
	Transaction		TransactionResponse `json:"transaction"`
	Status			string				`json:"status"`
	CreatedAt     	time.Time           `json:"created_at"`
	UpdatedAt     	time.Time           `json:"updated_at"`
}

type CartResponse struct {
	ID            	int                	`json:"id"`
	UserID        	int                	`json:"user_id"`
	Total			int					`json:"total"`
	TransactionID 	int                	`json:"transaction_id"`
	ProductID     	int                	`json:"product_id"`
	ToppingID     	[]int              	`json:"topping_id"`
	Product       	ProductTransaction 	`json:"product"`
	Topping       	[]Topping          	`json:"topping"`
}

func (CartResponse) TableName() string {
	return "carts"
}
