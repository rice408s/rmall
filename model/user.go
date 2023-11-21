package model

import "time"

// User 定义用户结构体
type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	Roles     []Role    `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
