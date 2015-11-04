package main

import "time"
import "log"

func MakeCoffee() {
	log.Println("Starting coffee.")
	time.Sleep(time.Duration(5) * time.Second)
	log.Println("Coffee is done.")
}

func MakeToast() {
	log.Println("Starting toast.")
	time.Sleep(time.Duration(3) * time.Second)
	log.Println("Toast is done.")
}

func main() {
	log.Println("Starting...")
	MakeCoffee()
	MakeToast()
	log.Println("Breakfast finished.")
}
