//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: message_hub_test.go
//     description:
//         created: 2015-12-13 15:29:26
//          author: wystan
//
//===============================================================================

package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

type sub struct {
	/* data */
	name string
}

func (s *sub) NewMessage(msg interface{}) {
	log.Printf("[%s]get new message %v", s.name, msg)
}

func TestTopic(t *testing.T) {
	hub := NewMessageHub()
	err := hub.NewTopic("wangyu")
	if err != nil {
		t.Errorf("new topic fail %s", err)
	}
	err = hub.NewTopic("pengyan")
	if err != nil {
		t.Errorf("new topic fail %s", err)
	}

	err = hub.DelTopic("wangyu")
	if err != nil {
		t.Errorf("del topic fail %s", err)
	}

	hub.Reset()
	hub = nil
}

func TestMessagePubSub(t *testing.T) {
	hub := NewMessageHub()

	err := hub.NewTopic("wong")
	if err != nil {
		t.Errorf("new topic fail %s", err)
	}
	err = hub.NewTopic("momo")
	if err != nil {
		t.Errorf("new topic fail %s", err)
	}

	s1 := &sub{name: "wong"}
	s2 := &sub{name: "momo"}
	hub.Subscribe("wong", s1)
	hub.Subscribe("momo", s2)

	hub.PublishMessage("wong", "hello wong")
	hub.PublishMessage("momo", "hello mo")

	hub.Reset()
	hub = nil
}

type EchoArchiver struct {
	/* data */
	ch     chan interface{}
	chQuit chan int
	wg     sync.WaitGroup
}

func (ar *EchoArchiver) IsBusy() bool {
	ret := len(ar.ch) == cap(ar.ch)
	return ret
}
func (ar *EchoArchiver) ArchiveMessage(msg interface{}) {
	ar.ch <- msg
}

func (ar *EchoArchiver) Start() error {
	go func() {
		ar.wg.Add(1)
	LOOP:
		for {
			select {
			case msg := <-ar.ch:
				log.Printf("archive msg %v", msg)
			case <-ar.chQuit:
				break LOOP
			}
		}
		log.Printf("echo archiver quit")
		ar.wg.Done()
	}()
	return nil
}
func (ar *EchoArchiver) Stop() {
	ar.chQuit <- 1
	ar.wg.Wait()
	log.Printf("echo archiver stoped")
	close(ar.ch)
	close(ar.chQuit)
}
func (ar *EchoArchiver) Flush() {
	for len(ar.ch) > 0 {
		time.Sleep(1 * time.Millisecond)
	}
}

func NewEchoArchiver() *EchoArchiver {
	ar := &EchoArchiver{
		ch:     make(chan interface{}, 20),
		chQuit: make(chan int, 1),
	}
	return ar
}

func TestMessageArchive(t *testing.T) {
	hub := NewMessageHub()
	ar := NewEchoArchiver()
	ar.Start()
	hub.RegisterArchiver(ar)

	err := hub.NewTopic("wong")
	if err != nil {
		t.Errorf("new topic fail %s", err)
	}

	s1 := &sub{name: "wong"}
	hub.Subscribe("wong", s1)

	for i := 0; i < 30; i++ {
		msg := fmt.Sprintf("hello wong %d", i)
		hub.PublishMessage("wong", msg)
		time.Sleep(10 * time.Millisecond)
	}

	hub.Flush()
	ar.Flush()
	ar.Stop()
	hub.Reset()
	ar = nil
	hub = nil
}

//==================================== END ======================================
