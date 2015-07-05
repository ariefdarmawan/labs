package main

import (
	"fmt"
	"github.com/eaciit/toolkit"
)

func main() {
	var e error
	l, e := toolkit.NewLog(true, true, "/Users/ariefdarmawan/Temp", "log-test-%s.log", "20060102")
	if e != nil {
		fmt.Println("Error: ", e.Error())
	}

	l.Info("Test Log Info")
}
