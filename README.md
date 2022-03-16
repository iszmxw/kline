# GoApiService

## 引入的库

> [gin-gonic/gin](https://github.com/gin-gonic/gin)   【Gin框架】
>
> [go-playground/validator](https://github.com/go-playground/validator)   【validator表单验证器】
>
> [concurrent-map](https://github.com/orcaman/concurrent-map)   【concurrent-map替换原生map解决并发读写】
>
> [gorm.io/gorm](https://gorm.io/gorm)   【Gorm数据查询工具】
>
> [go-redis/redis](https://github.com/go-redis/redis)   【go-redis缓存】
>
> [spf13/viper](https://github.com/spf13/viper)   【配置读取工具】
>
> [uber-go/zap](https://github.com/uber-go/zap)   【zap日志库】
>
> [gostudys/huobi_futures_go](https://github.com/gostudys/huobi_futures_go)   【火币网封装库】
>
> 日志TrackId链路追踪，包含mysql语句

## 打包编译

- 打包火币网socket服务

```shell
GOOS=linux GOARCH=amd64 go build -o huobiService service/huobi/socket.go
```

- 打包api接口

```shell
GOOS=linux GOARCH=amd64 go build -o apiService main.go
```

## 接口文档

### 1、k线图

#### 1.1、K线图历史行情

- 请求地址

```url
/v1/api/kline/history
```

- 请求方式

  `GET`

- 请求参数

| 字段名   | 类型   | 说明   |
| -------- | ------ | ------ |
| symbol | string | 例如：btc-usd，不传默认默认为"btc-usd" |
| period | string | 例如：1min，不传默认默认为"1min"   |
| limit | string | 例如：100,不传默认取出300条   |

```text
/v1/api/kline/history?symbol=btc-usd&period=1min
```

- 返回参数

| 字段名     | 类型     | 说明      |
| ---------- | -------- | --------- |
| code       | string      | 错误代码，1-正常，0-异常  |
| data       | object |     数据对象      |
|  - data.HistKlineLogs       | array |     历史行情      |
|  - data.beginDate       | string |     2022-03-15 14:00:27      |
|  - data.endDate       | string |     2022-03-15 14:01:27      |
|  - data.klineSource       | string |     来源：huobi      |
|  - data.totals       | int64 |     总条数：63      |
|  - data.contractCode       | string |     合约代码：ETH-USD      |
| reqId       | string |     请求Id     |

```json
{
  "code": 1,
  "data": {
    "HistKlineLogs": [
      {
        "amount": 0,
        "ch": "market.ETH-USD.index.1min",
        "contractCode": "ETH-USD",
        "close": 2543.271603555,
        "count": 0,
        "open": 2543.2662873644,
        "high": 0,
        "id": 1647327660,
        "klineSource": "huobi",
        "low": 2543.1410021688,
        "mrid": 0,
        "ts": 1647327687259,
        "vol": 0
      },
      {
        "amount": 0,
        "ch": "market.ETH-USD.index.1min",
        "contractCode": "ETH-USD",
        "close": 2543.271603555,
        "count": 0,
        "open": 2543.2662873644,
        "high": 0,
        "id": 1647327660,
        "klineSource": "huobi",
        "low": 2543.1410021688,
        "mrid": 0,
        "ts": 1647327686259,
        "vol": 0
      }
    ],
    "beginDate": "2022-03-15 14:00:27",
    "endDate": "2022-03-15 14:01:27",
    "klineSource": "huobi",
    "totals": 63,
    "contractCode": "ETH-USD"
  },
  "message": "success"
}
```

#### 1.2、K线图socket最新数据

- 请求地址

```url
/v1/api/kline/ws_index
```

- 请求方式

  `socket长连接`

- 请求参数

| 字段名   | 类型   | 说明   |
| -------- | ------ | ------ |
| body | string | market.BTC-USD.index.5min |

```text
market.BTC-USD.index.1min
```

- 返回参数

| 字段名     | 类型     | 说明      |
| ---------- | -------- | --------- |
| ch       | string      | "market.BTC-USD.index.5min"  |
| ts       | long |     响应生成时间点，单位：毫秒      |
| tick       | object array     |     tick返回，详情：推送tick参数      |
| - tick.id       | int64 |     指数K线id,也就是K线时间戳，K线起始时间          |
| - tick.vol       | string |     成交量张数为0      |
| - tick.count       | string |     成交笔数为0      |
| - tick.open       | string |     开盘指数价      |
| - tick.close       | string |     收盘指数价      |
| - tick.low       | string |     最低指数价      |
| - tick.high       | string |     最高指数价      |
| - tick.amount       | string |     数值为0      |

```json
{
  "ch": "market.BTC-USD.index.5min",
  "ts": 1647414930257,
  "tick": {
    "id": 1647414900,
    "vol": "0",
    "count": 0,
    "open": "39532.0925",
    "close": "39542.285",
    "low": "39530.6125",
    "high": "39542.285",
    "amount": "0"
  }
}
```

---

