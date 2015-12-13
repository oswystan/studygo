//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: message_hub.go
//     description:
//         created: 2015-12-13 13:49:25
//          author: wystan
//
//===============================================================================

package main

import (
	"container/list"
	"fmt"
	"log"
	"sync"
)

const (
	MESSAGE_CACHED_THRESHOLD = 20
)

type Subscriber interface {
	NewMessage(msg interface{})
}
type Archiver interface {
	ArchiveMessage(msg interface{})
	IsBusy() bool
}

type MessageHub struct {
	sync.RWMutex
	topics   map[string]*list.List
	archiver Archiver
	msglock  sync.RWMutex
	messages *list.List
}

func (hub *MessageHub) NewTopic(topic string) error {
	hub.Lock()
	defer hub.Unlock()
	if _, ok := hub.topics[topic]; ok {
		return fmt.Errorf("already have topic %s", topic)
	}

	hub.topics[topic] = list.New()
	return nil
}
func (hub *MessageHub) DelTopic(topic string) error {
	hub.Lock()
	defer hub.Unlock()

	v, ok := hub.topics[topic]
	if !ok {
		log.Printf("topic %s does not existed while deleting it")
		return nil
	}

	for v.Front() != nil {
		e := v.Front()
		v.Remove(e)
		e = nil
	}

	return nil
}
func (hub *MessageHub) PublishMessage(topic string, msg interface{}) error {
	hub.RLock()
	v, ok := hub.topics[topic]
	if !ok {
		hub.RUnlock()
		return fmt.Errorf("no topic %s found when publish", topic)
	}

	dup := list.New()
	dup.PushBackList(v)
	hub.RUnlock()
	for e := dup.Front(); e != nil; e = e.Next() {
		sub := e.Value.(Subscriber)
		sub.NewMessage(msg)
	}

	hub.msglock.Lock()
	defer hub.msglock.Unlock()
	hub.messages.PushBack(msg)

	if hub.archiver != nil {
		// trigger the archiver to store the message
		for hub.messages.Len() >= MESSAGE_CACHED_THRESHOLD && !hub.archiver.IsBusy() {
			e := hub.messages.Front()
			hub.archiver.ArchiveMessage(e.Value)
			hub.messages.Remove(e)
		}
	} else {
		hub.messages.Remove(hub.messages.Front())
	}

	return nil
}
func (hub *MessageHub) Flush() {
	hub.msglock.Lock()
	defer hub.msglock.Unlock()
	if hub.archiver == nil {
		hub.messages = hub.messages.Init()
		return
	}

	for e := hub.messages.Front(); e != nil; e = e.Next() {
		hub.archiver.ArchiveMessage(e.Value)
	}
	hub.messages = hub.messages.Init()
}

func (hub *MessageHub) Subscribe(topic string, sub Subscriber) error {
	hub.Lock()
	defer hub.Unlock()

	v, ok := hub.topics[topic]
	if !ok {
		return fmt.Errorf("no topic %s found when subscribe", topic)
	}

	v.PushBack(sub)
	return nil
}
func (hub *MessageHub) Unsubscribe(topic string, sub Subscriber) error {
	hub.Lock()
	defer hub.Unlock()

	v, ok := hub.topics[topic]
	if !ok {
		return fmt.Errorf("no topic %s found when unsubscribe", topic)
	}

	for e := v.Front(); e != nil; e = e.Next() {
		if e.Value.(Subscriber) == sub {
			v.Remove(e)
			return nil
		}
	}

	log.Printf("subscriber does not subscribe topic %s", topic)
	return nil
}
func (hub *MessageHub) RegisterArchiver(ar Archiver) error {
	if hub.archiver == nil {
		hub.archiver = ar
	} else {
		fmt.Errorf("already have a archiver %v", hub.archiver)
	}

	return nil
}
func (hub *MessageHub) Reset() {
	hub.Lock()
	defer hub.Unlock()

	for k, v := range hub.topics {
		v = v.Init()
		delete(hub.topics, k)
	}
	hub.messages = hub.messages.Init()
}

func NewMessageHub() *MessageHub {
	return &MessageHub{
		archiver: nil,
		topics:   make(map[string]*list.List),
		messages: list.New(),
	}
}

//==================================== END ======================================
