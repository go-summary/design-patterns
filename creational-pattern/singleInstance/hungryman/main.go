package hungryman

import "sync"

type Message struct {
	Count int
}

// 消息池
type messagePool struct {
	pool *sync.Pool
}

var msgPool = &messagePool{
	pool: &sync.Pool{
		New: func() any {
			return &Message{Count: 0}
		},
	},
}

// 单例 hungryman，初始化时就创建了实例
func Instance() *messagePool {
	return msgPool
}

func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}
