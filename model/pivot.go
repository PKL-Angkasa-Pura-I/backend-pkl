package model

import "time"

type Pivot_division_field struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string      `gorm:"unique;not null" json:"name"`
	CreatedAt     time.Time   `json:"created_at"`
	DivisionID    uint        `json:"division_id"`
	Division      Division    `json:"division" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Study_fieldID uint        `json:"study_field_id"`
	Study_field   Study_field `json:"study_field" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
