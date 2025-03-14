package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"test_go/internal/agent"
	"time"
)

var cnt int64
var addCounterChan chan int64

//var readCounterChan chan int64

//func AddCounter(ch chan int) {
//	ch <- 1
//}

func main() {

	//readCounterChan = make(chan int64, 100)

	//Start reader
	//go readApplicationLogs()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	cnt = 1
	for _ = range ticker.C {
		addCounterChan = make(chan int64)
		agent.NewMonitor(2, addCounterChan, cnt)

		cnt = cnt + <-addCounterChan
		//fmt.Println(cnt)
		//var id int
		//err := row.Scan(&id)

	}
	//Wait for exit
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//<-sigs
	log.Println("received", <-sigs)
	//log.Info("Received kill signal")
	//var tim
}
