CREATE TABLE `t_trading_assets_log`
(
    `id`                  bigint(20) NOT NULL AUTO_INCREMENT COMMENT '资产表主键ID',
    `item_code`           varchar(128)    DEFAULT NULL COMMENT '应用编码',
    `assets_name`         varchar(128)    DEFAULT NULL COMMENT '资产名称',
    `assets_code`         varchar(128)    DEFAULT NULL COMMENT '资产编码',
    `line_time_start`     bigint(15)      DEFAULT NULL COMMENT 'k线开始时间',
    `line_time_end`       bigint(15)      DEFAULT NULL COMMENT 'k线结束时间',
    `dot_count`           int(11)         DEFAULT NULL COMMENT '该段时间线的点数',
    `initial_value`       decimal(20, 12) DEFAULT NULL COMMENT '初始值',
    `end_value`           decimal(20, 12) DEFAULT NULL COMMENT '结束值',
    `rise_count`          int(11)         DEFAULT NULL COMMENT '上涨次数',
    `fall_count`          int(11)         DEFAULT NULL COMMENT '下跌次数',
    `flat_count`          int(11)         DEFAULT NULL COMMENT '平滑次数（前后两点相等次数）',
    `rise_rate`           decimal(20, 16)  DEFAULT NULL COMMENT '上涨率',
    `fall_rate`           decimal(20, 16)  DEFAULT NULL COMMENT '下跌率',
    `flat_rate`           decimal(20, 16)  DEFAULT NULL COMMENT '平滑率',
    `result`              tinyint(2)      DEFAULT NULL COMMENT '{0: "平滑", 1: "上涨", 2: "下跌"}',
    `create_date`         datetime        DEFAULT NULL COMMENT '创建日期',
    `update_date`         datetime        DEFAULT NULL COMMENT '更新日期',
    `return_rate_history` double(5, 2)    DEFAULT '0.79' COMMENT '历史收益率',
    `return_rate_now`     double(5, 2)    DEFAULT '0.79' COMMENT '新的收益率',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `item_code` (`item_code`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='资产表历史记录';