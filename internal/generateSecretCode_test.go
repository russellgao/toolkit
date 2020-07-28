package internal

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRandStringBytesMaskImpr(t *testing.T) {
	convey.Convey("测试生成随机数", t, func() {
		n := 8
		result := RandStringBytesMaskImpr(n)
		convey.So(len(result), convey.ShouldEqual, 8)
	})
}
