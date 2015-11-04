package main

import (
	"time"
)

func MakeCoffee(done chan string) {
	log.Println("Starting coffee.")
	time.Sleep(time.Duration(5) * time.Second)
	done <- "coffee"
}

func MakeToast(done chan string) {
	log.Println("Starting toast.")
	time.Sleep(time.Duration(3) * time.Second)
	done <- "toast"
}

func MakeEggs(done chan string) {
	log.Println("Starting eggs.")
	time.Sleep(time.Duration(5) * time.Second)
	done <- "eggs"
}

func Gather(done chan string) []string {
	var doneList []string
	for i := 0; i < 3; i++ {
		item := <-done
		doneList = append(doneList, item)
	}
	return doneList
}

func main() {
	doneCh := make(chan string)
	MakeCoffee(doneCh)
	MakeToast(doneCh)
	MakeEggs(doneCh)
	_ = Gather(doneCh)

}
