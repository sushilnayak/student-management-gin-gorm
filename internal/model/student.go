package model

import (
	"time"
)

type Student struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"size:100;not null" json:"firstName"`
	LastName  string    `gorm:"size:100;not null" json:"lastName"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Age       int       `gorm:"not null" json:"age"`
	Grade     string    `gorm:"size:2;not null" json:"grade"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
