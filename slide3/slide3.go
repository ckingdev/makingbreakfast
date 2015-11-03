package main

import (
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
)

func MakeCoffee() chan string {
	ret := make(chan string)
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		ret <- "Coffee is done."
	}()
	return ret
}

func MakeToast() chan string {
	ret := make(chan string)
	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		ret <- "Toast is done"
	}()
	return ret
}

func main() {
	runtime.GOMAXPROCS(2)
	log.Println("Starting...")
	coffee := MakeCoffee()
	toast := MakeToast()
	select {
	case msg := <-toast:
		log.Println(msg)
	case msg := <-coffee:
		log.Println(msg)
	}
	time.Sleep(time.Duration(10) * time.Second)
	log.Println("Breakfast finished.")
}
