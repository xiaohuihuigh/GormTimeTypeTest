package GormTImeTypeTest

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//var Conn _Conn
var Conn *gorm.DB


// Setup inject model with driver conn
func Setup() (err error) {
	//dbpath := "/Users/admin/Todo.db"
	dbpath := "root:z@tcp(127.0.0.1:3306)/timeTest?charset=utf8&loc=Local&parseTime=true"
	Conn,err= gorm.Open("mysql", dbpath)
	Conn.DB().SetMaxOpenConns(10)
	return Conn.DB().Ping()
}

//func SetOffsetAndLimit(conn *gorm.DB,userID int64,pageNum int){
//	if pageNum != 0 {
//		conn = conn.Offset((pageNum-1) * GetPageSize(userID)).Limit(GetPageSize(userID))
//	}else{
//		conn = conn.Limit(GetPageSize(userID))
//	}
//
//
