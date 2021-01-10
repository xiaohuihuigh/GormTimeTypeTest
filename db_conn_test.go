package GormTImeTypeTest

import (
	"fmt"
	"github.com/golib/assert"

	"testing"
)

func TestMain(m *testing.M) {
	err := Setup()
	fmt.Print(err)
	it := assert.Assertions{}
	it.Empty(err)
	m.Run()
}

