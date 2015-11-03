package main

import (
	"runtime"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

func MakeCoffee() {
	time.Sleep(time.Duration(5) * time.Second)
	log.Println("Coffee is done.")
}

func MakeToast() {
	time.Sleep(time.Duration(3) * time.Second)
	log.Println("Toast is done.")
}

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	log.Println("Starting...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		MakeCoffee()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		MakeToast()
	}()
	wg.Wait()
	log.Println("Breakfast finished.")
}
