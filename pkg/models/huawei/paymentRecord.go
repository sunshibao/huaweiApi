package huawei

type PaymentRecord struct {
	Id            uint64 `gorm:"column:id;primary_key;auto_increment:false;comment:'主键ID'"`
	PaymentID     string `gorm:"column:paymentID;type:varchar(128);default:'';comment:'支付唯一ID'"`
	Msisdn        string `gorm:"column:msisdn;type:varchar(128);default:'';comment:'手机号'"`
	ProductID     string `gorm:"column:productID;type:varchar(128);default:'';comment:'产品ID'"`
	ExtRef        string `gorm:"column:extRef;type:varchar(128);default:'';comment:'SP请求的流水号'"`
	Status        uint64 `gorm:"column:status;default:0;comment:'支付状态:1:PENDING:待确认 2:PAYMENTSUCCESS:支付成功 3:PAYMENTFAILED:支付失败'"`
	Amount        uint64 `gorm:"column:amount;default:0;comment:'费用(cents)'"`
	SubTime       string `gorm:"column:subTime;type:varchar(128);default:'';comment:'订购时间，格式:2018-10-17T13:05:48+02:00'"`
	StartTime     string `gorm:"column:startTime;type:varchar(128);default:'';comment:'合约生效时间，格式:2018-10- 17T13:05:48+02:00'"`
	EndTime       string `gorm:"column:endTime;type:varchar(128);default:'';comment:'合约失效时间，格式:2018-10- 17T13:05:48+02:00'"`
	SvcName       string `gorm:"column:svcName;type:varchar(128);default:'';comment:'业务名称'"`
	ChannelName   string `gorm:"column:channelName;type:varchar(128);default:'';comment:'订购渠道(WAP/USSD/SMS)'"`
	RenewalType   string `gorm:"column:renewalType;type:varchar(128);default:'';comment:'自动续订方式(AUTO)'"`
	BillingRate   uint64 `gorm:"column:billingRate;type:varchar(128);default:'';comment:'费用(cents)'"`
	BillingCycle  string `gorm:"column:billingCycle;type:varchar(128);default:'';comment:'周期 ONCE:按次 DAILY:包天 WEEKLY:包周 MONTHLY:包月'"`
	UpdatedAt     string `gorm:"column:updatedAt;type:varchar(128);default:'';comment:'最新更新时间。格式:2018-10- 17T13:05:48+02:00'"`
	LastBilledAt  string `gorm:"column:lastBilledAt;type:varchar(128);default:'';comment:'上一次扣费时间。格式:2018-10- 17T13:05:48+02:00'"`
	NextBillingAt string `gorm:"column:nextBillingAt;type:varchar(128);default:'';comment:'下一次扣费时间。格式:2018-10- 17T13:05:48+02:00'"`
}

func (PaymentRecord) TableName() string {
	return "payment_record"
}
