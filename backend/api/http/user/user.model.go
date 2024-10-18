package user

import "github.com/RaflyDioniswaraPramono/my-portofolio/api/http/role"

type User struct {
	Id       int       `gorm:"primaryKey;AUTO_INCREMENT;column:id" json:"id"`
	Email    string    `gorm:"unique;column:email" json:"email"`
	Username string    `gorm:"unique;column:username" json:"username"`
	Password string    `gorm:"column:password" json:"password"`
	RoleId   int       `gorm:"column:role_id" json:"role_id"`
	Role     role.Role `gorm:"foreignKey:RoleId;references:id"`
}

type UserSignInDto struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}
