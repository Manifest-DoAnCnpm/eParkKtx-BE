package entities

import (
	"time"
)

// `key:value` : tag cho thư viện đọc và mappping đúng giá trị
type User struct {

	// gorm.Model: nhúng các trường id, creatAt,.... vào struct
	// gorm.Model

	// field
	UserID      string    `gorm:"primaryKey;type:varchar(20);column:UserID" json:"user_id"`
	Password    string    `gorm:"type:varchar(100);not null;column:Password" json:"-"`
	Name        string    `gorm:"type:varchar(100);not null;column:Name" json:"name"`
	DoB         time.Time `gorm:"column:DoB" json:"dob"`
	Gender      string    `gorm:"type:varchar(10);column:Gender" json:"gender"`
	PhoneNumber string    `gorm:"type:varchar(20);not null;column:PhoneNumber" json:"phone_number"`
	Role        string    `gorm:"type:varchar(50);column:Role" json:"role"`
}
