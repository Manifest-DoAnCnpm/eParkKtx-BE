package entities

import (
	"time"
)

// field phải viết hoa đầu dòng
// `key:value` : struct tag cho thư viện đọc và mappping đúng giá trị 
type Contract struct {

	// field
    ContractID string `gorm:"primaryKey;type:varchar(20)"`;
	StartDate time.Time `gorm:"type:date;not null"`;
	EndDate time.Time `gorm:"type:date;not null"`;

	ContractType string `gorm:"type:varchar(30);not null"`;
	Cost int `gorm:"type:bigint;not null"`;

	NumberPlate Vehicle `gorm:"foreignKey:NumberPlate"`
	ParkManagementID ParkManagement `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

}