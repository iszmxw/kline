package models

import "time"

type DailyActiveData struct {
	Id                          int64     `gorm:"column:id; PRIMARY_KEY" json:"id"` //日活跃数据统计表
	NumberInstallers            int       `json:"number_installers"`                //安装人数
	Enrolment                   int       `json:"enrolment"`                        //注册人数
	RegistrationNumber          int       `json:"registration_number"`              //注册次数
	NewlyRegisteredRecharged    int       `json:"newly_registered_recharged"`       //新注册且充值人数
	NewAccountRechargeRate      float64   `json:"new_account_recharge_rate"`        //新户充值率
	AverageAmountInitialCharge  float64   `json:"average_amount_initial_charge"`    //新户首充平均金额
	RechargeNumber              int       `json:"recharge_number"`                  //充值人数
	RechargeTimes               int       `json:"recharge_times"`                   //充值次数
	TotalRechargeAmount         float64   `json:"total_recharge_amount"`            //充值总金额
	RechargeServiceFee          float64   `json:"recharge_service_fee"`             //充值服务费
	PerCapitaRechargeAmount     float64   `json:"per_capita_recharge_amount"`       //人均充值金额
	NumberWithdrawals           int       `json:"number_withdrawals"`               //提币人数
	WithdrawalsTimes            int       `json:"withdrawals_times"`                //提币次数
	TotalWithdrawalAmount       float64   `json:"total_withdrawal_amount"`          //提币总金额
	WithdrawalCommission        float64   `json:"withdrawal_commission"`            //提币手续费
	WithdrawalServiceFee        float64   `json:"withdrawal_service_fee"`           //提币服务费
	AmountBonusIssued           float64   `json:"amount_bonus_issued"`              //发放bonus金额
	BonusTransactions           int       `json:"bonus_transactions"`               //bonus交易人数
	BonusTransactionsAmount     float64   `json:"bonus_transactions_amount"`        //bonus交易金额
	BonusUserProfitability      float64   `json:"bonus_user_profitability"`         //bonus用户盈利
	NumberTransactions          int       `json:"number_transactions"`              //交易人数
	NumberTimeTransactions      int       `json:"number_time_transactions"`         //交易次数
	TotalTransactionAmount      float64   `json:"total_transaction_amount"`         //交易总金额
	UserLoss                    float64   `json:"user_loss"`                        //用户亏损
	NumberTransactionsPerCapita float64   `json:"number_transactions_per_capita"`   //人均交易次数
	PerCapitaTransactionAmount  float64   `json:"per_capita_transaction_amount"`    //人均交易金额
	Odds                        float64   `json:"odds"`                             //赔率
	ChargeWithdrawalDifference  float64   `json:"charge_withdrawal_difference"`     //充提差额
	PlatformProfitLoss          float64   `json:"platform_profit_loss"`             //平台盈亏
	CreateDate                  time.Time `json:"create_date"`                      //创建日期
	UpdateDate                  time.Time `json:"update_date"`                      //更新日期
}
