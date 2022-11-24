package models

import "time"

// 资产表历史记录

type TradingAssetsCalculate struct {
	Id                int64     `json:"id" gorm:"column:id; PRIMARY_KEY"` //资产表主键ID
	ItemCode          string    `json:"item_code"`                        //应用编码
	AssetsName        string    `json:"assets_name"`                      //资产名称
	AssetsCode        string    `json:"assets_code"`                      //资产编码
	LineTimeStart     int64     `json:"line_time_start"`                  //k线开始时间
	LineTimeEnd       int64     `json:"line_time_end"`                    //k线结束时间
	DotCount          int       `json:"dot_count"`                        //该段时间线的点数
	InitialValue      float64   `json:"initial_value"`                    //初始值
	EndValue          float64   `json:"end_value"`                        //结束值
	Avg               float64   `json:"avg"`                              //平均值
	UpCount           int       `json:"up_count"`                         //上涨次数
	DownCount         int       `json:"down_count"`                       //下跌次数
	FlatCount         int       `json:"flat_count"`                       //平滑次数（前后两点相等次数）
	WaveRate          float64   `json:"wave_rate"`                        //波动率
	UpRate            float64   `json:"up_rate"`                          //上涨率
	DownRate          float64   `json:"down_rate"`                        //下跌率
	FlatRate          float64   `json:"flat_rate"`                        //平滑率
	Result            int8      `json:"result"`                           //最终波动结果 1-大 2-小
	CreateDate        time.Time `json:"create_date"`                      //创建日期
	UpdateDate        time.Time `json:"update_date"`                      //更新日期
	ReturnRateHistory float64   `json:"return_rate_history"`              //历史收益率
	ReturnRateNow     float64   `json:"return_rate_now"`                  //新的收益率
}
