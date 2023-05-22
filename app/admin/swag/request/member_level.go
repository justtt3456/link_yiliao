package request

type MemberLevelUpdate struct {
	Id   int    `gorm:"column:id;primary_key"` //
	Name string `gorm:"column:name"`           //等级名称
	Img  string `gorm:"column:img"`            //图标
}
