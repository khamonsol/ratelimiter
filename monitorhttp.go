package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"math/rand"
)

var (
	port            = flag.String("port", "8080", "Port to listen on")
	minDelay = flag.Duration("minDelay", 0, "Minimum delay for request processing (e.g., 10ms, 100ms, 1s)")
	maxDelay = flag.Duration("maxDelay", 0, "Maximum delay for request processing (e.g., 10ms, 100ms, 1s)")
	concurrencyChan = make(chan int, 10000)
	rateChan        = make(chan int, 10000)
	totalChan        = make(chan int, 10000)
	highestRate     int
	lastRate int
	highestConcurrent int
)

func main() {
	flag.Parse()


	http.HandleFunc("/", handler)

	// Handling concurrency using channel
	go func() {
		var currentConcurrent int
		for change := range concurrencyChan {
			currentConcurrent += change
			if currentConcurrent > highestConcurrent {
				highestConcurrent = currentConcurrent
			}
		}
	}()

	// Handling rate measuring using channel
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		var currentRate int
		for {
			select {
			case <-rateChan:
				currentRate++
			case <-ticker.C:
				if currentRate > highestRate {
					highestRate = currentRate
				}
				if highestRate == 0 {
					highestRate = currentRate
				}
				lastRate = currentRate
				currentRate = 0
			}
		}
	}()

	// Goroutine for processing and displaying messages
	go func() {
		blinker := true
		var lhR int
		var lR int
		var lHc int
		for {
			time.Sleep(500 * time.Millisecond)
			var message string
			if lastRate == 0 {
				if lhR + lR + lHc > 0 {
					message=""
					fmt.Println(fmt.Sprintf("\rPrevious run: req/100ms(HWM): %d | req/100ms(last): %d |Concurrent(HWM): %d", lhR, lR, lHc))
					fmt.Println("\n")
					lhR = 0
					lR = 0
					lHc = 0
				}
				message = "Waiting for traffic ..."
			}else {
				lhR = highestRate
				lR = lastRate
				lHc = highestConcurrent
				message = fmt.Sprintf("req/100ms(HWM): %d | req/100ms(last): %d |Concurrent(HWM): %d" , lhR,lR, lHc)
			}

			if blinker {
				fmt.Printf("\r%s ", message)
			} else {
				fmt.Printf("\r%s", message)
			}
			blinker = !blinker
		}
	}()

	fmt.Printf("Server started on port %s with simulated processing time bewteen %s and %s. Enjoy the test\n\n",
		*port,
		minDelay.String(),
		maxDelay.String())
	http.ListenAndServe(":"+*port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	concurrencyChan <- 1
	rateChan <- 1
	totalChan <- 1

	if *minDelay > 0 || *maxDelay > 0 {
		randomSleep := *minDelay + time.Duration(rand.Int63n(int64(*maxDelay-*minDelay)))
		time.Sleep(randomSleep)
	}

	defer func() {
		concurrencyChan <- -1
	}()
	w.Write([]byte("Hello, World!"))
}
