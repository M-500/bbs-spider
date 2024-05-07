//@Author: wulinlin
//@Description:
//@File:  mq_2
//@Version: 1.0.0
//@Date: 2024/05/04 23:14

package channel

import "sync"

type Broker struct {
	consumer []func(msg Msg)
	lock     sync.RWMutex
}

func (b *Broker) Send(m Msg) error {
	b.lock.RLock()
	defer b.lock.RUnlock()
	for _, f := range b.consumer {
		go func() {
			f(m)
		}()
	}
	return nil
}

func (b *Broker) Subscribe(cb func(ms Msg)) error {
	b.lock.Lock()
	defer b.lock.RLock()
	b.consumer = append(b.consumer, cb)
	return nil
}
