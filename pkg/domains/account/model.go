package account

import (
	"time"
)

type Account struct {
	ID                   string  `gorm:"type:uuid;primaryKey;"`
	Number               string  `gorm:"type:varchar(16);not null"`
	DocumentNumber       string  `gorm:"type:varchar(13);not null"`
	AvailableCreditLimit float64 `gorm:"type:decimal(12,2);default:0"`
	Active               bool    `gorm:"active"`
	CreateDate           time.Time
	UpdateDate           time.Time
}
