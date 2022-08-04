package factory_method

// 工厂方法模式跟建造者模式类似，都是将对象创建的逻辑封装起来，为使用者提供一个简单易用的对象创建接口。
// 两者在应用场景上稍有区别，建造者模式更常用于需要传递多个参数来进行实例化的场景。

// 使用工厂方法来创建对象主要有两个好处：
// 1、代码可读性更好。相比于使用C++/Java中的构造函数，或者Go中的{}来创建对象，工厂方法因为可以通过函数名来表达代码含义，从而具备更好的可读性。
//比如，使用工厂方法productA := CreateProductA()创建一个ProductA对象，比直接使用productA := ProductA{}的可读性要好。
// 2、与使用者代码解耦。很多情况下，对象的创建往往是一个容易变化的点，通过工厂方法来封装对象的创建过程，可以在创建逻辑变更时，避免霰弹式修改。

// 工厂方法模式也有两种实现方式：
//（1）提供一个工厂对象，通过调用工厂对象的工厂方法来创建产品对象；
//（2）将工厂方法集成到产品对象中（C++/Java中对象的static方法，Go中同一package下的函数）

type Type uint8

const (
	Start Type = iota
	End
)

// 事件抽象接口
type Event interface {
	EventType() Type
	Content() string
}

// 开始事件，实现Event接口
type StartEvent struct {
	content string
}

func (e *StartEvent) EventType() Type {
	return Start
}
func (e *StartEvent) Content() string {
	return e.content
}

// 结束事件，实现Event接口
type EndEvent struct {
	content string
}

func (e *EndEvent) EventType() Type {
	return End
}
func (e *EndEvent) Content() string {
	return e.content
}

// 实现工厂对象, 集成到一个方法里
type Factory struct{}

func (e *Factory) Create(etype Type) Event {
	switch etype {
	case Start:
		return &StartEvent{
			content: "this is start event",
		}
	case End:
		return &EndEvent{
			content: "this is end event",
		}
	default:
		return nil
	}
}

// 第二种实现, 在factory里面给每个对象创建工厂方法
func OfStart() Event {
	return &StartEvent{
		content: "this is start event",
	}
}

func OfEnd() Event {
	return &EndEvent{
		content: "this is end event",
	}
}
