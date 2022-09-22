// Package memo provides
package memo

import "sync"

// A Memo caches the results of calling a Func.
type Memo struct {
	// f        Func
	// cache    map[string]*entry
	// mu       sync.Mutex // guards cache
	requests chan request
	cond     sync.Cond
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// New .
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.serve(f)
	return memo
}

// Get .
// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	// memo.mu.Lock()
	// e := memo.cache[key]
	// if e == nil {
	// 	e = &entry{ready: make(chan struct{})}
	// 	memo.cache[key] = e
	// 	memo.mu.Unlock()

	// 	e.res.value, e.res.err = memo.f(key)
	// 	close(e.ready)
	// } else {
	// 	memo.mu.Unlock()
	// 	<-e.ready
	// }
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) serve(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
