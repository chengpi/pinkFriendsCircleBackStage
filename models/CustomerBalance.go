package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CustomerBalance struct {
	BalanceId          int64     `orm:"column(balance_id);pk;auto"           json:"balance_id"         description:"金额日志ID号"`
	CustomerId         string    `orm:"column(customer_id);size(100)"        json:"customer_id"        description:"用户抖音ID号"`
	Source             int64     `orm:"column(source);"                      json:"source"             description:"记录来源（1他人打赏，0提现）"`
	SourceId           int64     `orm:"column(source_id)"                    json:"source_id"          description:"相关交易单据ID号"`
	CreateTime         time.Time `orm:"column(create_time)"                  json:"create_time"        description:"记录生成时间"`
	amount             float64   `orm:"column(amount)"                       json:"amount"             description:"变动金额"`
	RewardCustomerId   string    `orm:"column(reward_customer_id);size(100)" json:"reward_customer_id" description:"打赏人抖音ID号（记录来源为1时不为空）"`
	VideoWorksId       string    `orm:"column(video_works_id);size(100)"     json:"video_works_id"     description:"视频作品ID号（用户抖音ID号+序号）"`
	OpenId             string    `orm:"column(open_id);size(100)"            json:"open_id"            description:"用户在当前小程序的 ID号（即唯一标识ID）"`
}

func init() {
	orm.RegisterModel(new(CustomerBalance))
}
func (c *CustomerBalance) TableName() string {
	return "customer_balance_log"
}

//func InsertOrUpdateCustomerBalance() error{
//	o := orm.NewOrm()
//}
