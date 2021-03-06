package em_library

import (
	Package_Logrus "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Instance_Logrus *Package_Logrus.Logger

func Init_Logrus(logLevel string) {
	Instance_Logrus = Package_Logrus.New()
	// Instance_Logrus as JSON instead of the default ASCII formatter.
	Instance_Logrus.Formatter = new(Package_Logrus.JSONFormatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Instance_Logrus.Out = &lumberjack.Logger{
		Filename:   "storage/log/app.log",
		MaxSize:    500, // megabytes
		MaxAge:     30, //days
		Compress:   true, // disabled by default
	}

	// Only log the warning severity or above.
	level, err := Package_Logrus.ParseLevel(logLevel)
	if err != nil {
		level = Package_Logrus.WarnLevel
		initLog.Fatalln("[WARNING]", "Set Instance_Logrus Level Failed!", " Error:", err)
	} else {
		initLog.Println("[INFO]", "logrus initialized successfully.")
	}
	Instance_Logrus.Level = level
}

type logrus struct {

}

func NewLogrus() *logrus {
	return &logrus{}
}

func (this *logrus) Panic(args ...interface{}) {
	Instance_Logrus.Panic(args)
	return
}

func (this *logrus) Fatal(args ...interface{}) {
	Instance_Logrus.Fatal(args)
	return
}

func (this *logrus) Error(args ...interface{}) {
	Instance_Logrus.Error(args)
	return
}

func (this *logrus) Warning(args ...interface{}) {
	Instance_Logrus.Warning(args)
	return
}

func (this *logrus) Info(args ...interface{}) {
	Instance_Logrus.Info(args)
	return
}

func (this *logrus) Debug(args ...interface{}) {
	Instance_Logrus.Debug(args)
	return
}

func (this *logrus) Trace(args ...interface{}) {
	Instance_Logrus.Trace(args)
	return
}

