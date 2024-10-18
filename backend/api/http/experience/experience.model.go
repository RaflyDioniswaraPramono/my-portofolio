package experience

import "time"

type Experience struct {
	Id             int        `gorm:"primaryKey;AUTO_INCREMENT;column:id" json:"id"`
	Title          string     `gorm:"column:" json:""`
	CompanyName    string     `gorm:"column:" json:""`
	CompanyAddress string     `gorm:"column:" json:""`
	WorkAs         string     `gorm:"column:" json:""`
	StartData      *time.Time `gorm:"column:" json:""`
	EndDate        *time.Time `gorm:"column:" json:""`
}
