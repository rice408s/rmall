package service

import (
	"fmt"
	"rmall/global"
	"rmall/model/request"
	"strconv"
)

// AddPolicy 添加策略
func AddPolicy(req *request.AddPolicyReq) (bool, error) {
	return global.E.AddPolicy(req.V0, req.V1, req.V2)
}

func RemovePolicy(req *request.RemovePolicyReq) (bool, error) {
	return global.E.RemovePolicy(req.V0, req.V1, req.V2)
}

func UpdatePolicy(req *request.UpdatePolicyReq) (bool, error) {
	oldPolicy := []string{req.OldV0, req.OldV1, req.OldV2}
	newPolicy := []string{req.NewV0, req.NewV1, req.NewV2}
	fmt.Println(oldPolicy, newPolicy)
	return global.E.UpdatePolicy(oldPolicy, newPolicy)
}

// GetPolicy 获取策略
func GetPolicy() [][]string {
	return global.E.GetPolicy()
}

// GetPolicyByRole 通过角色获取策略
func GetPolicyByRole(req *request.GetPolicyByRoleReq) [][]string {
	return global.E.GetFilteredPolicy(0, strconv.Itoa(req.Id))
}

// GetPolicyByPath 通过路径获取策略
func GetPolicyByPath(path string) [][]string {
	return global.E.GetFilteredPolicy(1, path)
}

// GetPolicyByMethod 通过方法获取策略
func GetPolicyByMethod(method string) [][]string {
	return global.E.GetFilteredPolicy(2, method)
}

// GetAllRole 获取所有角色
func GetAllRole() []string {
	return global.E.GetAllRoles()
}

// RemovePolicyByRole 删除角色对应管理员的策略
func RemovePolicyByRole(role string) (bool, error) {
	users, err := global.E.GetUsersForRole(role)
	if err != nil {
		return false, err
	}
	for _, v := range users {
		ok, err := global.E.RemoveFilteredPolicy(0, v, role)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}
