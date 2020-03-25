package log

import (
	"fmt"
	// "github.com/sirupsen/logrus"
	"os"
	// "time"
)

//var log = logrus.New()
var file *os.File

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example

	// Only log the warning severity or above.
	//var err error
	file, _ = os.OpenFile("mylog", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	log.Out = file
	//} else {
	//	logrus.Info("Failed to log to file, using default stderr")
	//}
	//log.SetLevel(logrus.DebugLevel)
}

func Log(args ...interface{}) {
	//log.Debug(args...)
	file.WriteString(fmt.Sprintln(args...))
}

func LogVar(v interface{}) {
	file.WriteString(fmt.Sprintf("%#v\n", v))
}

func Print(val interface{}) {
	file.WriteString(fmt.Sprint(val))
}
