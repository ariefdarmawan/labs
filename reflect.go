package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

type Controller struct {
}

type IController interface {
	IsController() bool
	GetTitle() string
}

func (c *Controller) Do1() string {
	return "Do1"
}

func (c *Controller) Do2() string {
	return "Do2"
}

func (c *Controller) GetTitle() string {
	return "Controller0"
}

func (c *Controller) IsController() bool {
	return true
}

type Controller2 struct {
	Controller
}

func (c *Controller2) GetTitle() string {
	return "Controller2"
}

func parseController(m interface{}) error {
	r1 := reflect.TypeOf(m)

	//--- check if it is a controller
	if r1.Implements(reflect.TypeOf((*IController)(nil)).Elem()) == false {
		return errors.New("Not implementing IController")
	}

	v1 := reflect.ValueOf(m)
	//rElements := r1.Elem()
	//vElements := v1.Elem()

	vs := v1.MethodByName("GetTitle").Call(nil)
	name := vs[0].String()
	fmt.Println(name + " implements IController")

	numOfMethod := r1.NumMethod()
	for iMethod := 0; iMethod < numOfMethod; iMethod++ {
		method := r1.Method(iMethod)
		fmt.Println(method.Name)
		vs = v1.Method(iMethod).Call(nil)
		fmt.Printf("Call %s = %v \n", method.Name, vs[0])
	}
	fmt.Println("... Done ...")
	fmt.Println("")

	return nil
}

func main() {
	c := new(Controller)
	e := parseController(c)
	if e != nil {
		fmt.Println("Invalid :" + e.Error())
	}

	c2 := new(Controller2)
	e = parseController(c2)
	if e != nil {
		fmt.Println("Invalid :" + e.Error())
	}
}
