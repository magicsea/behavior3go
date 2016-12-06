package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
)

type Error struct {
	Action
	st string
}

func (this *Error) ctor() {
	this.SetName("Error")
	this.st = "AAAAAAAAAA"
}

func (this *Error) tick(tick *Tick) b3.Status {
	return b3.ERROR
}
