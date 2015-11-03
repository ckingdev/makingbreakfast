package main

import (
	"runtime"
	"sync"
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

	for i := 0; i < 2; i++ {
		select {
		case msg := <-toast:
			log.Println(msg)
		case msg := <-coffee:
			log.Println(msg)
		}
	}

	time.Sleep(time.Duration(10) * time.Second)
	log.Println("Breakfast finished.")
}
