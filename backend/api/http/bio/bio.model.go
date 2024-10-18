package bio

import (
	"github.com/RaflyDioniswaraPramono/my-portofolio/api/http/user"
)

type Bio struct {
	UserId        int
	User          *user.User
	FullName      string
	Age           int
	Address       string
	LastEducation string
	Job           string
	Photo         string
}
