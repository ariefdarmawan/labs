package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ip, e := getIP()
	handleError(e)
	fmt.Println("IP of this computer is ", ip)
	fmt.Println("*** Bye ***")
}

func getIP() ([]string, error) {

	ret := make([]string, 0)
	he := func(err error) ([]string, error) {
		return ret, err
	}

	ifaces, err := net.Interfaces()
	he(err)

	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		he(err)

		for _, addr := range addrs {
			interfaceTxt := addr.String()
			if strings.HasSuffix(interfaceTxt, "24") {
				interfaceTxt = interfaceTxt[0 : len(interfaceTxt)-3]
				ret = append(ret, interfaceTxt)
			}
		}
	}
	if len(ret) == 0 {
		ret = append(ret, "127.0.0.1")
	}
	return ret, nil
}

func handleError(e error) {
	if e != nil {
		fmt.Println("ERROR: ", e.Error())
		os.Exit(200)
		return
	}
}
