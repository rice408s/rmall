package resp

type LoginResp struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type RegisterResp struct {
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
