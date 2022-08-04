package prototype

// 原型复制抽象接口
type Prototype interface {
	Clone() Prototype
}
