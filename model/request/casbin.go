package request

// AddPolicyReq 添加策略
type AddPolicyReq struct {
	V0 string `json:"v0" binding:"required"`
	V1 string `json:"v1" binding:"required"`
	V2 string `json:"v2"`
}

// RemovePolicyReq 删除策略
type RemovePolicyReq struct {
	V0 string `json:"v0" binding:"required"`
	V1 string `json:"v1" binding:"required"`
	V2 string `json:"v2"`
}

// UpdatePolicyReq 修改策略
type UpdatePolicyReq struct {
	OldV0 string `json:"old_v0" binding:"required"`
	OldV1 string `json:"old_v1" binding:"required"`
	OldV2 string `json:"old_v2"`
	NewV0 string `json:"new_v0" binding:"required"`
	NewV1 string `json:"new_v1" binding:"required"`
	NewV2 string `json:"new_v2"`
}

type GetPolicyByRoleReq struct {
	Id int `json:"id" form:"id"`
}
