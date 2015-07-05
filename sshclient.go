package main

import (
	"fmt"
	"github.com/juragan360/hdfs"
	//"github.com/juragan360/sshclient"
	"path/filepath"
)

func main() {
	hdDir := "/user/ariefdarmawan/inbox/temp/ssh"
	h, _ := hdfs.NewHdfs(hdfs.NewHdfsConfig("http://localhost:50070", "ariefdarmawan"))

	h.Delete(hdDir, true)
	e := h.MakeDir(hdDir, "755")
	if e != nil {
		fmt.Println("Error: ", e.Error())
		return
	}

	sourceFile := "/Users/ariefdarmawan/Temp/goroutine_1.txt"
	hdDest := filepath.Join(hdDir, "ssh.txt")

	//sc := sshclient.NewSshClient("localhost", "ariefdarmawan")
	//defer sc.Close()

	h.Put(sourceFile, hdDest, "755", nil)
	/*
		shDest := filepath.Join("/Users/ariefdarmawan/Temp", "sh.txt")
				if e = sc.Scp(sourceFile, shDest); e != nil {
					fmt.Println("Unable to copy file: ", e.Error())
					return
				}

				hdCmd := fmt.Sprintf("/usr/local/hadoop/bin/hdfs dfs -put %s %s", shDest, hdDest)
				sc.Run(hdCmd)
				if _, e = sc.Run(hdCmd); e != nil {
					fmt.Println("Unable to run : "+hdCmd+" error: ", e.Error())
					return
				}

			delFile := fmt.Sprintf("rm %s", shDest)
			if _, e = sc.Run(delFile); e != nil {
				fmt.Println("Unable to delete file: ", e.Error())
				return
			}
	*/

	fmt.Println("Done")
}
