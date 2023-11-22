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

// Register 用户注册
func Register(req *request.RegisterReq) (int, error) {
	fmt.Println("hello")
	// 参数校验
	if req.Username == "" || req.Password == "" || req.Mobile == "" || req.Email == "" {
		return 0, errors.New("参数错误")
	}
	fmt.Println("hello")
	//检测用户名是否存在
	user, err := dao.FindUserByUsername(req.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
		return 0, err
	}

	if user != nil {
		return 0, errors.New("用户名已存在")
	}
	// 密码加密
	hashPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}
	// 保存用户信息
	user = &model.User{
		Username:  req.Username,
		Password:  hashPwd,
		Mobile:    req.Mobile,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	//插入数据库
	return dao.AddUser(user)

}

func Login(req *request.LoginReq) (string, int64, error) {
	// 参数校验
	if req.Username == "" || req.Password == "" {
		return "", 0, errors.New("参数错误")
	}
	//检测用户名是否存在
	user, err := dao.FindUserByUsername(req.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return "", 0, err
	}
	// 检测密码是否正确
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", 0, errors.New("密码错误")
	}

	token, expire, err := utils.CreateToken(user)
	if err != nil {
		return "", 0, err
	}
	return token, expire, nil

}

func FindUserById(Id int) (user *model.User, err error) {
	return dao.FindUserById(Id)
}

// FindUserByUsername 通过用户名查询用户信息
func FindUserByUsername(username string) (user *model.User, err error) {
	return dao.FindUserByUsername(username)
}
