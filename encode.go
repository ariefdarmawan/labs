package main

import (
	"fmt"
	tk "github.com/eaciit/toolkit"
)

type obj struct {
	Output interface{}
}

func main() {
	s := new(obj)
	s.Output = "Arief Darmawan"
	sb, _ := tk.EncodeByte(s)
	fmt.Printf("Encoded: %v \n", s)
	result := new(obj)
	e := tk.DecodeByte(sb, &result)
	if e != nil {
		fmt.Println("Error decoding: ", e.Error())
	}
	fmt.Printf("Decode: %v \n", result.Output)
}
