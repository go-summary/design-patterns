package composite

import (
	"fmt"
	"github.com/go-summary/design-patterns/creational-pattern/builder"
	"reflect"
	"strings"
)

// 组合模式
// 在面向对象编程中，有两个常见的对象设计方法，组合和继承，两者都可以解决代码复用的问题，但是使用后者时容易出现继承层次过深，对象关系过于复杂的副作用，从而导致代码的可维护性变差。
// 因此，一个经典的面向对象设计原则是：组合优于继承。

// 我们都知道，组合所表示的语义为“has-a”，也就是部分和整体的关系，最经典的组合模式描述如下：
// 将对象组合成树形结构以表示“部分-整体”的层次结构，使得用户对单个对象和组合对象的使用具有一致性。
// Go语言天然就支持了组合模式，而且从它不支持继承关系的特点来看，Go也奉行了组合优于继承的原则，鼓励大家在进行程序设计时多采用组合的方法。
// Go实现组合模式的方式有两种，分别是直接组合（Direct Composition）和嵌入组合（Embedding Composition），下面我们一起探讨这两种不同的实现方法。

// 直接组合（Direct Composition）的实现方式类似于Java/C++，就是将一个对象作为另一个对象的成员属性。
// 比如之前设计的是后，Message结构中国就是由Header和Body组成，其成员就是一个对象

// 插件运行状态
type Status uint

const (
	Stopped Status = iota
	Started
)

type Plugin interface {
	Start()
	Stop()
	Status() Status
}

type Input interface {
	Plugin
	Receive() *builder.Message
}

type Filter interface {
	Plugin
	Process(msg *builder.Message) *builder.Message
}

type Output interface {
	Plugin
	Send(msg *builder.Message)
}

// 实现3个插件
// input插件名称与类型的映射关系，主要用于通过反射创建input对象
var inputNames = make(map[string]reflect.Type)

type HelloInput struct {
	status Status
}

func (h *HelloInput) Receive() *builder.Message {
	// 如果插件未启动，则返回nil
	if h.status != Started {
		fmt.Println("Hello input plugin is not running, input nothing.")
		return nil
	}
	return builder.Builder().
		WithHeaderItem("content", "text").
		WithBodyItem("Hello World").
		Build()
}

func (h *HelloInput) Start() {
	h.status = Started
	fmt.Println("Hello input plugin started.")
}

func (h *HelloInput) Stop() {
	h.status = Stopped
	fmt.Println("Hello input plugin stopped.")
}

func (h *HelloInput) Status() Status {
	return h.status
}

// init的时候初始化映射表
func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
}

// filter插件名称与类型的映射关系，主要用于通过反射创建filter对象
var filterNames = make(map[string]reflect.Type)

type UpperFilter struct {
	status Status
}

func (u *UpperFilter) Process(msg *builder.Message) *builder.Message {
	if u.status != Started {
		fmt.Println("Upper filter plugin is not running, filter nothing.")
		return msg
	}
	for i, val := range msg.Body.Items {
		msg.Body.Items[i] = strings.ToUpper(val)
	}
	return msg
}

func (u *UpperFilter) Start() {
	u.status = Started
	fmt.Println("Upper filter plugin started.")
}

func (u *UpperFilter) Stop() {
	u.status = Stopped
	fmt.Println("Upper filter plugin stopped.")
}

func (u *UpperFilter) Status() Status {
	return u.status
}

// 初始化filter插件映射关系表
func init() {
	filterNames["upper"] = reflect.TypeOf(UpperFilter{})
}

// output插件名称与类型的映射关系，主要用于通过反射创建output对象
var outputNames = make(map[string]reflect.Type)

type ConsoleOutput struct {
	status Status
}

func (c *ConsoleOutput) Send(msg *builder.Message) {
	if c.status != Started {
		fmt.Println("Console output is not running, output nothing.")
		return
	}
	fmt.Printf("Output:\n\tHeader:%+v, Body:%+v\n", msg.Header.Items, msg.Body.Items)
}

func (c *ConsoleOutput) Start() {
	c.status = Started
	fmt.Println("Console output plugin started.")
}

func (c *ConsoleOutput) Stop() {
	c.status = Stopped
	fmt.Println("Console output plugin stopped.")
}

func (c *ConsoleOutput) Status() Status {
	return c.status
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
