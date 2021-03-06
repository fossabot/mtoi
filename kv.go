// +build go1.9

package mtoi

import "sync"

type itemKV struct {
	Key   string
	Value interface{}
}

type KV struct {
	cap    int
	data   sync.Map
	stream chan *itemKV
}

func NewKV(cap int) *KV {
	if cap <= 2 {
		cap = 2
	}
	c := &KV{
		cap:    cap,
		stream: make(chan *itemKV, cap),
	}
	c.start()
	return c
}

func (c *KV) start() {
	go func() {
		for v, ok := <-c.stream; ok; v, ok = <-c.stream {
			if v != nil {
				for ; v != nil && ok; v, ok = <-c.stream {
					c.data.Store(v.Key, v.Value)
				}
			}
		}
	}()
}

func (c *KV) Close() {
	close(c.stream)
}

func (c *KV) Put(k string, v interface{}) {
	c.stream <- &itemKV{k, v}
	c.stream <- nil
}

func (c *KV) MulPut() (func(k string, v interface{}), func()) {
	return func(k string, v interface{}) { c.stream <- &itemKV{k, v} },
		func() { c.stream <- nil }
}

func (c KV) Get(k string) (interface{}, bool) {
	return c.data.Load(k)
}

func (c KV) Contain(k string) bool {
	_, ok := c.Get(k)
	return ok
}
