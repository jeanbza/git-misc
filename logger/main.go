package woodsman

import (
    "fmt"
    "flag"
    "os"

    "github.com/golang/glog"
    "github.com/blackjack/syslog"
)

var logType = flag.String("log_type", "stderr", "-log_type [stderr|syslog]")
var logName = flag.String("log_name", "awesome_app", "-log_name <some name>")

func init() {
    flag.Parse()
    
    switch {
        case *logType == "stderr":
            defer glog.Flush()
        case *logType == "syslog":
            syslog.Openlog(*logName, syslog.LOG_PID, syslog.LOG_USER)
            defer syslog.Closelog()
        default:
            fmt.Println("Incorrect flag passed! Run again with -h to see flag usage")
            os.Exit(1)
    }
}

func Info(msg string) {
    switch {
        case *logType == "stderr":
            glog.Infoln(msg)
        case *logType == "syslog":
            syslog.Info(msg)
    }
}

func Error(msg string) {
    switch {
        case *logType == "stderr":
            glog.Errorln(msg)
        case *logType == "syslog":
            syslog.Err(msg)
    }
}