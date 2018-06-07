package models

type UserType uint

const (
	AdminUser   UserType = 1
	ManagerUser UserType = 2
	NormalUser  UserType = 3
)

type User struct {
	ID   uint
	Name string
	Type UserType
}
