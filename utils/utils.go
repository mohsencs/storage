package utils

import (
	"log"
	"runtime"
)

func PrintError(err error) {
	pc, filename, line, _ := runtime.Caller(1)
	log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
}
