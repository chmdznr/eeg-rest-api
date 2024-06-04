package models

import "time"

// NewbornCv represents the structure of the newborn_cvs table
type NewbornCv struct {
	ID        uint       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TrialCode string     `json:"trial_code"`
	DataType  string     `json:"data_type"`
	Notes     string     `json:"notes"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName returns the name of the newborn_cvs table
func (*NewbornCv) TableName() string {
	return "newborn_cvs"
}
