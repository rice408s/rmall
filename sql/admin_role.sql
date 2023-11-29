CREATE TABLE `admin_role` (
                              `admin_id` INT NOT NULL,
                              `role_id` INT NOT NULL,
                              PRIMARY KEY (`admin_id`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员角色关联表';