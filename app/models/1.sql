CREATE TABLE `t_trading_assets_monitoring` (
                                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '资产表主键ID',
                                               `item_code` varchar(128) DEFAULT NULL COMMENT '应用编码',
                                               `assets_code` varchar(128) DEFAULT NULL COMMENT '资产名称',
                                               `is_valid` int(1) DEFAULT NULL COMMENT '是否有效1-是 0-否',
                                               `area_timer` int(15) DEFAULT NULL COMMENT '设置区域时长',
                                               `interval_timer` int(15) DEFAULT NULL COMMENT '设置间隔触发时间',
                                               `inadvance_timer` int(15) DEFAULT NULL COMMENT '设置提前触发时间（拿这个时间间隔作为判断条件）',
                                               `loss_threshold` double(10,2) DEFAULT NULL COMMENT '触发属性：区域平台亏损金额>=xx美金',
                                               `shock_amplitude` double(10,2) DEFAULT NULL COMMENT '震荡幅度',
                                               `volatility_amplitude` double(10,2) DEFAULT NULL COMMENT '波动幅度',
                                               `volatility_frequency` varchar(255) DEFAULT NULL COMMENT '波动频率',
                                               `creator` varchar(128) DEFAULT NULL COMMENT '创建人',
                                               `updater` varchar(128) DEFAULT NULL COMMENT '更新人',
                                               `create_date` datetime DEFAULT NULL COMMENT '创建日期',
                                               `update_date` datetime DEFAULT NULL COMMENT '更新日期'
) ENGINE=InnoDB AUTO_INCREMENT=152 DEFAULT CHARSET=utf8mb4 COMMENT='资产表';