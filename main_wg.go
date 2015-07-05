package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

const maxProc = 10
const workPath = "/home/juragan360/temp/"

func main() {
	log.SetOutput(w)
	wg := sync.WaitGroup{}
	wg.Add(maxProc)
	t0 := time.Now()
	for i := 1; i <= maxProc; i++ {
		go func(i int) {
			defer wg.Done()
			iprocessed := i
			itimer := time.Duration(maxProc-iprocessed) * time.Second
			time.Sleep(itimer)
			log.Printf("Print sequence: %v wait for %v \n", iprocessed, itimer)
			filename := fmt.Sprintf(workPath+"goroutine_%d.txt", iprocessed)
			if e := ioutil.WriteFile(filename, []byte("Test GoRoutine"), 0644); e != nil {
				log.Println("Error: " + e.Error())
			}
		}(i)
	}
	isDone := false
	defer func() {
		duration := time.Since(t0)
		if !isDone {
			log.Println("Not Yet Done")
		} else {
			log.Println("Done")
		}
		log.Printf("Time lapsed: %v \n", duration)
	}()

	wg.Wait()
	isDone = true
}
