package dao

import (
	"rmall/global"
	"rmall/model"
)

func AddUser(user *model.User) (id int, err error) {
	insertStr := `insert into user(username,password,mobile,email,created_at) values(?,?,?,?,?)`
	res, err := global.Db.Exec(insertStr, user.Username, user.Password, user.Mobile, user.Email, user.CreatedAt)
	if err != nil {
		return 0, err
	}
	insertId, err := res.LastInsertId()
	return int(insertId), err

}

func FindUserByUsername(username string) (user *model.User, err error) {
	queryStr := `select * from user where username=?`
	user = &model.User{}
	err = global.Db.Get(user, queryStr, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserById(Id int) (user *model.User, err error) {
	queryStr := `select * from user where id=?`
	user = &model.User{}
	err = global.Db.Get(user, queryStr, Id)
	if err != nil {
		return nil, err
	}
	return user, nil

}
