CREATE DATABASE `casbin_test` DEFAULT CHARACTER SET utf8;
GRANT Alter, Alter Routine, Create, Create Routine, Create Temporary Tables, Create View, Delete, Drop, Event, Execute, Index, Insert, Lock Tables, References, Select, Show View, Trigger, Update ON `casbin\_test`.* TO `ops`@`%`;
FLUSH PRIVILEGES;
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
                               `p_type` varchar(100) DEFAULT NULL COMMENT '规则类型',
                               `v0` varchar(100) DEFAULT NULL COMMENT '角色ID',
                               `v1` varchar(100) DEFAULT NULL COMMENT 'api路径',
                               `v2` varchar(100) DEFAULT NULL COMMENT 'api访问方法',
                               `v3` varchar(100) DEFAULT '',
                               `v4` varchar(100) DEFAULT '',
                               `v5` varchar(100) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限规则表';
/*插入操作casbin api的权限规则*/
# INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`) VALUES ('p', 'admin', '/api/v1/casbin', 'POST');
# INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`) VALUES ('p', 'admin', '/api/v1/casbin/list', 'GET');