package decorators

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
)

/**
 * The Inverter decorator inverts the result of the child, returning `SUCCESS`
 * for `FAILURE` and `FAILURE` for `SUCCESS`.
 *
 * @module b3
 * @class Inverter
 * @extends Decorator
**/
type Inverter struct {
	Decorator
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Inverter) OnTick(tick *Tick) b3.Status {
	if this.GetChild() == nil {
		return b3.ERROR
	}

	var status = this.GetChild().Execute(tick)
	if status == b3.SUCCESS {
		status = b3.FAILURE
	} else if status == b3.FAILURE {
		status = b3.SUCCESS
	}

	return status
}
