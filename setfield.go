package main

import (
	"fmt"
	"github.com/eaciit/toolkit"
)

type S struct {
	Id    string
	Title string
	Value int
}

func main() {
	o := new(S)
	//toolkit.SetField(o, "Id", "Test")
	//toolkit.SetField(o, "Value", 90000)
	fmt.Printf("%v", o)
}
