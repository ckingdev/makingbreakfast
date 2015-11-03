package main

import (
	"runtime"
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
	log.Println("Starting...")
	go MakeCoffee()
	go MakeToast()
	time.Sleep(time.Duration(10) * time.Second)
	log.Println("Breakfast finished.")
}
