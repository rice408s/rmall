package request

// AdminLoginReq 登陆请求
type AdminLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminRegisterReq 注册请求
type AdminRegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type AdminAssignRoleReq struct {
	AdminId int `json:"admin_id" binding:"required"`
	RoleId  int `json:"role_id" binding:"required"`
}

type GetAdminListReq struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
