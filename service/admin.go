package service

import (
	"database/sql"
	"errors"
	"fmt"
	"rmall/dao"
	"rmall/global"
	"rmall/model"
	"rmall/model/request"
	"rmall/utils"
	"strconv"
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
	return utils.CreateAdminToken(admin)
}

// GetAdminList 获取管理员列表
func GetAdminList(req *request.GetAdminListReq) ([]*model.Admin, error) {
	start := (req.PageNum - 1) * req.PageSize
	return dao.FindAdminByPage(start, req.PageSize)
}

// AdminAssignRole 管理员分配角色
func AdminAssignRole(req *request.AdminAssignRoleReq) (bool, error) {
	// 参数校验
	if req.AdminId == 0 || req.RoleId == 0 {
		return false, errors.New("参数错误")
	}
	// 检测管理员是否存在
	admin, err := dao.FindAdminById(req.AdminId)
	if err != nil {
		return false, err
	}
	if admin == nil {
		return false, errors.New("管理员不存在")
	}
	// 检测角色是否存在
	role, err := dao.FindRoleById(req.RoleId)
	if err != nil {
		return false, err
	}
	if role == nil {
		return false, errors.New("角色不存在")
	}
	// 检测管理员是否已经分配了角色
	roles, err := global.E.GetRolesForUser(strconv.Itoa(req.AdminId))
	if err != nil {
		return false, err
	}
	//遍历判断是否已经分配了此角色
	for _, v := range roles {
		if v == strconv.Itoa(role.Id) {
			return false, errors.New("此管理员已经分配了此角色")
		}
	}
	// 给管理员分配角色
	ok, err := global.E.AddRoleForUser(strconv.Itoa(req.AdminId), strconv.Itoa(req.RoleId))
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("分配角色失败")
	}
	return true, nil
}
