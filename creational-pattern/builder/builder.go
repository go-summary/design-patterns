package builder

import (
	prototype "github.com/go-summary/design-patterns/creational-pattern/prototype"
	"sync"
)

// 针对对象成员较多，创建对象逻辑较为繁琐的场景
// 适合使用建造者模式来进行优化

// 造者模式的作用有如下几个：
// 1、封装复杂对象的创建过程，使对象使用者不感知复杂的创建逻辑。
// 2、可以一步步按照顺序对成员进行赋值，或者创建嵌套对象，并最终完成目标对象的创建。
// 3、对多个对象复用同样的对象创建逻辑。

type Message struct {
	Header *Header
	Body   *Body
}

type Header struct {
	SrcAddr  string `json:"srcAddr, omitempty"`
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}
type Body struct {
	Items []string
}

// prototype接口
func (m *Message) Clone() prototype.Prototype {
	msg := *m
	return &msg
}

// 建造者模式的方式
type builder struct {
	once *sync.Once
	msg  *Message
}

func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg:  &Message{&Header{}, &Body{}},
	}
}

// 成员构建
func (b *builder) WithSrcPort(srcPort uint64) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}
func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}
func (b *builder) WithDestPort(destPort uint64) *builder {
	b.msg.Header.DestPort = destPort
	return b
}
func (b *builder) WithSrcAddr(srcAddr string) *builder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}
func (b *builder) WithHeaderItem(key, value string) *builder {
	b.once.Do(func() {
		b.msg.Header.Items = map[string]string{}
	})
	b.msg.Header.Items[key] = value
	return b
}

func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

func (b *builder) Build() *Message {
	return b.msg
}
