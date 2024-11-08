package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Init() {
	// 添加日志切割输出
	var hook = NewLfsHook("logs", time.Hour*12, 6)
	logrus.AddHook(hook)
	// 忽略控制台打印
	logrus.SetOutput(io.Discard)
	// 展示日志行数
	logrus.SetReportCaller(true)
}

// NewLfsHook 日志钩子(日志拦截，并重定向)
func NewLfsHook(logName string, rotationTime time.Duration, leastDay uint) logrus.Hook {
	_ = os.Mkdir(logName, os.ModeDir)
	// 可设置按不同level创建不同的文件名，咱们把6中日志都写到同一个writer中
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: NewWriter(filepath.Join(logName, "debug"), rotationTime, leastDay),
		logrus.InfoLevel:  NewWriter(filepath.Join(logName, "info"), rotationTime, leastDay),
		logrus.WarnLevel:  NewWriter(filepath.Join(logName, "warn"), rotationTime, leastDay),
		logrus.ErrorLevel: NewWriter(filepath.Join(logName, "error"), rotationTime, leastDay),
		logrus.FatalLevel: NewWriter(filepath.Join(logName, "fatal"), rotationTime, leastDay),
		logrus.PanicLevel: NewWriter(filepath.Join(logName, "panic"), rotationTime, leastDay),
	}, nil) //&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}
	return lfsHook
}

func NewWriter(logName string, rotationTime time.Duration, leastDay uint) io.Writer {
	writer, err := rotatelogs.New(
		// 1 日志文件名字
		logName+".%Y-%m-%d_%H_%M", // _%S
		// 2 日志周期(默认每86400秒/一天旋转一次)
		rotatelogs.WithRotationTime(rotationTime),
		// 3 清除历史 (WithMaxAge和WithRotationCount只能选其一)
		//rotatelogs.WithMaxAge(time.Hour*24*7), //默认每7天清除下日志文件
		rotatelogs.WithRotationCount(leastDay), //只保留最近的N个日志文件
	)
	if err != nil {
		log.Panic(err)
	}
	return writer
}
