package logger

import (
	"fmt"
	"runtime/debug"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	Init()
	logrus.Infoln("asdadsasd")
	time.Sleep(5 * time.Second)
	logrus.Infoln(string(debug.Stack()))
	logrus.WithFields(logrus.Fields{"RoleUid": "12313123123"}).Warn("asdasd")
}

func TestLogger2(t *testing.T) {
	str := "goroutine 68 [running]:\\nruntime/debug.Stack()\\n\\tD:/work/environment/GO/src/runtime/debug/stack.go:24 +0x65\\ncolly_v2/pkg/zlib.Try.func1()\\n\\tE:/gopackage2/2023.10/colly_v2/pkg/zlib/try.go:12 "
	fmt.Println(len([]byte(str)))
}

func TestLogger3(t *testing.T) {
	fmt.Println(time.Now().Unix() - 259200)
}
