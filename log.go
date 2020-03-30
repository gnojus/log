package log

import (
	"fmt"
	// "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	// "time"
)

//var log = logrus.New()
var file *os.File

func init() {
	c := exec.Command(`x-terminal-emulator`, `-e`, `bash -c "if [ ! -e ./mylog ]; then mkfifo mylog; trap \"rm ./mylog\" EXIT; while true; do cat mylog && echo && echo \"---\" && echo; done; fi"`)
	c.Start()
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
