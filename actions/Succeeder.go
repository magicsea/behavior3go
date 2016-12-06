package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
)

type Succeeder struct {
	Action
}

func (this *Succeeder) ctor() {
	this.SetName("Succeeder")
}

func (this *Succeeder) tick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
