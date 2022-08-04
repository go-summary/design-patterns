package abstract_factory

// 消息管道
type Pipeline struct {
	input  Input
	filter Filter
	output OutPut
}

type PipelineConfig struct {
	Input, Filter, OutPut PluginConfig
}

// msg处理流程 input -> filter -> output
func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

// 保存用于创建插件的工厂实例，key为插件类型，value为抽象工厂接口
var pluginFactories = make(map[Type]Factory)

func factoryOf(t Type) Factory {
	factory, _ := pluginFactories[t]
	return factory
}

// pipeline 工厂方法，根据配置创建一个Pipelie实例
func Of(conf PipelineConfig) *Pipeline {
	p := &Pipeline{}
	p.input = factoryOf(InputType).Create(conf.Input).(Input)
	p.filter = factoryOf(FilterType).Create(conf.Filter).(Filter)
	p.output = factoryOf(OutputType).Create(conf.OutPut).(OutPut)
	return p
}

func DefaultConfig() PipelineConfig {
	conf := PipelineConfig{}
	conf.Input = PluginConfig{Name: "msg"}
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
