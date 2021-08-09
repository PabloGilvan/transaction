package transaction

import "time"

type Transaction struct {
	Id              string  `gorm:"type:uuid;primaryKey"`
	AccountID       string  `gorm:"type:uuid;not null;column:account_id"`
	OperationTypeID string  `gorm:"type:uuid;not null;column:operation_type_id"`
	Amount          float64 `gorm:"type:numeric(12,2)"`
	EventDate       time.Time
	Approved        bool
	RejectionMotive string `json:"rejection_motive"` //Rejection can't be in a table, because you would to know the rules to reject, if we need to know there is no how to set it dynamically
}
