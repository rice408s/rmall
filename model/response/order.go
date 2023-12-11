package response

import "rmall/model"

type CreateOrderResp struct {
	Id int `json:"id"`
}

type GetOrderListResp struct {
	Total int            `json:"total"`
	List  []*model.Order `json:"list"`
}

type GetOrderByIdResp struct {
	*model.Order
}

type DeleteOrderResp struct {
	Id int `json:"id"`
}

type UpdateOrderResp struct {
	Id int `json:"id"`
}

type GetOrderByPageResp struct {
	Total int            `json:"total"`
	List  []*model.Order `json:"list"`
}

type GetOrderByUidResp struct {
	Total int            `json:"total"`
	List  []*model.Order `json:"list"`
}

type GetOrderByPidResp struct {
	Total int            `json:"total"`
	List  []*model.Order `json:"list"`
}
