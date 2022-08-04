package test

import (
	"github.com/go-summary/design-patterns/creational-pattern/abstract-factory"
	"testing"
)

func TestPipeline(t *testing.T) {
	// 其中pipeline.DefaultConfig()的配置内容见【抽象工厂模式示例图】
	// 消息处理流程为 HelloInput -> UpperFilter -> ConsoleOutput
	p := abstract_factory.Of(abstract_factory.DefaultConfig())
	p.Exec()
}
