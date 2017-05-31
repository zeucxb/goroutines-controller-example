package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup
var stop chan bool

func init() {
	stop = make(chan bool, 1)
}

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Second * 10)

	wg.Add(2)

	go func() {
		for range ticker.C {
			for i := 0; i < 10; i++ {
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
				select {
				case <-stop:
					wg.Done()
					return
				default:
					wg.Add(1)
					go process()
				}
			}
		}
	}()

	go func() {
		defer wg.Done()

		<-sigs
		stop <- true
		fmt.Println("\nCalma ae, ja vai")
		ticker.Stop()
	}()

	wg.Wait()
}

func process() {
	defer wg.Done()
	fmt.Println("INIT Tick")
	time.Sleep(time.Second * 9)
	fmt.Println("END Tick")
}
