package model

type OrderLog struct {
	Id         int   `gorm:"column:id;primary_key"`             //关联用户id
	UId        int   `gorm:"column:uid"`                        //关联用户id
	OrderId    int   `gorm:"column:order_id"`                   //关联订单id
	Type       int   `gorm:"column:type"`                       //类型，1下单2结单
	Before     int64 `gorm:"column:before"`                     //变动前金额
	Amount     int64 `gorm:"column:amount"`                     //消费金额
	After      int64 `gorm:"column:after"`                      //变动后的金额
	CreateTime int64 `gorm:"column:create_time;autoCreateTime"` //结算时间
}

// TableName sets the insert table name for this struct type
func (o *OrderLog) TableName() string {
	return "c_order_log"
}
