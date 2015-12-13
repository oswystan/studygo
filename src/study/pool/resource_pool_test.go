//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: resource_pool_test.go
//     description:
//         created: 2015-12-12 22:40:11
//          author: wystan
//
//===============================================================================
package main

import (
	"fmt"
	"testing"
)

type item struct {
	/* data */
	value int
}

func newItem() interface{} {
	return &item{value: 1}
}

func TestLazyPool(t *testing.T) {
	p := NewResourcePool(LAZY_POOL, 10, newItem)
	var slice []*item = make([]*item, 10)
	for i := 0; i < 10; i++ {
		v := p.Acquire()
		if v == nil {
			t.Errorf("fail to acquire resource")
		}
		slice[i] = v.(*item)
		if slice[i].value != 1 {
			t.Errorf("invalid item value=%d", slice[i].value)
		}
		fmt.Printf("%s\n", p.Status())
	}

	v := p.Acquire()
	if v != nil {
		t.Errorf("acquire too much resource but return one")
	}

	for i := 0; i < 10; i++ {
		p.Release(slice[i])
	}
	fmt.Printf("%s\n", p.Status())

	for i := 0; i < 10; i++ {
		v := p.Acquire()
		if v == nil {
			t.Errorf("fail to acquire resource")
		}
		slice[i] = v.(*item)
		if slice[i].value != 1 {
			t.Errorf("invalid item value=%d", slice[i].value)
		}
	}

	p.Reset()
}

func TestEagerPool(t *testing.T) {
	p := NewResourcePool(EAGER_POOL, 10, newItem)
	var slice []*item = make([]*item, 10)
	for i := 0; i < 10; i++ {
		v := p.Acquire()
		if v == nil {
			t.Errorf("fail to acquire resource")
		}
		slice[i] = v.(*item)
	}

	v := p.Acquire()
	if v != nil {
		t.Errorf("acquire too much resource but return one")
	}

	for i := 0; i < 10; i++ {
		p.Release(slice[i])
	}
}

//==================================== END ======================================
