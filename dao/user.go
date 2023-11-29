package dao

import (
	"rmall/global"
	"rmall/model"
)

// AddUser 添加用户
func AddUser(user *model.User) (id int, err error) {
	insertStr := `insert into user(username,password,mobile,email,created_at) values(?,?,?,?,?)`                 // 插入语句
	res, err := global.Db.Exec(insertStr, user.Username, user.Password, user.Mobile, user.Email, user.CreatedAt) //执行插入语句
	if err != nil {
		return 0, err
	}
	insertId, err := res.LastInsertId()
	return int(insertId), err //返回插入成功的id

}

// FindUserByUsername 根据用户名查找用户
func FindUserByUsername(username string) (user *model.User, err error) {
	queryStr := `select * from user where username=?`
	user = &model.User{}
	err = global.Db.Get(user, queryStr, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserById 根据用户id查找用户
func FindUserById(Id int) (user *model.User, err error) {
	queryStr := `select * from user where id=?`
	user = &model.User{}
	err = global.Db.Get(user, queryStr, Id)
	if err != nil {
		return nil, err
	}
	return user, nil

}
