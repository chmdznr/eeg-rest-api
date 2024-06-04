package models

import "time"

// NewbornData represents the structure of the newborn_datas table
type NewbornData struct {
	ID         uint       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TrialCode  string     `json:"trial_code"`
	Name       string     `json:"name"`
	MotherName string     `json:"mother_name"`
	MotherAge  int        `json:"mother_age"`
	Gravidity  int        `json:"gravidity"`
	Parity     int        `json:"parity"`
	AccelX     float64    `json:"accel_x"`
	AccelY     float64    `json:"accel_y"`
	AccelZ     float64    `json:"accel_z"`
	Thermal    float64    `json:"thermal"`
	Spo2       int        `json:"spo_2"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

// TableName returns the name of the newborn_datas table
func (*NewbornData) TableName() string {
	return "newborn_datas"
}
