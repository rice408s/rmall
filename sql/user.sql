CREATE TABLE User (
                      Id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
                      Username VARCHAR(50) comment '用户名',
                      Password VARCHAR(50) comment '密码',
                      Mobile VARCHAR(20) comment '手机号',
                      Email VARCHAR(50) comment '邮箱',
                      CreatedAt TIMESTAMP comment '创建时间',
                      UpdatedAt TIMESTAMP comment '更新时间'
);