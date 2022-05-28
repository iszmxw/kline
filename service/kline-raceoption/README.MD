# 外汇交易时间

印尼时间： 周一早上4点开始，周六早上4点结束 北京时间： 周一早上5点开始，周六早上5点结束

# 黄金，白银交易时间

2022-05-28 03:59:55.296

印尼时间： 周一早上5点开始，周六早上5点结束 北京时间： 周一早上6点开始，周五早上6点结束

## 上线注意事项

1、更新数据表结构 -- 表：t_kline_log

新增字段 upstream_ts bigint(20)  上游K线图返回的ts

新增索引

upstream_ts NORMAL BTREE

-- 表：t_trading_assets

新增字段

point_number int(6)  保留小数点位数

assets_status int(6)  资产状态 1正常 2休市 3下架

sleep_time varchar(128)  休市时间 {"start":"6-05:00","ts_long":"172800"} start开始时间，从周几开始几点开始，休息多少秒






{"start":"6-04:00","ts_long":"172800"}
{"start":"5-05:00","ts_long":"259200"}
{"start":"5-05:00","ts_long":"259200"}
{"start":"6-04:00","ts_long":"172800"}
{"start":"6-04:00","ts_long":"172800"}
{"start":"6-04:00","ts_long":"172800"}


