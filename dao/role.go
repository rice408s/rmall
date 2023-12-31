package dao

import (
	"fmt"
	"rmall/global"
	"rmall/model"
	"rmall/utils"
	"strconv"
)

// AddRole 添加角色
func AddRole(role *model.Role) (id int, err error) {
	insertStr := `insert into role(role_name) values(?)` // 插入语句
	res, err := global.DB.Exec(insertStr, role.RoleName) //执行插入语句
	if err != nil {
		return 0, err
	}
	insertId, err := res.LastInsertId()
	return int(insertId), err //返回插入成功的id
}

// FindRoleByRoleName 根据角色名查找角色
func FindRoleByRoleName(roleName string) (role *model.Role, err error) {
	queryStr := `select * from role where role_name=?`
	role = &model.Role{}
	err = global.DB.Get(role, queryStr, roleName)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// FindRoleById 根据角色id查找角色
func FindRoleById(Id int) (role *model.Role, err error) {
	queryStr := `select * from role where id=?`
	role = &model.Role{}
	err = global.DB.Get(role, queryStr, Id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// FindRoleByPage 查找所有角色
func FindRoleByPage(offset, pageSize int) (roles []*model.Role, err error) {
	queryStr := `select * from role` + " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa(offset)
	roles = []*model.Role{}
	fmt.Println(queryStr)
	err = global.DB.Select(&roles, queryStr)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// UpdateRole 更新角色
func UpdateRole(role *model.Role) (err error) {
	updateStr := `update role set role_name=? where id=?`
	_, err = global.DB.Exec(updateStr, role.RoleName, role.Id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoleById 根据角色id删除角色
func DeleteRoleById(Id int) (err error) {
	deleteStr := `delete from role where id=?`
	_, err = global.DB.Exec(deleteStr, Id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoleByRoleName 根据角色名删除角色
func DeleteRoleByRoleName(roleName string) (err error) {
	deleteStr := `delete from role where role_name=?`
	_, err = global.DB.Exec(deleteStr, roleName)
	if err != nil {
		return err
	}
	return nil
}

// FindRoleByAdminId 根据管理员id查找角色
func FindRoleByAdminId(adminId int) (roles []*model.Role, err error) {
	//查询出所有的角色的id
	roleIds, err := global.E.GetRolesForUser(strconv.Itoa(adminId))
	if err != nil {
		return nil, err
	}
	intArray, err := utils.ConvertStringArrayToIntArray(roleIds)
	if err != nil {
		return nil, err
	}

	//根据角色id查询出所有的角色
	roles = []*model.Role{}
	for _, roleId := range intArray {
		role, err := FindRoleById(roleId)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
