package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Phone     string
	Address   string
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
