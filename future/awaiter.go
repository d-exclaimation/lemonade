//
//  awaiter.go
//  main
//
//  Created by d-exclaimation on 11:56 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package future

// Await (Future / Promise for void function)
//
// Data structure to run a time-consuming operation on a separate goroutine
// and monitor whether the operation finished
type Await struct {
	isDone bool           // Field for indicating whether the operation is done
	_check chan chan bool // Channel for requesting `isDone` concurrent safely
	_done  chan bool      // Channel for the operation to indicate finished operation
}

// Wait
//
// Run a function in an Await
func Wait(op func()) *Await {
	a := &Await{
		isDone: false,
		_check: make(chan chan bool),
		_done:  make(chan bool),
	}
	a.run()
	a.dispatch(op)
	return a
}

// Dispatch the operation and set up the callback when finished.
func (a *Await) dispatch(op func()) {
	go func() {
		defer func() { a._done <- true }()
		op()
	}()
}

// Run the Actor-like `onReceive` method on a separate goroutine.
func (a *Await) run() {
	go a.onReceive()
}

// Actor-like `onReceive` method to receive message from channels (concurrent-safe).
func (a *Await) onReceive() {
	for {
		select {
		case ref := <-a._check:
			ref <- a.isDone
		case _ = <-a._done:
			a.isDone = true
		}
	}
}

// IsDone
//
// Check whether operation finished (concurrent-safe)
func (a *Await) IsDone() bool {
	respond := make(chan bool)
	a._check <- respond
	return <-respond
}
