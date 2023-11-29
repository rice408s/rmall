package response

type AdminLoginResp struct {
	Token string `json:"token"`
	// 过期时间
	Expire int64 `json:"expire"`
}

type AdminRegisterResp struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

type AdminInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}
