package bio

import (
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/user"
)

type Bio struct {
	Id            int        `gorm:"primaryKey;AUTO_INCREMENT;not null;column:id" json:"id"`
	UserId        int        `gorm:"column:user_id" json:"user_id"`
	User          *user.User `gorm:"foreignKey:UserId;references:id"`
	FullName      string     `gorm:"column:full_name" json:"full_name"`
	Age           int        `gorm:"column:age" json:"age"`
	Address       string     `gorm:"column:address" json:"address"`
	LastEducation string     `gorm:"column:last_education" json:"last_education"`
	Job           string     `gorm:"column:job" json:"job"`
	Photo         string     `gorm:"column:photo" json:"photo"`
}
