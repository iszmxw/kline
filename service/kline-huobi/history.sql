CREATE TABLE `osx_hist_kline_logs`
(
    `auto_id`      bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自动递增id',
    `symbol`       varchar(20) NOT NULL COMMENT '交易代码,btcusdt, ethbtc...等',
    `period`       varchar(6)  NOT NULL COMMENT 'K线周期 1min, 5min, 15min, 30min, 60min, 4hour, 1day, 1mon, 1week, 1year',
    `ch`           varchar(60)     DEFAULT NULL COMMENT 'market.ETH-USD.index.1min',
    `ts`           bigint(13) DEFAULT NULL COMMENT '系统响应时间',
    `id`           bigint(20) NOT NULL COMMENT 'unix时间，同时作为K线ID',
    `amount`       decimal(25, 8)  DEFAULT NULL COMMENT '成交量',
    `count`        int(11) DEFAULT NULL COMMENT '成交笔数',
    `contractCode` varchar(255)    DEFAULT NULL COMMENT '合约代码',
    `open`         decimal(25, 8)  DEFAULT NULL COMMENT '开盘价',
    `close`        decimal(25, 8)  DEFAULT NULL COMMENT '收盘价（当K线为最晚的一根时，是最新成交价）',
    `low`          decimal(25, 8)  DEFAULT NULL COMMENT '最低价',
    `high`         decimal(25, 8)  DEFAULT NULL COMMENT '最高价',
    `vol`          decimal(25, 10) DEFAULT NULL COMMENT '成交额, 即 sum(每一笔成交价 * 该笔的成交量)',
    `klineSource`  varchar(255)    DEFAULT NULL COMMENT 'K线图来源',
    `mrid`         int(11) DEFAULT NULL COMMENT 'mrid',
    `created_at`   datetime    NOT NULL COMMENT '创建时间',
    `updated_at`   datetime    NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`   datetime        DEFAULT NULL COMMENT '删除时间，为 null 则是没删除',
    PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='K线图历史记录';


CREATE TABLE `t_kline_log_v2`
(
    `log_id`        bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'K线表主键ID',
    `contract_code` varchar(128) NOT NULL COMMENT '合约代码 BTC-USD ',
    `ch`            varchar(128)    DEFAULT NULL COMMENT 'K线类型请求kline的wss请求地址',
    `ts`            bigint(20) DEFAULT '0' COMMENT 'kilne返回日期时间戳',
    `id`            bigint(20) DEFAULT '0' COMMENT '返回的ID',
    `mrid`          decimal(20, 2)  DEFAULT '0.00' COMMENT 'ID',
    `c_open`        decimal(20, 10) DEFAULT '0.0000000000' COMMENT '起始峰值',
    `c_close`       decimal(20, 10) DEFAULT '0.0000000000' COMMENT '结束峰值',
    `high`          decimal(20, 10) DEFAULT '0.0000000000' COMMENT '最高峰值',
    `low`           decimal(20, 10) DEFAULT '0.0000000000' COMMENT '最低峰值',
    `amount`        decimal(20, 10) DEFAULT NULL COMMENT '成交金额',
    `vol`           decimal(20, 0)  DEFAULT NULL COMMENT '成交手',
    `c_count`       decimal(20, 0)  DEFAULT NULL COMMENT '交易笔数',
    `ts_date`       datetime        DEFAULT NULL COMMENT 'Kline返回日期',
    `create_date`   datetime        DEFAULT NULL COMMENT '创建日期',
    `kline_source`  varchar(64)     DEFAULT '' COMMENT 'k线来源 huobi',
    PRIMARY KEY (`log_id`),
    KEY             `contract_code` (`contract_code`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='K线数据表';