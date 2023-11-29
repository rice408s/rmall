CREATE TABLE admin
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    username   VARCHAR(50)  NOT NULL comment '用户名', -- unique
    password   VARCHAR(255) NOT NULL comment '密码',   -- hashed password
    email      VARCHAR(255) NOT NULL comment '邮箱',   -- unique
    mobile     VARCHAR(20)  NOT NULL comment '手机号', -- unique
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    deleted_at TIMESTAMP comment '删除时间'
);