# 外汇交易时间

印尼时间： 周一早上4点开始，周六早上4点结束 北京时间： 周一早上5点开始，周六早上5点结束

# 黄金，白银交易时间

印尼时间： 周一早上5点开始，周五早上5点结束 北京时间： 周一早上6点开始，周五早上6点结束

# https://olymptrade.com/platform

该平台需要登录才可以socket握手，登录需要session，每次都需要去官网登录获取一下session,一个session的有效期 为10天，请注意保持session更新

### 增加自动登录脚本，登录过期自动重新登录

## plus 链接 

wss://ws-plus.olymptrade.com/connect

```json
[{"e":1000,"t":2,"d":{"pairs":["XAUUSD"]},"uuid":"2"}]
```