package loadbalance

// https://talks.golang.org/2012/waza.slide

import "container/heap"

/////////////////////////////////////////////////////////////////
//
// Request
//
/////////////////////////////////////////////////////////////////

type Request struct {
	fn func() interface{} // The operation to perform
	c  chan interface{}   // The channel to return the result
}

/////////////////////////////////////////////////////////////////
//
// Worker
//
/////////////////////////////////////////////////////////////////

type Worker struct {
	requests chan Request // works to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests // get Request from balancer
		req.c <- req.fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}

/////////////////////////////////////////////////////////////////
//
// Pool
//
/////////////////////////////////////////////////////////////////

type Pool []*Worker

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	w := x.(*Worker)
	w.index = n
	*p = append(*p, w)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	w := old[n-1]
	w.index = -1 // for safety
	*p = old[0 : n-1]
	return w
}

/////////////////////////////////////////////////////////////////
//
// Balancer
//
/////////////////////////////////////////////////////////////////

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) dispatch(req Request) {
	// grab the least loaded worker
	w := heap.Pop(&b.pool).(*Worker)
	// ... send it the task
	w.requests <- req
	// One more in its work queue
	w.pending++
	// Put it into its place on the heap
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue
	w.pending--
	// Remove it from heap
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap
	heap.Push(&b.pool, w)
}

func (b *Balancer) balance(requests chan Request) {
	for {
		select {
		case req := <-requests: // received a Request
			b.dispatch(req) // ... so send it to a worker
		case worker := <-b.done: // a worker has finished
			b.completed(worker) // ... so update its info
		}
	}
}
