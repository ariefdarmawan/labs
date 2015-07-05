package main

import (
	"fmt"
	. "github.com/juragan360/hdfs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	h, e := NewHdfs(NewHdfsConfig("http://localhost:50070", ""))
	if e != nil {
		fmt.Println(e.Error())
	}
	h.Config.PoolSize = 5

	h.Delete("/user", true)
	h.MakeDirs([]string{"/user/ariefdarmawan/inbox/ecfz/json", "/user/ariefdarmawan/inbox/temp"}, "")

	dirname := "/Users/ariefdarmawan/Temp/ECFZ/TempInsurance/JSON"
	fmt.Println(">>>> TEST COPY DIR <<<<")
	/*
		e, es := h.PutDir("/Users/ariefdarmawan/Temp/ECFZ/TempInsurance/JSON", "/user/ariefdarmawan/inbox/ecfz/json")
		if es != nil {
			for k, v := range es {
				fmt.Println(fmt.Sprintf("Error when create %v : %v \n", k, v))
			}
		}
	*/
	dirs := ioutil.ReadDir(dirname)
	filenames := []string{}
	for _, dir := range dirs {
		if !strings.HasPrefix(dir.Name(), ".") {
			filenames = append(filenames, filepath.Join(dirname, dir.Name()))
		}
	}

	poolSize := 10
	workCount := len(filenames)

	fmt.Println(">>>> TEST GET STATUS <<<<")
	hdata, e := h.List("/user/ariefdarmawan/inbox/ecfz/json")
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("Data Processed :\n%v\n", len(hdata.FileStatuses.FileStatus))
	}

	fmt.Println("Done\n")
}
