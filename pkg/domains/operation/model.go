package operation

import "time"

type OperationType struct {
	ID          string `gorm:"type:uuid;primaryKey"`      //Operation UUID identifier
	Name        string `gorm:"type:varchar(50);not null"` //Operation type name.
	Description string `gorm:"type:varchar(255);"`        //Operation type description. Used to clarify usages and situations.
	Active      bool   `gorm:"active"`
	CreateDate  time.Time
	UpdateDate  time.Time
}
