//===============================================================================
//                      Copyright (C) 2015 wystan
//
//        filename: conn_pool.go
//     description:
//         created: 2015-12-06 23:09:25
//          author: wystan
//
//===============================================================================
package imc_lib

import "sync"

type ConnPool struct {
	/* data */
	pool map[string]*userConn
	lock sync.RWMutex
}

func (p *ConnPool) GetUserConn(name string) *userConn {
	p.lock.RLock()
	defer p.lock.RUnlock()
	if v, ok := p.pool[name]; ok {
		return v
	}

	return nil
}
func (p *ConnPool) PutUserConn(u *userConn) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.pool[u.userName] = u
}

func (p *ConnPool) DelUserConn(u *userConn) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.pool, u.userName)
}
func (p *ConnPool) Clear() {
	p.lock.Lock()
	defer p.lock.Unlock()
	for k, _ := range p.pool {
		delete(p.pool, k)
	}
}

var connectionPoolInstance = &ConnPool{pool: make(map[string]*userConn)}

func GetConnPool() *ConnPool {
	return connectionPoolInstance
}

//==================================== END ======================================
