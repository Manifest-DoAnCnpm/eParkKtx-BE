package entities

import (
	
)

// field phải viết hoa đầu dòng
// `key:value` : struct tag cho thư viện đọc và mappping đúng giá trị 
type Garage struct {

	// field
    GarageID string `gorm:"primaryKey;type:varchar(20)"`
    GarageName string `gorm:"type:varchar(30);not null"` //add not null
	Size int `gorm:"type:int;not null"`
	UserID string `gorm:"type:varchar(20);not null"`
	
	ParkManagement ParkManagement `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	// ee history 
	//........

}