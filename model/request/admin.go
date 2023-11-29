package request

type AdminLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminRegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
