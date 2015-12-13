//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: resource_pool.go
//     description:
//         created: 2015-12-12 22:32:03
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
	LAZY_POOL = iota
	EAGER_POOL
)

type Creator func() interface{}

type ResourcePool struct {
	sync.RWMutex            //protect the following resources
	creator      Creator    //function for creating resource
	avaliable    *list.List //the resources that can be acquired
	amount       int        //total amount that created
	max          int        //how may resources we have
}

func (lp *ResourcePool) Acquire() interface{} {
	lp.Lock()
	defer lp.Unlock()
	if lp.avaliable.Len() > 0 {
		e := lp.avaliable.Front()
		lp.avaliable.Remove(e)
		return e.Value
	}

	if lp.amount < lp.max {
		log.Printf("create a new one amount=%d", lp.amount)
		rv := lp.creator()
		lp.amount++
		return rv
	}

	return nil
}

func (lp *ResourcePool) Release(r interface{}) {
	lp.Lock()
	if lp.avaliable.Len()+1 > lp.amount {
		panic("some resouce is not in this pool")
	}
	lp.avaliable.PushBack(r)
	lp.Unlock()
}

func (lp *ResourcePool) Status() string {
	lp.RLock()
	defer lp.RUnlock()
	return fmt.Sprintf("max:%d allocated:%d acquired:%d avaliable:%d",
		lp.max,
		lp.amount,
		lp.amount-lp.avaliable.Len(),
		lp.avaliable.Len())
}

func (lp *ResourcePool) Reset() {
	lp.Lock()
	defer lp.Unlock()

	for lp.avaliable.Front() != nil {
		e := lp.avaliable.Front()
		lp.avaliable.Remove(e)
	}
	lp.avaliable = nil
	lp.creator = nil
	lp.amount = 0
	lp.max = 0
}

func NewResourcePool(poolType int, max int, fn Creator) *ResourcePool {
	p := &ResourcePool{
		creator:   fn,
		avaliable: list.New(),
		max:       max,
		amount:    0,
	}

	if poolType == EAGER_POOL {
		for i := 0; i < max; i++ {
			v := p.creator()
			p.avaliable.PushBack(v)
		}
		p.amount = max
	}

	return p
}

//==================================== END ======================================
