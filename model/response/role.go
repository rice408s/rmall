package response

import "rmall/model"

// AddRoleResp role添加结构体
type AddRoleResp struct {
	Id int `json:"id"`
}

// UpdateRoleResp role更新结构体
type UpdateRoleResp struct {
	Id int `json:"id"`
}

type DeleteRoleResp struct {
	Id int `json:"id"`
}

// FindRoleResp 查询角色结构体
type FindRoleResp struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`
}

// GetRoleListResp 获取角色列表结构体
type GetRoleListResp struct {
	Total int           `json:"total"`
	List  []*model.Role `json:"list"`
}
