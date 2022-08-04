package test

import (
	"github.com/go-summary/design-patterns/creational-pattern/factory-method"
	"testing"
)

func TestEventFactory(t *testing.T) {
	f := factory_method.Factory{}
	e := f.Create(factory_method.Start)
	if e.EventType() != factory_method.Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = f.Create(factory_method.End)
	if e.EventType() != factory_method.End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}

func TestEventOF(t *testing.T) {
	e := factory_method.OfStart()
	if e.EventType() != factory_method.Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = factory_method.OfEnd()
	if e.EventType() != factory_method.End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}
