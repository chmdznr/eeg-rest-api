package models

import "time"

// PregnantData represents the structure of the pregnant_datas table
type PregnantData struct {
	ID        uint       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TrialCode string     `json:"trial_code"`
	Name      string     `json:"name"`
	Gravidity int        `json:"gravidity"`
	Parity    int        `json:"parity"`
	Age       int        `json:"age"`
	HR        int        `json:"hr"`
	Spo2      int        `json:"spo_2" gorm:"column:spo_2"`
	RespCount int        `json:"resp_count"`
	Sistole   int        `json:"sistole"`
	Diastole  int        `json:"diastole"`
	FetalHR   int        `json:"fetal_hr"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName returns the name of the pregnant_datas table
func (*PregnantData) TableName() string {
	return "pregnant_datas"
}
