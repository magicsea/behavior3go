package composites

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
)

type Priority struct {
	Composite
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Priority) OnTick(tick *Tick) b3.Status {
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.FAILURE {
			return status
		}
	}
	return b3.FAILURE
}
