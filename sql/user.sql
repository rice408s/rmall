CREATE TABLE User
(
    id        INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username  VARCHAR(50) comment '用户名',
    password  VARCHAR(100) comment '密码',
    mobile    VARCHAR(20) comment '手机号',
    email     VARCHAR(50) comment '邮箱',
    created_at TIMESTAMP comment '创建时间',
    updated_at TIMESTAMP comment '更新时间',
    deleted_at TIMESTAMP comment '删除时间'
);