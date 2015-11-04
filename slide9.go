package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func MakeCoffee() chan string {
	log.Println("Starting coffee.")
	ret := make(chan string)
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		ret <- "Coffee is done."
		close(ret)
	}()
	return ret
}

func MakeToast() chan string {
	log.Println("Starting toast.")
	ret := make(chan string)
	go func() {
		time.Sleep(time.Duration(3) * time.Second)
		ret <- "Toast is done"
		close(ret)
	}()
	return ret
}

func fanIn(chs ...chan string) chan string {
	fannedIn := make(chan string)
	var wg sync.WaitGroup
	for _, ch := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range ch {
				fannedIn <- msg
			}
		}()
	}
	go func() {
		wg.Wait()
		close(fannedIn)
	}()
	return fannedIn
}

func main() {
	runtime.GOMAXPROCS(2)
	log.Println("Starting...")

	coffee := MakeCoffee()
	toast := MakeToast()

	for msg := range fanIn(coffee, toast) {
		log.Println(msg)
	}
	time.Sleep(time.Duration(10) * time.Second)
	log.Println("Breakfast finished.")
}
