package lazyman

import (
	"github.com/go-summary/design-patterns/singleInstance/hungryman"
	"sync"
)

var once = &sync.Once{}

// 消息池
type messagePool struct {
	pool *sync.Pool
}

var msgPool *messagePool

func Instance() *messagePool {
	once.Do(func() {
		msgPool = &messagePool{
			// 如果消息池里没有消息，则新建一个Count值为0的Message实例
			pool: &sync.Pool{
				New: func() any {
					return &hungryman.Message{Count: 0}
				},
			},
		}
	})
	return msgPool
}

func (m *messagePool) AddMsg(msg *hungryman.Message) {
	m.pool.Put(msg)
}

func (m *messagePool) GetMsg() *hungryman.Message {
	return m.pool.Get().(*hungryman.Message)
}
