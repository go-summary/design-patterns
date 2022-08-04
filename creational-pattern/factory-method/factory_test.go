package factory_method

import "testing"

func TestEventFactory(t *testing.T) {
	f := Factory{}
	e := f.Create(Start)
	if e.EventType() != Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = f.Create(End)
	if e.EventType() != End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}

func TestEventOF(t *testing.T) {
	e := OfStart()
	if e.EventType() != Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = OfEnd()
	if e.EventType() != End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}
