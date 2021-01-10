package GormTImeTypeTest

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DemoModel struct {
	*gorm.DB
	*DemoMyTime
}
 var DemoTable = "demo"

 func NewDemoModel() *DemoModel{
 	return &DemoModel{DemoMyTime: &DemoMyTime{},DB:Conn}
 }
func (u *DemoModel)SetTable() *gorm.DB  {
	conn := u.Table(DemoTable)
	return conn
}

func (u *DemoModel) FindByID(id int64) (err error) {
	err = u.SetTable().Where("id = ?", id).Limit(1).Find(&u.DemoMyTime).Error
	return
}
func (u *DemoModel) Insert()(err error) {
	err = u.SetTable().Debug().Create(&u.DemoMyTime).Error
	return
}
func (u *DemoModel) Update(id int64)(err error){
	old := NewDemoModel()
	err = old.FindByID(id)
	if err != nil{
		return
	}
	old.Name = u.Name
	old.DT.Time = time.Time{}
	old.CT.Time = time.Time{}
	err = u.SetTable().Debug().Where("id = ?",id).Update(&old.DemoMyTime).Error
	return
}
