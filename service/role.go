package service

import (
	"rmall/dao"
	"rmall/model"
	"rmall/model/request"
)

// Path: service/role.go
// Compare this snippet from dao/role.go:
// package dao
//

// AddRole 添加角色
func AddRole(req *request.AddRoleReq) (id int, err error) {
	role := &model.Role{
		RoleName: req.Name,
	}
	return dao.AddRole(role)
}

// UpdateRole 更新角色
func UpdateRole(req *request.UpdateRoleReq) (err error) {
	role := &model.Role{
		Id:       req.Id,
		RoleName: req.Name,
	}
	return dao.UpdateRole(role)
}

// DeleteRole 删除角色
func DeleteRole(id int) (err error) {
	return dao.DeleteRoleById(id)
}

// FindRoleByPage 查找所有角色
func FindRoleByPage(req *request.GetRoleListReq) (roles []*model.Role, err error) {
	// 分页
	start := (req.PageNum - 1) * req.PageSize
	// 从数据库中查找所有角色
	return dao.FindRoleByPage(start, req.PageSize)
}
