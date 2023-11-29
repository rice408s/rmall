package model

// Role 定义角色结构体
type Role struct {
	Id       int    `json:"id" db:"id"`
	RoleName string `json:"role_name" db:"role_name"`
}
