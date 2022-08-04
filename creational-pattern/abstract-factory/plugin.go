package abstract_factory

import (
	"fmt"
	"reflect"
	"strings"
)

// 抽象工厂模式通过给工厂类新增一个抽象层解决了该问题，假如FactoryA和FactoryB都实现·抽象工厂接口，分别用于创建ProductA和ProductB。
// 如果后续新增了ProductC，只需新增一个FactoryC即可，无需修改原有的代码；
// 因为每个工厂只负责创建一个产品，因此也遵循了单一职责原则。
//
// Go实现
// 考虑需要如下一个插件架构风格的消息处理系统，pipeline是消息处理的管道，其中包含了input、filter和output三个插件。
// 我们需要实现根据配置来创建pipeline ，加载插件过程的实现非常适合使用工厂模式，
// 其中input、filter和output三类插件的创建使用抽象工厂模式，而pipeline的创建则使用工厂方法模式。

// 插件抽象接口定义
// 抽象工厂模式
type Plugin interface{}
type Type uint8

const (
	InputType = iota
	FilterType
	OutputType
)

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

// 实现3个插件
// input插件名称与类型的映射关系，主要用于通过反射创建input对象
var inputNames = make(map[string]reflect.Type)

// Hello input插件，接收"Hello World"消息
type HelloInput struct{}

func (h *HelloInput) Receive() string {
	return "Hello World"
}

type MsgInput struct{}

func (m *MsgInput) Receive() string {
	return "Msg Send"
}

// init的时候初始化映射表
func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
	inputNames["msg"] = reflect.TypeOf(MsgInput{})
}

// filter插件名称与类型的映射关系，主要用于通过反射创建filter对象
var filterNames = make(map[string]reflect.Type)

// Upper filter插件，将消息全部字母转成大写
type UpperFilter struct{}

func (u *UpperFilter) Process(msg string) string {
	return strings.ToUpper(msg)
}

// 初始化filter插件映射关系表
func init() {
	filterNames["upper"] = reflect.TypeOf(UpperFilter{})
}

// output插件名称与类型的映射关系，主要用于通过反射创建output对象
var outputNames = make(map[string]reflect.Type)

// Console output插件，将消息输出到控制台上
type ConsoleOutput struct{}

func (c *ConsoleOutput) Send(msg string) {
	fmt.Println(msg)
}

// 初始化output插件映射关系表
func init() {
	outputNames["console"] = reflect.TypeOf(ConsoleOutput{})
}

type PluginConfig struct {
	Name string
}

// 定义抽象工厂接口
type Factory interface {
	Create(conf PluginConfig) Plugin
}

// input插件工厂对象， 实现Factory接口
type InputFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *InputFactory) Create(conf PluginConfig) Plugin {
	t, _ := inputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type FilterFactory struct{}

func (f *FilterFactory) Create(conf PluginConfig) Plugin {
	t, _ := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type OutPutFactory struct{}

func (f *OutPutFactory) Create(conf PluginConfig) Plugin {
	t, _ := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}
