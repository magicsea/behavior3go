package actions

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
)

type Runner struct {
	Action
}

func (this *Runner) ctor() {
	this.SetName("Runner")
}

func (this *Runner) tick(tick *Tick) b3.Status {
	return b3.RUNNING
}
