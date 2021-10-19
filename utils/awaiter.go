//
//  awaiter.go
//  main
//
//  Created by d-exclaimation on 11:56 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package utils

type Await struct {
	isDone bool
	_check chan chan bool
	_done  chan bool
}

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

func (a *Await) dispatch(op func()) {
	go func() {
		defer a.end()
		op()
	}()
}

func (a *Await) run() {
	go a.onReceive()
}

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

func (a *Await) end() {
	a._done <- true
}

func (a *Await) IsDone() bool {
	respond := make(chan bool)
	a._check <- respond
	return <-respond
}
