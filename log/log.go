package log

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/logfileGo.log"

func Log(errToLog error, prefix string) {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("utils.logger.Log(): ", err)
	} else {
		iLog := log.New(f, prefix, log.LstdFlags)
		iLog.Println(errToLog.Error())
		fmt.Println(errToLog.Error())
	}
}
