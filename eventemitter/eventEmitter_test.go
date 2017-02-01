package eventemitter

import (
	"testing"

	"github.com/gopherjs/gopherjs/js"
)

var (
	e = New()
)

func TestEventNames(t *testing.T) {
	e.On("hi", func(args ...*js.Object) {
		println("on hi:", len(args))
	})
	ns := e.EventNames()
	if ns[0] != "hi" {
		t.Fail()
	}
}

func TestListener(t *testing.T) {
	e.On("hi2", func(args ...*js.Object) {
		println("on hi2:", len(args), args[0], args[1])
		if len(args) != 2 {
			t.Fail()
		}
		if args[0].Int() != 10 {
			t.Fail()
		}
		if args[1].String() != "str" {
			t.Fail()
		}
	})
	e.Emit("hi2", 10, "str")
}

func TestOnEvent(t *testing.T) {
	e.OnEvent("hi3", func(em *EventEmitter, args ...*js.Object) {
		t.Log("EventNames:", em.EventNames())
		if em != e {
			t.Fatal(" em != e ")
		}
		if len(em.EventNames()) != len(e.EventNames()) {
			t.Fatal(" enents mismatch ")
		}
	})
	e.Emit("hi3")
}
