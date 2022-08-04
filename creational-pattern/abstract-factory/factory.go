package abstract_factory

// 抽象工厂模式通过给工厂类新增一个抽象层解决了该问题，假如FactoryA和FactoryB都实现·抽象工厂接口，分别用于创建ProductA和ProductB。
// 如果后续新增了ProductC，只需新增一个FactoryC即可，无需修改原有的代码；
// 因为每个工厂只负责创建一个产品，因此也遵循了单一职责原则。
//
// Go实现
// 考虑需要如下一个插件架构风格的消息处理系统，pipeline是消息处理的管道，其中包含了input、filter和output三个插件。
// 我们需要实现根据配置来创建pipeline ，加载插件过程的实现非常适合使用工厂模式，
// 其中input、filter和output三类插件的创建使用抽象工厂模式，而pipeline的创建则使用工厂方法模式。

// 插件抽象接口定义
// 工厂方法模式
type Plugin interface{}

// 抽象工厂模式
// 过滤插件接口
type Filter interface {
	Plugin // 实现插件接口，继承插件接口
	Process(msg string) string
}

// 输入插件接口
type Input interface {
	Plugin
	Receive() string
}

// 输出插件接口
type OutPut interface {
	Plugin
	Send(msg string)
}
