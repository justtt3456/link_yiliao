package global

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func Log() {
	abs := CONFIG.Log.Abs()
	err, _ := rotatelogs.New(
		abs+"_%Y%m%d",
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithLinkName(abs),              // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	hook := NewHook(WriterMap{
		logrus.ErrorLevel: err,
		logrus.FatalLevel: err,
		logrus.PanicLevel: err,
	}, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.AddHook(hook)
	logrus.SetOutput(ioutil.Discard)

	logrus.Info("==========输出日志==============")
	logrus.Error("===========错误日志=============")
}
