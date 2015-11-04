package main

import (
	"time"

	"log"
)

func MakeCoffee(ch chan string) {
	log.Println("Starting coffee.")
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		ch <- "Coffee is done."
	}()
}

func MakeToast(ch chan string) {
	log.Println("Starting toast.")
	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		ch <- "Toast is done"
	}()
}

func main() {
	log.Println("Starting...")
	ch := make(chan string)
	MakeCoffee(ch)
	MakeToast(ch)
	for i := 0; i < 2; i++ {
		msg := <-ch
		log.Println(msg)
	}
	log.Println("Breakfast finished.")
}
