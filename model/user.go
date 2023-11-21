package model

import (
	"database/sql"
	"time"
)

// User 定义用户结构体
type User struct {
	Id        int          `json:"id" db:"id"`
	Username  string       `json:"username" db:"username"`
	Password  string       `json:"password" db:"password"`
	Mobile    string       `json:"mobile" db:"mobile"`
	Email     string       `json:"email" db:"email"`
	Roles     []Role       `json:"roles"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	//删除
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
