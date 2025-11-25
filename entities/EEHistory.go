package entities

import (
	"time"
)

// field phải viết hoa đầu dòng
// `key:value` : struct tag cho thư viện đọc và mappping đúng giá trị
type EEHistory struct {

	// field
	TimeDate    time.Time `gorm:"primaryKey"`
	NumberPlate string    `gorm:"primaryKey"`

	Status   string `gorm:"type:varchar(10);check:status IN ('in','out');not null"`
	GarageID string `gorm:"column:GarageID;not null"`

	Vehicle Vehicle `gorm:"foreignKey:NumberPlate;constraint:OnDelete:CASCADE"`
	Garage  Garage  `gorm:"foreignKey:GarageID;constraint:OnDelete:CASCADE"`
}
