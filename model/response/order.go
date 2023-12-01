package response

import "rmall/model"

type CreateOrderResp struct {
	Id int64 `json:"id"`
}

type GetOrderListResp struct {
	Total int64          `json:"total"`
	List  []*model.Order `json:"list"`
}

type GetOrderByIdResp struct {
	*model.Order
}

type DeleteOrderResp struct {
	Id int64 `json:"id"`
}
