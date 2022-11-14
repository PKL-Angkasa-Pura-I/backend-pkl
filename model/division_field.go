package model

import "time"

type Pivot_division_field struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt     time.Time   `json:"created_at"`
	DivisionID    uint        `json:"division_id"`
	Division      Division    `json:"division" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Study_fieldID uint        `json:"study_field_id"`
	Study_field   Study_field `json:"study_field" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type List_division_field struct {
	Division_Name string `json:"division_name"`
	Field_Name    string `json:"field_name"`
}
