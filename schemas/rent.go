package schemas

import "time"

type Rent struct {
	ID             uint      `gorm:"primaryKey"`
	CustomerID     uint      `gorm:"not null"`
	Customer       Customer  `gorm:"foreignKey:CustomerID"`
	StartDate      time.Time `gorm:"not null"`
	EndDate        time.Time `gorm:"not null"`
	Contract       string
	TotalAmount    float64 `gorm:"not null"`
	AmountPaid     float64 `gorm:"not null"`
	Paid           bool    `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	RentedProducts []RentedProduct `gorm:"foreignKey:RentID"`
}

type RentedProduct struct {
	ID        uint    `gorm:"primaryKey"`
	RentID    uint    `gorm:"not null"`
	Rent      Rent    `gorm:"foreignKey:RentID"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `gorm:"default:1"`
}
