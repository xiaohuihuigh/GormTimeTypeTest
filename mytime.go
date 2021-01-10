package GormTImeTypeTest

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05" //常规类型
)

type MyTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t MyTime) MarshalJSON()([]byte,error){
	formatted:= fmt.Sprintf("\"%s\"",t.Format(YYYYMMDDHHMMSS))
	return []byte(formatted),nil
}
// Value insert timestamp into mysql need this function.
func (t MyTime) Value()(driver.Value,error){
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano(){
		return nil,nil
	}
	return t.Time,nil
}
//Scan valueof time.Time
func (t *MyTime)Scan(v interface{})error{
	value,ok := v.(time.Time)
	if ok {
		*t = MyTime{Time:value}
		return nil
	}
	return fmt.Errorf("can not conver %v to timestamp",v)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *MyTime) UnmarshalJSON(data []byte) error{
	// Ignore null, like in the main JSON package
	if string(data) == "null"{
		return nil
	}
	var err error
	// Fractional seconds are handled implicitly by Parse
	(*t).Time,err = time.ParseInLocation(`"`+YYYYMMDDHHMMSS+`"`,string(data),time.Local)
	return err
}