CREATE TABLE `t_trading_assets_source`
(
    `id`               int(11) NOT NULL AUTO_INCREMENT COMMENT '数据源表主键ID',
    `item_code`        varchar(128) DEFAULT NULL COMMENT '应用编码',
    `assets_code`      varchar(128) DEFAULT NULL COMMENT '资产编码',
    `kline_source`     varchar(64)  DEFAULT NULL COMMENT 'k线来源,数据源',
    `subscribe_symbol` varchar(128) DEFAULT NULL COMMENT '数据源对应的订阅符号',
    `sleep_time`       varchar(128) DEFAULT NULL COMMENT '休市时间 {"start":"6-04:00","ts_long":"172800"} start开始时间，从周几开始几点开始，休息多少秒',
    `is_valid`         int(3)       DEFAULT '1' COMMENT '是否有效 1是 0否',
    `remark`           varchar(255) DEFAULT NULL COMMENT '数据源备注',
    `creator`          varchar(128) DEFAULT NULL COMMENT '创建人',
    `updater`          varchar(128) DEFAULT NULL COMMENT '更新人',
    `create_date`      datetime     DEFAULT NULL COMMENT '创建日期',
    `update_date`      datetime     DEFAULT NULL COMMENT '更新日期',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `item_code` (`item_code`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 31
  DEFAULT CHARSET = utf8mb4 COMMENT ='数据源表，一个数据源对应多个资产，不同的数据源有不同的休市时间';

INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (1, 'IQO', 'EUR-USD', 'raceoption', 'EUR/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '欧元/美元',
        'system', '叶志明', '2022-05-20 15:20:54', '2022-05-28 10:39:59');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (2, 'IQO', 'GBP-USD', 'raceoption', 'GBP/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '英镑/美元',
        'system', 'tdw', '2022-05-20 15:20:54', '2022-05-28 11:08:56');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (3, 'IQO', 'AUD-CAD', 'raceoption', 'AUD/CAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/加拿大元',
        'system', 'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:09:15');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (4, 'IQO', 'USD-JPY', 'raceoption', 'USD/JPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '美元/日元',
        'system', 'deimen', '2022-05-20 15:20:54', '2022-05-28 09:03:50');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (5, 'IQO', 'GOLD', 'raceoption', 'XAU/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '伦敦金', 'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (6, 'IQO', 'SILVER', 'raceoption', 'XAG/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '伦敦银',
        'system', 'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (7, 'IQO', 'EUR-USD', 'qxbroker', 'EURUSD', NULL, 0, '欧元/美元', 'system', '叶志明', '2022-05-20 15:20:54',
        '2022-05-28 10:39:59');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (8, 'IQO', 'GBP-USD', 'qxbroker', 'GBPUSD', NULL, 0, '英镑/美元', 'system', 'tdw', '2022-05-20 15:20:54',
        '2022-05-28 11:08:56');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (9, 'IQO', 'AUD-CAD', 'qxbroker', 'AUDCAD', NULL, 1, '澳大利亚元/加拿大元', 'system', 'Lucy', '2022-05-20 15:20:54',
        '2022-05-28 09:09:15');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (10, 'IQO', 'USD-JPY', 'qxbroker', 'USDJPY', NULL, 0, '美元/日元', 'system', 'deimen', '2022-05-20 15:20:54',
        '2022-05-28 09:03:50');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (11, 'IQO', 'GOLD', 'qxbroker', 'XAUUSD', NULL, 0, '伦敦金', 'system', 'deimen', '2022-05-20 15:20:54',
        '2022-05-28 09:03:50');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (12, 'IQO', 'SILVER', 'qxbroker', 'XAGUSD', NULL, 0, '伦敦银', 'system', 'deimen', '2022-05-20 15:20:54',
        '2022-05-28 09:03:50');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (13, 'IQO', 'Microsoft', 'qxbroker', 'MSFT_otc', NULL, 1, '微软', 'system', 'tdw', '2022-05-20 15:20:54',
        '2022-05-27 17:31:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (14, 'IQO', 'Intel', 'qxbroker', 'INTC_otc', NULL, 1, '英特尔', 'system', 'tdw', '2022-05-20 15:20:54',
        '2022-05-27 17:31:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (15, 'IQO', 'EUR-USD', 'olymptrade', 'EURUSD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/美元',
        'system', '叶志明', '2022-05-20 15:20:54', '2022-05-28 10:39:59');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (16, 'IQO', 'GBP-USD', 'olymptrade', 'GBPUSD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/美元',
        'system', 'tdw', '2022-05-20 15:20:54', '2022-05-28 11:08:56');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (17, 'IQO', 'AUD-CAD', 'olymptrade', 'AUDCAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/加拿大元',
        'system', 'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (18, 'IQO', 'USD-JPY', 'olymptrade', 'USDJPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/日元',
        'system', 'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (19, 'IQO', 'GOLD', 'olymptrade', 'XAUUSD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '伦敦金', 'system',
        'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (20, 'IQO', 'SILVER', 'olymptrade', 'XAGUSD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '伦敦银',
        'system', 'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (21, 'IQO', 'BTC-USD', 'huobi', 'BTC-USD', NULL, 0, '比特币/美元', 'system', 'lauren', '2022-05-20 15:20:54',
        '2022-05-28 15:23:43');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (22, 'IQO', 'ETH-USD', 'huobi', 'ETH-USD', NULL, 0, '以太坊/美元', 'system', 'lauren', '2022-05-20 15:20:54',
        '2022-05-28 15:23:43');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (23, 'IQO', 'BTC-USD', 'binance', 'BTC-USD', NULL, 1, '比特币/美元', 'system', 'lauren', '2022-05-20 15:20:54',
        '2022-05-28 15:23:43');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (24, 'IQO', 'ETH-USD', 'binance', 'ETH-USD', NULL, 1, '以太坊/美元', 'system', 'lauren', '2022-05-20 15:20:54',
        '2022-05-28 15:23:43');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (25, 'IQO', 'EUR-USD', 'hsb', '153080', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/美元', 'system',
        '叶志明', '2022-05-20 15:20:54', '2022-05-28 10:39:59');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (26, 'IQO', 'GBP-USD', 'hsb', '153086', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/美元', 'system',
        'tdw', '2022-05-20 15:20:54', '2022-05-28 11:08:56');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (27, 'IQO', 'USD-JPY', 'hsb', '153083', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/日元', 'system',
        'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (28, 'IQO', 'GOLD', 'hsb', '153105', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '伦敦金', 'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (29, 'IQO', 'SILVER', 'hsb', '153104', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '伦敦银', 'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');



INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (30, 'IQO', 'Apple', 'olymptrade', 'AAPL', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '苹果',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (31, 'IQO', 'Tesla', 'olymptrade', 'TSLA', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '特斯拉',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (32, 'IQO', 'CAD-JPY', 'olymptrade', 'CADJPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '加拿大元/日元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (33, 'IQO', 'EUR-AUD', 'olymptrade', 'EURAUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/澳大利亚元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (34, 'IQO', 'USD-CAD', 'olymptrade', 'USDCAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/加拿大元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');



INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (35, 'IQO', 'CAD-JPY', 'raceoption', 'CAD/JPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '加拿大元/日元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (36, 'IQO', 'EUR-AUD', 'raceoption', 'EUR/AUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/澳大利亚元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (37, 'IQO', 'USD-CAD', 'raceoption', 'USD/CAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/加拿大元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');



INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (38, 'IQO', 'CAD-JPY', 'qxbroker', 'CADJPY_otc', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '加拿大元/日元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (39, 'IQO', 'EUR-AUD', 'qxbroker', 'EURAUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/澳大利亚元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (40, 'IQO', 'USD-CAD', 'qxbroker', 'USDCAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/加拿大元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');



INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (41, 'IQO', 'EUR-AUD', 'hsb', '153084', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/澳大利亚元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`,
                                      `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`,
                                      `creator`, `updater`, `create_date`, `update_date`)
VALUES (42, 'IQO', 'USD-CAD', 'hsb', '153088', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/加拿大元',
        'system',
        'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');

-- 新增四条
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (43, 'IQO', 'Adidas', 'olymptrade', 'Adidas', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '阿迪达斯', 'system', 'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (44, 'IQO', 'NGAS', 'olymptrade', 'NGAS', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, 'NGAS', 'system', 'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (45, 'IQO', 'BRENT', 'olymptrade', 'BRENT', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, 'BRENT', 'system', 'system', '2022-05-20 15:20:54', '2022-05-20 15:20:54');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (46, 'IQO', 'Microsoft', 'olymptrade', 'MSFT', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '微软', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-06-16 14:54:04');

-- 新增四个数据源
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (47, 'IQO', 'Amazon', 'olymptrade', 'AMZN', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '亚马逊', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (48, 'IQO', 'AUD-CHF', 'olymptrade', 'AUDCHF', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '澳大利亚元/瑞士法郎', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (49, 'IQO', 'AUD-JPY', 'olymptrade', 'AUDJPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '澳大利亚元/日元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (50, 'IQO', 'GBP-AUD', 'olymptrade', 'GBPAUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '英镑/澳大利亚元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');

INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (51, 'IQO', 'AUD-JPY', 'hsb', '153087', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/日元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (52, 'IQO', 'GBP-AUD', 'hsb', '153092', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/澳大利亚元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (53, 'IQO', 'AUD-CHF', 'qxbroker', 'AUDCHF', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/瑞士法郎', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (54, 'IQO', 'AUD-JPY', 'qxbroker', 'AUDJPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/日元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (55, 'IQO', 'GBP-AUD', 'qxbroker', 'GBPAUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/澳大利亚元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (56, 'IQO', 'AUD-CHF', 'raceoption', 'AUD/CHF', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/瑞士法郎', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (57, 'IQO', 'AUD-JPY', 'raceoption', 'AUD/JPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/日元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (58, 'IQO', 'GBP-AUD', 'raceoption', 'GBP/AUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/澳大利亚元', 'system', 'system', '2022-05-20 15:20:54', '2022-06-16 14:54:04');
