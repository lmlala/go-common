// Date : 4/1/17 PM7:55
// Copyright: TradeShift.com
// Author = 'liming'
// Email = 'lim@cn.tradeshift.com'
package common

// Date : 18/10/16 PM11:22
// Copyright: TradeShift.com
// Author = 'liming'
// Email = 'lim@cn.tradeshift.com'

import (
	//"github.com/astaxie/beego/logs"
	"os"
	log "github.com/Sirupsen/logrus"
	"time"
	"fmt"
)

var logger *log.Logger
var loglevels = map[string]log.Level{
		"info":	log.InfoLevel,
		"debug": log.DebugLevel,
		"warning": log.WarnLevel,
		"error": log.ErrorLevel,
		"panic": log.PanicLevel,
		"fatal": log.FatalLevel,
	}

//func GetLogger_Beego() (logger *logs.BeeLogger){
//
//	format := `{"filename":"{{.}}","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`
//	log_format, err := StringFormat(format, logfile)
//	if err != nil{
//		fmt.Println("Log Format Error!!!!")
//	}
//	logger = logs.NewLogger()
//	logger.SetLogger(logs.AdapterFile, log_format)
//	// enable line no and file name
//	logger.EnableFuncCallDepth(true)
//	// async and async buffer
//	logger.Async(1e3)
//
//	return logger
//
//}

func GETLOGGER(logfile, loglevel string) *log.Logger {
	//log.WithFields(log.Fields{
	//"animal": "walrus",
	//"size":   10,
	//}).Info("A group of walrus emerges from the ocean")
	
	_, err := os.Stat(logfile)
	if err != nil {
		// try to create this file if not exist
		if os.IsNotExist(err) {
			_, err := os.Create(logfile)
			if err != nil {
				fmt.Println("Can not create log file: " + logfile)
			}
		} else {
			fmt.Println("Illegl Log File Content: " + logfile)
		}
	}

	if logger != nil{
		return logger
	}

	logger = log.New()
	f, _ := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	log.SetOutput(f)
	//logger.Out = f
	logger.Level = loglevels[loglevel]
	format := &log.TextFormatter{}
	format.ForceColors = true
	format.FullTimestamp = true
	format.TimestampFormat = time.RFC822
	logger.Formatter = format
	//defer f.Close()
	return logger
}