package entities

import (
	
	"time"
)

// `key:value` : tag cho thư viện đọc và mappping đúng giá trị 
type User struct {

	// gorm.Model: nhúng các trường id, creatAt,.... vào struct
	// gorm.Model

	// field
    UserID      string    `gorm:"primaryKey;type:varchar(20)"`
    Password    string    `gorm:"type:varchar(100);not null"`
    Name        string    `gorm:"type:varchar(100);not null"`
    DoB         time.Time `json:"dob"`
    Gender      string    `gorm:"type:varchar(10)"`
    PhoneNumber string    `gorm:"type:varchar(20);not null"`

}