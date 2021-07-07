-- ----------------------------
-- 用户数据库
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`          int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `account`     varchar(11) DEFAULT NULL COMMENT '登录账号',
    `password`    varchar(40) DEFAULT NULL COMMENT '登录密码',
    `user_name`   varchar(11) DEFAULT NULL COMMENT '用户名',
    `tele_phone`  varchar(11) DEFAULT NULL COMMENT '电话号码',
    `status`      int(1) DEFAULT NULL COMMENT '用户权限(0系统管理员 1项目管理员 2普通人员)',
    `create_time` varchar(21) DEFAULT NULL COMMENT '创建时间',
    `update_time` varchar(21) DEFAULT NULL COMMENT '修改时间',
    `is_delete`   int(1) DEFAULT NULL COMMENT '逻辑删除(0未删 1删除)',
    PRIMARY KEY (`id`)
)charset utf8 collate utf8_general_ci;
INSERT INTO `users`
VALUES (1, 'system', '7c222fb2927d828af22f592134e8932480637c0d', '系统管理员', '12345678901', 0, '2021-04-25 09:00:00',
        '2021-04-25 09:00:00', 0);
INSERT INTO `users`
VALUES (2, 'admin1', '7c222fb2927d828af22f592134e8932480637c0d', '管理员1', '12345678901', 1, '2021-04-25 09:00:00',
        '2021-04-25 09:00:00', 0);
INSERT INTO `users`
VALUES (3, 'member1', '7c222fb2927d828af22f592134e8932480637c0d', '成员1', '12345678901', 2, '2021-04-25 09:00:00',
        '2021-04-25 09:00:00', 0);
-- ----------------------------
-- 首页总览数据库
-- ----------------------------
DROP TABLE IF EXISTS `overviews`;
CREATE TABLE `overviews`
(
    `all_project`   int DEFAULT 0 COMMENT '项目总数',
    `all_monitor`   int DEFAULT 0 COMMENT '监控总数',
    `all_algorithm` int DEFAULT 0 COMMENT '算法总数',
    `all_warning`   int DEFAULT 0 COMMENT '告警总数'
)charset utf8 collate utf8_general_ci;
INSERT INTO `overviews`
VALUES (1, 2, 2, 2);
-- ----------------------------
-- 算法数据库
-- ----------------------------
DROP TABLE IF EXISTS `algorithms`;
CREATE TABLE `algorithms`
(
    `id`           int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `algorithm_id` varchar(32)  DEFAULT NULL COMMENT '项目id',
    `english_name` varchar(30)  DEFAULT NULL COMMENT '算法英文名',
    `version`      varchar(10)  DEFAULT NULL COMMENT '版本号',
    `type`         varchar(20)  DEFAULT NULL COMMENT '类型',
    `writer`       varchar(20)  DEFAULT NULL COMMENT '作者',
    `description`  varchar(100) DEFAULT NULL COMMENT '算法描述',
    `address`      varchar(100) DEFAULT NULL COMMENT '算法储存地址',
    `create_time`  varchar(21)  DEFAULT NULL COMMENT '创建时间',
    `update_time`  varchar(21)  DEFAULT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`)
)charset utf8 collate utf8_general_ci;
INSERT INTO `algorithms`
VALUES (3, '1234567890abcdefghigklmnopqrstuv', 'test_algorithm1', '1.0.16', 'no_type', 'test_manager',
        'algorithm1 for test',
        'no way', '2021-04-25 09:00:00', '2021-04-25 09:00:00');
INSERT INTO `algorithms`
VALUES (4, 'abcdefghigklmnopqrstuv1234567890', 'test_algorithm2', '1.0.13', '1_type', 'test_manager',
        'algorithm2 for test',
        'no way', '2021-04-25 09:00:00', '2021-04-25 09:00:00');
-- ----------------------------
-- 根任务数据库
-- ----------------------------
DROP TABLE IF EXISTS `root_tasks`;
CREATE TABLE `root_tasks`
(
    `id`                 int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `task_name`          varchar(30)  DEFAULT NULL COMMENT '任务名称',
    `task_id`            varchar(32)  DEFAULT NULL COMMENT '任务id',
    `camera_name`        varchar(30)  DEFAULT NULL COMMENT '摄像头ip地址',
    `camera_ip`          varchar(128) DEFAULT NULL COMMENT '摄像头ip地址',
    `camera_number`      varchar(50)  DEFAULT NULL COMMENT '摄像头序列号',
    `start_time`         varchar(21)  DEFAULT NULL COMMENT '开始时间',
    `end_time`           varchar(21)  DEFAULT NULL COMMENT '结束时间',
    `algorithm_return`   varchar(20)  DEFAULT NULL COMMENT '算法返回类型',
    `algorithm_callback` varchar(100) DEFAULT NULL COMMENT '算法回调地址',
    `is_tour_check`      int(1) DEFAULT NULL COMMENT '是否为巡检任务(0不是 1是)',
    `is_start`           int(1) DEFAULT NULL COMMENT '任务是否开始(0未开始 1开始)',
    `create_time`        varchar(21)  DEFAULT NULL COMMENT '创建时间',
    `update_time`        varchar(21)  DEFAULT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`)
) charset utf8 collate utf8_general_ci;
INSERT INTO `root_tasks`
VALUES (1, 'test_check1', '1234567890abcdefghigklmnopqrstuv', 'camera1', 'localhost', '12345678', '2021-04-25 08:00',
        '2021-04-25 09:00',
        'string', '127.1.1.3', 0, 0, '2021-04-25 09:00:00',
        '2021-04-25 09:00:00');
INSERT INTO `root_tasks`
VALUES (2, 'test_check2', 'abcdefghigklmnopqrstuv1234567890', 'camera2', '127.0.0.2', '12345678', '2021-04-25 08:00',
        '2021-04-25 10:00',
        'string', '128.1.1.2', 1, 1, '2021-04-25 09:00:00',
        '2021-04-25 09:00:00');
-- ----------------------------
-- 子任务数据库
-- ----------------------------
DROP TABLE IF EXISTS `son_tasks`;
CREATE TABLE `son_tasks`
(
    `id`           int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `task_name`    varchar(30)  DEFAULT NULL COMMENT '任务名称',
    `task_id`      varchar(32)  DEFAULT NULL COMMENT '任务id',
    `algorithm_id` varchar(32)  DEFAULT NULL COMMENT '应用的算法id',
    `parent_id`    varchar(32)  DEFAULT NULL COMMENT '父任务id',
    `area`         varchar(100) DEFAULT NULL COMMENT '作用区域',
    `catch_photo`  int          DEFAULT NULL COMMENT '抓图次数',
    `catch_time`   int          DEFAULT NULL COMMENT '持续时间',
    `ptz`          varchar(100) DEFAULT NULL COMMENT '云台参数',
    PRIMARY KEY (`id`)
) charset utf8 collate utf8_general_ci;
INSERT INTO `son_tasks`
VALUES (1, 'test_son_check1', '12345678911111111higklmnopqfwadg', '算法id1', '1234567890abcdefghigklmnopqrstuv',
        '37,61,29,18', 17, 30, 'ptz1');
INSERT INTO `son_tasks`
VALUES (2, 'test_son_check2', 'abcdefd22222222opqrstuv12sdf7890', '算法id2', '1234567890abcdefghigklmnopqrstuv',
        '17,6,3,178', 15, 40, 'ptz1');
INSERT INTO `son_tasks`
VALUES (3, 'test_son_check3', '234567890abcq33333igklmnopqfwadg', '算法id1', 'abcdefghigklmnopqrstuv1234567890',
        '37,34,9,18', 30, 30, 'ptz0');
INSERT INTO `son_tasks`
VALUES (4, 'test_son_check4', 'bcdefdghigklmn244444tuv12sdf7890', '算法id2', 'abcdefghigklmnopqrstuv1234567890',
        '17,6,3123,178', 10, 70, 'ptz0');
-- ----------------------------
-- 告警数据库
-- ----------------------------
DROP TABLE IF EXISTS `warnings`;

CREATE TABLE `warnings`
(
    `id`             int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `warning_id`     varchar(32) DEFAULT NULL COMMENT '项目id',
    `camera_number`  varchar(50) DEFAULT NULL COMMENT '摄像头序列号',
    `warning_type`   varchar(10) DEFAULT NULL COMMENT '告警类型',
    `warning_time`   int         DEFAULT NULL COMMENT '告警时间',
    `warning_degree` int(1) DEFAULT NULL COMMENT '等级',
    PRIMARY KEY (`id`)
) charset utf8 collate utf8_general_ci;
INSERT INTO `warnings`
VALUES (1, 'abcdefghigklmnopqrstuv1234567890', 'camera1', 'warning1',
        12345, 0);
INSERT INTO `warnings`
VALUES (2, '1234567890abcdefghigklmnopqrstuv', 'camera2', 'warning',
        123456, 0);
-- ----------------------------
-- IPC数据库
-- ----------------------------
DROP TABLE IF EXISTS `ipcs`;
CREATE TABLE `ipcs`
(
    `id`            int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `number`        varchar(50)  DEFAULT NULL COMMENT '设备序列号',
    `name`          varchar(30)  DEFAULT NULL COMMENT '设备名称',
    `account`       varchar(30)  DEFAULT NULL COMMENT '设备账号',
    `password`      varchar(30)  DEFAULT NULL COMMENT '设备密码',
    `ip`            varchar(128) DEFAULT NULL COMMENT '设备ip',
    `nvr`           varchar(128) DEFAULT NULL COMMENT '关联的nvr序列号',
    `factory`       varchar(10)  DEFAULT NULL COMMENT '厂商',
    `type`          varchar(20)  DEFAULT NULL COMMENT '设备型号',
    `area`          varchar(20)  DEFAULT NULL COMMENT '安装区域',
    `place`         varchar(20)  DEFAULT NULL COMMENT '安装位置',
    `is_online`     int(1) DEFAULT NULL COMMENT '是否在线(0在线 1不在线)',
    `is_tour_check` int(1) DEFAULT NULL COMMENT '是否能进行巡检(0不能 1能)',
    `is_start`      int(1) DEFAULT NULL COMMENT '是否激活(0激活 1未激活)',
    PRIMARY KEY (`id`)
) charset utf8 collate utf8_general_ci;
INSERT INTO `ipcs`
VALUES (1, '1234567890abcdefghigklmnopqrstuv', 'camera1', '12345678', '12345678',
        '10.0.120.147', '1234567890abcdefghigklmnopqrstuv', '海信', '1.3.8NCR', '西区', '房间顶', 0, 1, 0);
INSERT INTO `ipcs`
VALUES (2, 'abcdefghigklmnopqrstuv1234567890', 'camera2', '12345678', '12345678',
        '10.0.120.148', 'abcdefghigklmnopqrstuv1234567890', '海信', '1.3.66NCR', '东区', '门口', 0, 0, 0);
INSERT INTO `ipcs`
VALUES (3, 'abcdefghigklmnopqrstuv123456test', 'camera_test', 'admin', 'password.123',
        '10.0.120.146', 'abcdefghigklmnopqrstuv1234567890', '海信', '1.3.66NCR', '主楼', '1005', 1, 0, 0);
-- ----------------------------
-- NVR数据库
-- ----------------------------
DROP TABLE IF EXISTS `nvrs`;
CREATE TABLE `nvrs`
(
    `id`        int NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `number`    varchar(50)  DEFAULT NULL COMMENT '设备序列号',
    `name`      varchar(30)  DEFAULT NULL COMMENT '设备名称',
    `account`   varchar(30)  DEFAULT NULL COMMENT '设备账号',
    `password`  varchar(30)  DEFAULT NULL COMMENT '设备密码',
    `ip`        varchar(128) DEFAULT NULL COMMENT '设备ip',
    `factory`   varchar(10)  DEFAULT NULL COMMENT '厂商',
    `type`      varchar(20)  DEFAULT NULL COMMENT '设备型号',
    `is_online` int(1) DEFAULT NULL COMMENT '是否在线(0在线 1不在线)',
    PRIMARY KEY (`id`)
) charset utf8 collate utf8_general_ci;
INSERT INTO `nvrs`
VALUES (1, '1234567890abcdefghigklmnopqrstuv', 'nvr1', '12345678', '12345678',
        '127.101.8.3', '海信', '1.3.8nvr', 0);
INSERT INTO `nvrs`
VALUES (2, 'abcdefghigklmnopqrstuv1234567890', 'nvr2', '12345678', '12345678',
        '127.101.8.31', '海信', '1.3.66nvr', 0);