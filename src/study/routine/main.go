/*
 * *******************************************************************************
 *                     Copyright (c) 2015, wystan
 *       Filename:  main.go
 *
 *    Description:
 *
 *        Created:  2015-11-08 14:16:58
 *       Revision:  none
 *
 *         Author:  wystan
 *
 * *******************************************************************************
 */

package main

import (
	"fmt"
	"log"
	"sync"
)

type Handler interface {
	/* methods */
	HandleMessage(interface{}) interface{}
}

const (
	SERVICE_STOP = iota
	SERVICE_RUNNING
)

type MicroServiceRunner struct {
	handler Handler
	wg      sync.WaitGroup
	chr     chan interface{}
	chw     chan interface{}
	chq     chan int
	status  int
}

func (sr *MicroServiceRunner) loop() {
	sr.wg.Add(1)

LOOP:
	for {
		select {
		case data := <-sr.chr:
			result := sr.handler.HandleMessage(data)
			if result != nil {
				sr.chw <- result
			}

		case <-sr.chq:
			log.Printf("request exit..")
			break LOOP
		}
	}

	log.Printf("loop exit!!")
	sr.status = SERVICE_STOP
	sr.wg.Done()
}

func (sr *MicroServiceRunner) Start(h Handler) {
	if sr.status != SERVICE_STOP {
		return
	}

	if h != nil {
		sr.handler = h
	}
	sr.chr = make(chan interface{}, 10)
	sr.chw = make(chan interface{}, 10)
	sr.chq = make(chan int)

	sr.status = SERVICE_RUNNING
	go sr.loop()
}
func (sr *MicroServiceRunner) Stop() {
	// do not check the current status.
	// do the following step anyway
	sr.chq <- 1
	sr.wg.Wait()
	close(sr.chr)
	close(sr.chw)
	close(sr.chq)
}
func (sr *MicroServiceRunner) Status() int {
	return sr.status
}
func (sr *MicroServiceRunner) Send(data interface{}) error {
	if sr.status != SERVICE_RUNNING {
		return fmt.Errorf("service is not running")
	}
	sr.chr <- data
	return nil
}
func (sr *MicroServiceRunner) Receive() (interface{}, error) {
	if sr.status != SERVICE_RUNNING {
		return nil, fmt.Errorf("service is not running")
	}
	data := <-sr.chw
	return data, nil
}

func NewMicroServiceRunner(h Handler) *MicroServiceRunner {
	return &MicroServiceRunner{
		handler: h,
		status:  SERVICE_STOP,
	}
}

type addserver struct {
	/* data */
}

func (s *addserver) HandleMessage(d interface{}) interface{} {
	a := d.(int)
	a = a + 10
	return a
}

type strserver struct {
	/* data */
}

func (s *strserver) HandleMessage(d interface{}) interface{} {
	a := d.(string)
	a = "hello " + a
	return a
}

/*
 * main entry
 */
func main() {
	as := NewMicroServiceRunner(&strserver{})
	as.Start(nil)

	for i := 0; i < 10; i++ {
		as.Send(fmt.Sprintf("%d", i))
	}
	for i := 0; i < 10; i++ {
		result, err := as.Receive()
		if err == nil {
			fmt.Printf("result=%s\n", result.(string))
		}
	}

	fmt.Printf("status : %d\n", as.Status())
	as.Stop()
	fmt.Printf("status : %d\n", as.Status())
}

/**************************************** END ***********************************/
