package GormTImeTypeTest

import (
	"time"
)
//gorm的时间类型很还是很强大的所以依旧使用time.Time类型
//因为mysql中的update要有某一项发生改变才会生效，所以添加实际数据Name，否则无法进行测试
//但是在进行json转换时没有办法正常显示
type DemoTimeTime struct {
	ID int64 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	DT time.Time `gorm:"column:dt;default:null" json:"dt"`
	CT time.Time `gorm:"column:ct;default:null" json:"ct"`
}
type DemoString struct {
	ID  int64 `gorm:"column:id"`
	DT string `gorm:"column:dt"`
	TS string `gorm:"column:ts"`
}
type DemoMyTime struct {
	ID int64 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	DT MyTime `gorm:"column:dt;default:null" json:"dt"`
	CT MyTime `gorm:"column:ct;default:null" json:"ct"`
}
