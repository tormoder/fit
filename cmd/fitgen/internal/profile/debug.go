package profile

import (
	"log"
	"os"
	"strconv"
)

var debugfg, _ = strconv.ParseBool(os.Getenv("FITGEN_DEBUG"))

func debug(v ...interface{}) {
	if debugfg {
		log.Print(v...)
	}
}

func debugf(format string, v ...interface{}) {
	if debugfg {
		log.Printf(format, v...)
	}
}

func debugln(v ...interface{}) {
	if debugfg {
		log.Println(v...)
	}
}
