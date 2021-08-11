package transaction

import "time"

type Transaction struct {
	Id              string  `gorm:"type:uuid;primaryKey"`
	AccountID       string  `gorm:"type:uuid;not null;column:account_id"`
	OperationTypeID int     `gorm:"type:int;not null;column:operation_type_id"`
	Amount          float64 `gorm:"type:decimal(12,2)"`
	Balance         float64 `gorm:"type:decimal(12,2)"`
	Paid            bool
	EventDate       time.Time
	Approved        bool
	RejectionMotive string `json:"rejection_motive"`
}
