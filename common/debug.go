package common

import "log"

var VerboseLogging = false

func Debugln(v ...interface{}) {
	if VerboseLogging {
		log.Println(v...)
	}
}
