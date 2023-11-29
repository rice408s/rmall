CREATE TABLE `role` (
                        `id` INT NOT NULL AUTO_INCREMENT,
                        `role_name` VARCHAR(100) NOT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE INDEX `role_name_UNIQUE` (`role_name` ASC) VISIBLE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';