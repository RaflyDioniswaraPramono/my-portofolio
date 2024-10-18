package portofolio

import "time"

type Portofolio struct {
	Id          int        `gorm:"primaryKey;AUTO_INCREMENT;column:id" json:"id"`
	Title       string     `gorm:"column:" json:""`
	Description string     `gorm:"column:" json:""`
	Thumbnail   string     `gorm:"column:" json:""`
	CreatedAt   *time.Time `gorm:"column:" json:""`
}
