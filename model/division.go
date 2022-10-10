package model

import "time"

type Division struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Quota     *uint     `gorm:"default:0;not null" json:"quota"`
	CreatedAt time.Time `json:"created_at"`
}
