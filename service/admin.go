package service

import (
	"database/sql"
	"errors"
	"fmt"
	"rmall/dao"
	"rmall/model"
	"rmall/model/request"
	"rmall/utils"
	"time"
)

// AdminRegister 管理员注册
func AdminRegister(req *request.AdminRegisterReq) (int, error) {
	// 参数校验
	if req.Username == "" || req.Password == "" || req.Mobile == "" || req.Email == "" {
		return 0, errors.New("参数错误")
	}
	//检测管理员名是否存在
	admin, err := dao.FindAdminByUsername(req.Username)
	fmt.Println(admin)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	// 检测用户名是否存在
	if admin != nil {
		return 0, errors.New("用户名已存在")
	}
	// 密码加密
	hashPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}
	// 保存管理员信息
	admin = &model.Admin{
		Username:  req.Username,
		Password:  hashPwd,
		Mobile:    req.Mobile,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
	//插入数据库
	return dao.AddAdmin(admin)

}

// AdminLogin 管理员登录
func AdminLogin(req *request.AdminLoginReq) (string, int64, error) {
	// 参数校验
	if req.Username == "" || req.Password == "" {
		return "", 0, errors.New("参数错误")
	}
	//检测用户名是否存在
	admin, err := dao.FindAdminByUsername(req.Username)
	if err != nil {
		return "", 0, err
	}
	// 检测密码是否正确
	if !utils.CheckPasswordHash(req.Password, admin.Password) {
		return "", 0, errors.New("密码错误")
	}
	// 生成token
	return utils.CreateToken(admin)
}
