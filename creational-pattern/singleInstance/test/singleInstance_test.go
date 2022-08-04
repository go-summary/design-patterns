package test

import (
	"github.com/go-summary/design-patterns/creational-pattern/singleInstance/hungryman"
	"github.com/go-summary/design-patterns/creational-pattern/singleInstance/lazyman"
	"testing"
)

func TestInstance0(t *testing.T) {
	msg0 := hungryman.Instance().GetMsg()
	if msg0.Count != 0 {
		t.Errorf("expect msg count %d, but actual %d.", 0, msg0.Count)
	}
	msg0.Count = 1
	hungryman.Instance().AddMsg(msg0)
	msg1 := hungryman.Instance().GetMsg()
	if msg1.Count != 1 {
		t.Errorf("expect msg count %d, but actual %d.", 1, msg1.Count)
	}
}

func TestInstance1(t *testing.T) {
	msg0 := lazyman.Instance().GetMsg()
	if msg0.Count != 0 {
		t.Errorf("expect msg count %d, but actual %d.", 0, msg0.Count)
	}
	msg0.Count = 1
	lazyman.Instance().AddMsg(msg0)
	msg1 := lazyman.Instance().GetMsg()
	if msg1.Count != 1 {
		t.Errorf("expect msg count %d, but actual %d.", 1, msg1.Count)
	}
}
