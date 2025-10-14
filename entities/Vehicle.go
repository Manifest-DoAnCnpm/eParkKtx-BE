package entities

import (
	"time"
)

// field phải viết hoa đầu dòng
// `key:value` : struct tag cho thư viện đọc và mappping đúng giá trị 
type Vehicle struct {

	// field
    NumberPlate string `gorm:"primaryKey;type:varchar(20)"`
    VehicleType string `gorm:"type:varchar(30)"`
	RegisterDate time.Time `gorm:"type:date;not null"`
	Color string `gorm:"type:varchar(20)"`

	StudentID Student `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE"`
	ParkManagementID ParkManagement `gorm:"foreignKey:ParkManagementID;constraint:OnDelete:CASCADE"`

}