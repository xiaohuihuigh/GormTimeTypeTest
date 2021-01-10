package GormTImeTypeTest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golib/assert"
	"testing"
)

func TestDemoModel(t *testing.T) {
	it := assert.New(t)
	demo1 := NewDemoModel()
	demo1.Name = "sdf"
	err := errors.New("")
	err = demo1.Insert()
	it.Empty(err)
	demo1.Name  = "dff"
	err = demo1.Update(6)
	it.Empty(err)
	s,err := json.Marshal(demo1)
	fmt.Print(string(s))
}
