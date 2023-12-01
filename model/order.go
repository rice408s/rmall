package model

import "time"

//CREATE TABLE `order`
//(
//`id`          bigint unsigned NOT NULL AUTO_INCREMENT,
//`uid`         bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
//`pid`         bigint unsigned NOT NULL DEFAULT '0' COMMENT '产品ID',
//`amount`      int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单金额',
//`status`      tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态',
//`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//PRIMARY KEY (`id`),
//KEY           `idx_uid` (`uid`),
//KEY           `idx_pid` (`pid`)
//) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
//

// Order 订单
type Order struct {
	Id         int64     `json:"id" db:"id"`
	Uid        int64     `json:"uid" db:"uid"`
	Pid        int64     `json:"pid" db:"pid"`
	Amount     int       `json:"amount" db:"amount"`
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
