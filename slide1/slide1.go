package main

import (
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
	log.Println("Starting...")
	MakeCoffee()
	MakeToast()
	log.Println("Breakfast finished.")
}
