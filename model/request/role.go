package request

type AddRoleReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateRoleReq struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteRoleReq struct {
	Id int `json:"id" binding:"required"`
}

// GetRoleListReq 获取角色列表结构体
type GetRoleListReq struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
