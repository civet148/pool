package main

import (
	"github.com/civet148/log"
	"github.com/civet148/pool"
)

var count int

type Object struct {
	Count int
}

func main() {
	p := pool.New(CreateObject)
	for i := 0; i < 10; i++ {
		obj := p.Get()
		if obj == nil {
			log.Panic("Get object from pool return nil")
		}
		p.Put(obj)
		log.Infof("object[%d] obj %+v total %d", i, obj, p.Len())
	}
	p.RemoveAll()
}

func CreateObject() interface{} {
	count++
	return &Object{
		Count: count,
	}
}
