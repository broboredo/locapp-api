package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       float32 `sql:"type:decimal(10,2);"`
	Quantity    int     `gorm:"default:1"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deteledAt,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Quantity    int       `json:"quantity"`
}
