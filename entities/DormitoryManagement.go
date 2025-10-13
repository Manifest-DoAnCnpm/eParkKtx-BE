package entities

import (
	
)

// field phải viết hoa đầu dòng
// `key:value` : struct tag cho thư viện đọc và mappping đúng giá trị 
type DormitoryManagement struct {


	// field
    UserID      string    `gorm:"primaryKey;type:varchar(20)"`
    Building    string    `gorm:"type:varchar(30);not null"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}