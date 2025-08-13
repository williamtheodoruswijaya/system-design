package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int            `gorm:"primaryKey;autoIncrement"`
	Username   string         `gorm:"type:varchar(100);unique;not null" json:"username"`
	Fullname   string         `gorm:"type:varchar(100);not null" json:"fullname"`
	Email      string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password   string         `gorm:"type:varchar(255);not null" json:"-"`
	ProfileUrl sql.NullString `gorm:"type:varchar(255)" json:"profileurl"`
	CreatedAt  time.Time      `json:"createdat"`
	// Posts      []Post         `gorm:"foreignKey:UserID"` // contoh kalau mau ada foreign key
}
