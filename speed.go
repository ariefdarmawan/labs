package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {
	procno := flag.Int("proc", runtime.NumCPU(), "Number of processor to be used")
	flag.Parse()

	runtime.GOMAXPROCS(procno)
	fmt.Println("Running process using ", *procno, " processor(s)")
	jobcount := 5
	t0 := time.Now()
	for i := 0; i < jobcount; i++ {
		printFileAndSize()
	}
	t1 := time.Since(t0)
	fmt.Printf("No concurency completed in %v \n", t1)

	t0 = time.Now()
	done := make(chan bool)
	for i := 0; i < jobcount; i++ {
		go func(d chan bool) {
			printFileAndSize()
			d <- true
		}(done)
	}

	quit := false
	processed := 0
	for !quit {
		select {
		case <-done:
			processed = processed + 1
			if jobcount == processed {
				quit = true

			}

		case <-time.After(3 * time.Second):
			quit = true
		}
	}

	t1 = time.Since(t0)
	fmt.Printf("Concurency completed in %v \n", t1)
}

func printFileAndSize() {
	t0 := time.Now()
	dirname := "/Users/ariefdarmawan/Dropbox"
	c, s := getFileAndSize(dirname)
	t1 := time.Since(t0)
	fmt.Printf("Processing %s in %v => has %d files size of %v Mb \n", dirname, t1, c, (s / 1024 / 1024))
}

func getFileAndSize(dirname string) (int, int64) {
	var fileCount int = 0
	var size int64 = 0
	fis, _ := ioutil.ReadDir(dirname)
	for _, fi := range fis {
		if !fi.IsDir() {
			if !strings.HasPrefix(fi.Name(), ".") {
				fileCount = fileCount + 1
				size = size + fi.Size()
			}
		} else {
			if !strings.HasPrefix(fi.Name(), ".") {
				dCount, dSize := getFileAndSize(filepath.Join(dirname, fi.Name()))
				fileCount = fileCount + dCount
				size = size + dSize
			}
		}
	}
	return fileCount, size
}
