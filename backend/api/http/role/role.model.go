package role

type Role struct {
	Id       int    `gorm:"primaryKey;AUTO_INCREMENT;column:id" json:"id"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
}
