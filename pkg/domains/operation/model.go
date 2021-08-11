package operation

import "time"

type OperationType struct {
	ID                            int    `gorm:"type:int;primaryKey"`
	Name                          string `gorm:"type:varchar(50);not null"`
	Description                   string `gorm:"type:varchar(255);"`
	MultiplicationFactor          int    `gorm:"type:int"`
	ShouldUseMultiplicationFactor bool
	Active                        bool
	CreateDate                    time.Time
	UpdateDate                    time.Time
}

func (t OperationType) IsCreditTransaction() bool {
	return !t.ShouldUseMultiplicationFactor && t.MultiplicationFactor >= 0
}
