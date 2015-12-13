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
	"log"
	"testing"
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

//==================================== END ======================================
