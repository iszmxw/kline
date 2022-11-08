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


-- iqoption

INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (59, 'IQO', 'EUR-USD', 'iqoption', 'EUR/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/美元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (60, 'IQO', 'GBP-USD', 'iqoption', 'GBP/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/美元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:28');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (61, 'IQO', 'AUD-CAD', 'iqoption', 'AUD/CAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/加拿大元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (62, 'IQO', 'USD-JPY', 'iqoption', 'USD/JPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/日元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (63, 'IQO', 'GOLD', 'iqoption', 'XAU/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '伦敦金', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-03 04:07:22');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (64, 'IQO', 'EUR-AUD', 'iqoption', 'EUR/AUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/澳大利亚元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (65, 'IQO', 'USD-CAD', 'iqoption', 'USD/CAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/加拿大元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (66, 'IQO', 'AUD-CHF', 'iqoption', 'AUD/CHF', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/瑞士法郎', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-01 04:03:58');

-- Fxcm

INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (67, 'IQO', 'EUR-USD', 'fxcm', 'EUR/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/美元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (68, 'IQO', 'GBP-USD', 'fxcm', 'GBP/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/美元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:28');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (69, 'IQO', 'AUD-CAD', 'fxcm', 'AUD/CAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/加拿大元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (70, 'IQO', 'USD-JPY', 'fxcm', 'USD/JPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/日元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (71, 'IQO', 'GOLD', 'fxcm', 'XAU/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '伦敦金', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-03 04:07:22');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (72, 'IQO', 'SILVER', 'fxcm', 'XAG/USD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '伦敦银', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-03 04:01:22');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (73, 'IQO', 'CAD-JPY', 'fxcm', 'CAD/JPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '加拿大元/日元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (74, 'IQO', 'EUR-AUD', 'fxcm', 'EUR/AUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/澳大利亚元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (75, 'IQO', 'USD-CAD', 'fxcm', 'USD/CAD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/加拿大元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (76, 'IQO', 'AUD-CHF', 'fxcm', 'AUD/CHF', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/瑞士法郎', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-01 04:03:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (77, 'IQO', 'AUD-JPY', 'fxcm', 'AUD/JPY', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '澳大利亚元/日元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-10-31 19:16:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (78, 'IQO', 'GBP-AUD', 'fxcm', 'GBP/AUD', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '英镑/澳大利亚元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-01 04:10:58');


-- 添加OTC数据源

INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (79, 'IQO', 'EUR-USD-OTC', 'qxbroker', 'EURUSD_otc', NULL, 1, '欧元/美元', 'system', '叶志明', '2022-05-20 15:20:54', '2022-11-03 13:56:08');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (80, 'IQO', 'GBP-USD-OTC', 'qxbroker', 'GBPUSD_otc', NULL, 0, '英镑/美元', 'system', '叶志明', '2022-05-20 15:20:54', '2022-11-03 13:56:05');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (81, 'IQO', 'AUD-CAD-OTC', 'qxbroker', 'AUDCAD_otc', NULL, 0, '澳大利亚元/加拿大元', 'system', '叶志明', '2022-05-20 15:20:54', '2022-11-08 09:34:44');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (82, 'IQO', 'USD-JPY-OTC', 'qxbroker', 'USDJPY_otc', NULL, 1, '美元/日元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-05 04:14:08');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (83, 'IQO', 'GOLD-OTC', 'qxbroker', 'XAUUSD_otc', NULL, 1, '伦敦金', 'system', '叶志明', '2022-05-20 15:20:54', '2022-11-03 13:56:00');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (84, 'IQO', 'SILVER-OTC', 'qxbroker', 'XAGUSD_otc', NULL, 1, '伦敦银', 'system', '叶志明', '2022-05-20 15:20:54', '2022-11-03 13:55:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (85, 'IQO', 'Microsoft-OTC', 'qxbroker', 'MSFT_otc', NULL, 1, '微软', 'system', '叶志明', '2022-05-20 15:20:54', '2022-11-08 10:08:48');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (86, 'IQO', 'Intel-OTC', 'qxbroker', 'INTC_otc', NULL, 1, '英特尔', 'system', 'tdw', '2022-05-20 15:20:54', '2022-05-27 17:31:58');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (87, 'IQO', 'CAD-JPY-OTC', 'qxbroker', 'CADJPY_otc', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '加拿大元/日元', 'system', 'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (88, 'IQO', 'EUR-AUD-OTC', 'qxbroker', 'EURAUD_otc', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '欧元/澳大利亚元', 'system', '叶志明', '2022-05-20 15:20:54', '2022-11-04 17:20:02');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (89, 'IQO', 'USD-CAD-OTC', 'qxbroker', 'USDCAD_otc', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 0, '美元/加拿大元', 'system', 'Lucy', '2022-05-20 15:20:54', '2022-05-28 09:10:30');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (90, 'IQO', 'AUD-CHF-OTC', 'qxbroker', 'AUDCHF_otc', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '澳大利亚元/瑞士法郎', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-05 23:59:38');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (91, 'IQO', 'AUD-JPY-OTC', 'qxbroker', 'AUDJPY_otc', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '澳大利亚元/日元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-05 12:07:38');
INSERT INTO `t_trading_assets_source`(`id`, `item_code`, `assets_code`, `kline_source`, `subscribe_symbol`, `sleep_time`, `is_valid`, `remark`, `creator`, `updater`, `create_date`, `update_date`) VALUES (92, 'IQO', 'GBP-AUD-OTC', 'qxbroker', 'GBPAUD_otc', '{\"start\":\"6-04:00\",\"ts_long\":\"172800\"}', 1, '英镑/澳大利亚元', 'system', '系统自动切换', '2022-05-20 15:20:54', '2022-11-05 23:59:38');



-- 添加OTC资产

INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'EUR/USD (OTC)', 'EUR-USD-OTC', 1, 5, 1, 1, 'https://fileupload.ssmartoption.com/file/download/47ca26e0-4693-43aa-ae44-4c64dcb602bc.png', 9, '2022-08-10 11:06:16', '2022-09-29 13:29:25', 'zz', '嘉鑫', 0.70, NULL, 4, 'exchange', 'Forex', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'GBP/USD (OTC)', 'GBP-USD-OTC', 1, 5, 1, 1, 'https://fileupload.ssmartoption.com/file/download/2fc3be0d-47a9-47c8-8761-3dfc4f0a88e2.png', 8, '2022-08-10 11:06:16', '2022-11-07 15:32:20', 'zz', '嘉鑫', 0.70, NULL, 4, 'exchange', 'Forex', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'AUD/CAD (OTC)', 'AUD-CAD-OTC', 1, 5, 1, 1, 'https://fileupload.ssmartoption.com/file/download/d52bfe77-4c23-4457-9882-02813bd0ca12.png', 7, '2022-08-10 11:06:16', '2022-10-29 13:06:36', 'zz', '嘉鑫', 0.65, NULL, 4, 'exchange', 'Forex', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'USD/JPY (OTC)', 'USD-JPY-OTC', 1, 4, 1, 1, 'https://fileupload.ssmartoption.com/file/download/9bacf651-dd7b-4f1d-82fd-3bb553891cbb.png', 6, '2022-08-10 11:06:16', '2022-09-22 17:19:25', 'zz', '嘉鑫', 0.60, NULL, 4, 'exchange', 'Forex', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'GOLD (OTC)', 'GOLD-OTC', 1, 3, 1, 1, 'https://fileupload.ssmartoption.com/file/download/69b987a7-380d-46ea-b618-73a5260c1655.png', 44, '2022-08-10 11:06:16', '2022-09-22 09:57:16', 'zz', '嘉鑫', 0.60, NULL, 6, 'commodity', 'Commodities', 'India', 0);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'SILVER (OTC)', 'SILVER-OTC', 1, 3, 1, 1, 'https://fileupload.ssmartoption.com/file/download/6a8b4562-6a98-49ce-9ef1-6cd1e9ae9fc2.png', 43, '2022-08-10 11:06:16', '2022-09-23 08:31:31', 'zz', '嘉鑫', 0.45, NULL, 6, 'commodity', 'Commodities', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'Microsoft (OTC)', 'Microsoft-OTC', 1, 3, 1, 1, 'https://fileupload.ssmartoption.com/file/download/dc2b9e40-b722-4a5e-ae6e-81b7a48f0f23.png', 2, '2022-08-10 11:06:16', '2022-11-08 11:13:29', 'zz', 'kline-system', 0.75, NULL, 3, 'stock', 'Stocks', 'India', 0);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'Intel (OTC)', 'Intel-OTC', 1, 3, 1, 1, 'https://fileupload.ssmartoption.com/file/download/533ab0cc-5c96-4d6b-962a-3c25435cd3ad.png', 3, '2022-08-10 11:06:16', '2022-11-07 15:32:41', 'zz', '嘉鑫', 0.70, NULL, 3, 'stock', 'Stocks', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'CAD/JPY (OTC)', 'CAD-JPY-OTC', 1, 3, 1, 1, 'https://fileupload.ssmartoption.com/file/download/984c31ea-af40-4597-840d-e4809d195f9c.png', 6, '2022-08-10 11:06:16', '2022-11-07 15:32:29', 'zz', '嘉鑫', 0.70, NULL, 4, 'exchange', 'Forex', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'EUR/AUD (OTC)', 'EUR-AUD-OTC', 1, 6, 1, 1, 'https://fileupload.ssmartoption.com/file/download/4cb570ae-0fd6-489f-a3e3-d3997b5bcba3.png', 4, '2022-08-10 11:06:16', '2022-10-21 10:00:19', 'zz', 'kline-system', 0.70, NULL, 4, 'exchange', 'Forex', 'India', 0);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'USD/CAD (OTC)', 'USD-CAD-OTC', 1, 6, 1, 1, 'https://fileupload.ssmartoption.com/file/download/ccaf9505-18d9-4b7a-9f80-b382b19d874f.png', 5, '2022-08-10 11:06:16', '2022-11-07 15:32:35', 'zz', '嘉鑫', 0.70, NULL, 4, 'exchange', 'Forex', 'India', 1);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'AUD/CHF (OTC)', 'AUD-CHF-OTC', 1, 5, 1, 1, 'https://fileupload.ssmartoption.com/file/download/13c117ac-5773-45b0-b0d0-2f8f964ec6c6.png', 0, '2022-08-10 11:06:16', '2022-11-08 11:13:29', 'zz', 'kline-system', 0.75, NULL, 4, 'exchange', 'Forex', 'India', 0);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'AUD/JPY (OTC)', 'AUD-JPY-OTC', 1, 3, 1, 1, 'https://fileupload.ssmartoption.com/file/download/c60a5aa8-ad9a-431c-8ce2-7c11ff51214e.png', 0, '2022-08-10 11:06:16', '2022-11-08 11:13:29', 'zz', 'kline-system', 0.75, NULL, 4, 'exchange', 'Forex', 'India', 0);
INSERT INTO `t_trading_assets`( `item_code`, `assets_name`, `assets_code`, `assets_status`, `point_number`, `is_valid`, `is_show`, `assets_logo_url`, `sort`, `create_date`, `update_date`, `creator`, `updater`, `return_rate`, `sleep_time`, `assets_type_id`, `assets_type_code`, `assets_type_name`, `nation`, `is_whitelist`) VALUES ('SMO', 'GBP/AUD (OTC)', 'GBP-AUD-OTC', 1, 5, 1, 1, 'https://fileupload.ssmartoption.com/file/download/9283e82d-41bb-416b-bdec-aa4c500b4b5d.png', 0, '2022-08-10 11:06:16', '2022-09-29 15:54:52', 'zz', '嘉鑫', 0.70, NULL, 4, 'exchange', 'Forex', 'India', 1);

