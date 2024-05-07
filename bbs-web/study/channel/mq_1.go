//@Author: wulinlin
//@Description:
//@File:  mq_1
//@Version: 1.0.0
//@Date: 2024/05/04 22:23

package channel

import (
	"errors"
	"sync"
)

type Msg struct {
	Context string
}

type MQV1 struct {
	lock  sync.RWMutex
	chans []chan Msg
}

func (m *MQV1) Send(data Msg) error {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, ch := range m.chans {
		//ch <- data // 这么写不行，因为会阻塞住
		select {
		case ch <- data:
		default:
			return errors.New("消息队列已经满了")
		}
	}
	return nil
}

func (m *MQV1) Close() error {
	m.lock.Lock()
	chans := m.chans // 小技巧
	m.chans = nil
	m.lock.Unlock()
	for _, ch := range chans {
		close(ch)
	}
	return nil
}

func (m *MQV1) Subscribe(capacity int) (<-chan Msg, error) {
	tempChan := make(chan Msg, capacity)
	m.lock.Lock()
	defer m.lock.Unlock()
	m.chans = append(m.chans, tempChan)
	return tempChan, nil
}
