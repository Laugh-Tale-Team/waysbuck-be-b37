package models

import "time"

// Product model struct
type Transaction struct {
	ID        int64               `json:"id"`
	UserId    int                 `json:"user_id" gorm:"type: int"`
	User      User				 `json:"user"`
	Status    string              `json:"status"`
	Total     int                 `json:"total" gorm:"type: int"`
	Cart      []Cart              `json:"carts"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}

type TransactionResponse struct {
	ID     int64 `json: "id"`
	UserID int   `json: "user_id"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
