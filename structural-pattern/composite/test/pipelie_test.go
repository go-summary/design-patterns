package test

import (
	"github.com/go-summary/design-patterns/structural-pattern/composite"
	"testing"
)

func TestPipeline(t *testing.T) {
	p := composite.Of(composite.DefaultConfig())
	p.Start()
	p.Exec()
	p.Stop()
}
