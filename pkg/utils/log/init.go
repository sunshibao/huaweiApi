package log

import (
	"flag"
	"fmt"
	"path"
	"regexp"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	// 	some vendor will use glog as logger, which will create logfile under /tmp when error is logged.
	// 	this will cause the program exit, if the /tmp directory is not writable.
	// 	so we disable Glog, and prevent glog to create logfile
	logToStdErr := flag.Lookup("logtostderr")
	if logToStdErr != nil && logToStdErr.Value != nil {
		_ = logToStdErr.Value.Set("true")
	}

	// Set log formatter to ouput source code file name, line number and function name.
	var re = regexp.MustCompile(`^dev-gitlab.wanxingrowth.com`)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			fileName := path.Base(f.File)
			return fmt.Sprintf("%s()", re.ReplaceAllString(f.Function, "")), fmt.Sprintf("%s:%d", fileName, f.Line)
		},
	})
}
