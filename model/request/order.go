package request

type CreateOrderReq struct {
	Uid    int64 `json:"uid" binding:"required"`
	Pid    int64 `json:"pid" binding:"required"`
	Amount int   `json:"amount" binding:"required"`
}

type GetOrderListReq struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type GetOrderByIdReq struct {
	Id int `json:"id" form:"id"`
}

type DeleteOrderReq struct {
	Id int `json:"id" form:"id"`
}

type UpdateOrderReq struct {
	Id     int64 `json:"id" binding:"required"`
	Uid    int64 `json:"uid" binding:"required"`
	Pid    int64 `json:"pid" binding:"required"`
	Amount int   `json:"amount" binding:"required"`
	Status int   `json:"status" binding:"required"`
}

type GetOrderByUidReq struct {
	Uid int `json:"uid" form:"uid"`
}

type GetOrderByPidReq struct {
	Pid int `json:"pid" form:"pid"`
}
