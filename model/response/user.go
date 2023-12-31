package response

type UserLoginResp struct {
	Token string `json:"token"`
	// 过期时间
	Expire int64 `json:"expire"`
}

type UserRegisterResp struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}
