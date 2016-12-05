package loadbalance

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestSingleWorkerWithSingleJob(t *testing.T) {
	worker := Worker{
		requests: make(chan Request, 10),
		pending: 0,
		index: -1,
	}

	pool := Pool{&worker,}

	balancer := Balancer{
		pool: pool,
		done: make(chan *Worker, 10),
	}

	// goroutine for work
	go func() {
		worker.work(balancer.done)
	}()

	requests := make(chan Request)

	// goroutine for balanced routing
	go func() {
		balancer.balance(requests)
	}()

	response := make(chan interface{})

	requests <- Request{
		fn: func() interface{} {
			time.Sleep(time.Duration(rand.Int63n(2000)) * time.Millisecond)
			return "excuted"
		},
		c: response,
	}

	value := <-response
	fmt.Printf("%v\n", value)
}