package loadbalance

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
}

func TestSingleWorkerWithSingleJob(t *testing.T) {
	worker := Worker{
		requests: make(chan Request, 10),
		pending:  0,
		index:    0,
	}

	pool := Pool{&worker}

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

type NumberedWorker struct {
	Worker
	no int
}

func TestMultipleWorker(t *testing.T) {
	requestChannel := make(chan Request, 10)
	responseChannel := make(chan interface{}, 10)
	jobDoneSignalChannel := make(chan *Worker, 10)

	defer func() {
		defer close(requestChannel)
		defer close(responseChannel)
		defer close(jobDoneSignalChannel)
	}()

	pool := make(Pool, 0, 5)
	for i := 0; i < cap(pool); i++ {
		w := Worker{
			requests: make(chan Request, 20),	// this buffer size is important to avoid deadlock(all goroutines are waiting)
			pending:  0,
			index:    -1,
		}

		pool.Push(&w)

		go func() {
			w.work(jobDoneSignalChannel)
		}()
	}

	balancer := Balancer{pool, jobDoneSignalChannel}

	go func() {
		balancer.balance(requestChannel)
	}()

	var wg sync.WaitGroup

	go func() {
		for v := range responseChannel {
			go func(val interface{}) {
				log.Printf("<- %d", val)
				wg.Done()
			}(v)
		}
	}()

	nWorker := int64(len(pool))
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 100; i++ {
		genDurationMsecs := time.Duration(rand.Int63n(nWorker * 1000))
		workDurationMsecs := time.Duration(rand.Int63n(nWorker * 1000))

		go func(val interface{}) {
			time.Sleep(genDurationMsecs * time.Millisecond)

			requestChannel <- Request{
				fn: func() interface{} {
					time.Sleep(workDurationMsecs * time.Millisecond)
					return val
				},
				c: responseChannel,
			}
			log.Printf("%d(msecs)-load generated : expected return is %v", workDurationMsecs, val)
		}(i)

		wg.Add(1)
	}

	wg.Wait()
}
