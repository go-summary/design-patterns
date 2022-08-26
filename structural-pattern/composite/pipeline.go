package composite

import "fmt"

type Type uint8

const (
	InputType = iota
	FilterType
	OutputType
)

// 消息管道
type Pipeline struct {
	status Status
	input  Input
	filter Filter
	output Output
}

// msg处理流程 input -> filter -> output
func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

func (p *Pipeline) Start() {
	p.output.Start()
	p.filter.Start()
	p.input.Start()
	p.status = Started
	fmt.Println("Hello input plugin started.")
}

func (p *Pipeline) Stop() {
	p.output.Stop()
	p.filter.Stop()
	p.input.Stop()
	p.status = Stopped
	fmt.Println("Hello input plugin Stopped.")
}

func (p *Pipeline) Status() Status {
	return p.status
}

type PipelineConfig struct {
	Input, Filter, OutPut PluginConfig
}

// 保存用于创建插件的工厂实例，key为插件类型，value为抽象工厂接口
var pluginFactories = make(map[Type]Factory)

func factoryOf(t Type) Factory {
	factory, _ := pluginFactories[t]
	return factory
}

func Of(conf PipelineConfig) *Pipeline {
	p := &Pipeline{}
	p.input = factoryOf(InputType).Create(conf.Input).(Input)
	p.filter = factoryOf(FilterType).Create(conf.Filter).(Filter)
	p.output = factoryOf(OutputType).Create(conf.OutPut).(Output)
	return p
}

func DefaultConfig() PipelineConfig {
	conf := PipelineConfig{}
	conf.Input = PluginConfig{Name: "hello"}
	conf.Filter = PluginConfig{Name: "upper"}
	conf.OutPut = PluginConfig{Name: "console"}
	return conf
}

// 初始化工厂插件
func init() {
	pluginFactories[InputType] = &InputFactory{}
	pluginFactories[FilterType] = &FilterFactory{}
	pluginFactories[OutputType] = &OutPutFactory{}
}
