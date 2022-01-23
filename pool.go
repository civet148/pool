package pool

import (
	"container/list"
	"sync"
)

type Pool struct {
	queue  list.List
	locker sync.RWMutex
	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() interface{}
}

func New(factory func() interface{}) *Pool {
	return &Pool{
		New: factory,
	}
}

func (p *Pool) Get() interface{} {
	p.locker.Lock()
	defer p.locker.Unlock()
	e := p.queue.Front()

	var v interface{}
	if e == nil {
		v = p.New()
	} else {
		v = e.Value
		p.queue.Remove(e)
	}
	return v
}

func (p *Pool) Put(v interface{}) {
	p.locker.Lock()
	defer p.locker.Unlock()
	p.queue.PushBack(v)
}

func (p *Pool) Len() int {
	p.locker.RLock()
	defer p.locker.RUnlock()
	return p.queue.Len()
}

func (p *Pool) RemoveAll() {
	p.locker.Lock()
	defer p.locker.Unlock()
	var next *list.Element
	for e := p.queue.Front(); e != nil; e = next {
		next = e.Next()
		p.queue.Remove(e)
	}
}
