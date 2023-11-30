package dao

import (
	"rmall/global"
	"rmall/model"
	"strconv"
)

// AddAdmin 添加管理员
func AddAdmin(admin *model.Admin) (id int, err error) {
	insertStr := `insert into admin(username,password,mobile,email,created_at) values(?,?,?,?,?)`
	res, err := global.Db.Exec(insertStr, admin.Username, admin.Password, admin.Mobile, admin.Email, admin.CreatedAt)
	if err != nil {
		return 0, err
	}
	insertId, err := res.LastInsertId()
	return int(insertId), err
}

// FindAdminByUsername 根据用户名查找管理员
func FindAdminByUsername(username string) (admin *model.Admin, err error) {
	queryStr := `select * from admin where username=?`
	admin = &model.Admin{}
	err = global.Db.Get(admin, queryStr, username)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// FindAdminById 根据管理员id查找管理员
func FindAdminById(Id int) (admin *model.Admin, err error) {
	queryStr := `select * from admin where id=?`
	admin = &model.Admin{}
	err = global.Db.Get(admin, queryStr, Id)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func FindAdminByPage(offset, pageSize int) (admins []*model.Admin, err error) {
	queryStr := `select * from admin` + " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa(offset)
	admins = []*model.Admin{}
	err = global.Db.Select(&admins, queryStr)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

//func FindAdminRoleByAdminId(adminId int) (role *model.Role, err error) {
//
//}
