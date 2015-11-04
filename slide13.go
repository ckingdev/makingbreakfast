package main

import (
	"sync"
	"time"
)

func MakeCoffee(done *Done) {
	log.Println("Starting coffee.")
	time.Sleep(time.Duration(5) * time.Second)
	done.Append("coffee")
}

func MakeToast(done *Done) {
	log.Println("Starting toast.")
	time.Sleep(time.Duration(3) * time.Second)
	done.Append("toast")
}

func MakeEggs(done *Done) {
	log.Println("Starting eggs.")
	time.Sleep(time.Duration(5) * time.Second)
	done.Append("eggs")
}

type Done struct {
	List string
	Lock sync.Mutex
}

func (d *Done) Append(item string) {
	d.Lock.Lock()
	d.List = append(d.List, item)
	d.Lock.Unlock()
}

func main() {
	var done Done
	MakeCoffee(done)
	MakeEggs(done)
	MakeToast(done)
}
