package models

import "time"

// NewbornEEG represents the structure of the newborn_eegs table
type NewbornEEG struct {
	ID        uint       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TrialCode string     `json:"trial_code"`
	Channel   string     `json:"channel"`
	Value     float64    `json:"value"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName returns the name of the newborn_eegs table
func (*NewbornEEG) TableName() string {
	return "newborn_eegs"
}
