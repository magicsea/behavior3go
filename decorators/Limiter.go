package decorators

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

/**
 * This decorator limit the number of times its child can be called. After a
 * certain number of times, the Limiter decorator returns `FAILURE` without
 * executing the child.
 *
 * @module b3
 * @class Limiter
 * @extends Decorator
**/
type Limiter struct {
	Decorator
	maxLoop int
}

/**
 * Initialization method.
 *
 * Settings parameters:
 *
 * - **milliseconds** (*Integer*) Maximum time, in milliseconds, a child
 *                                can execute.
 *
 * @method Initialize
 * @param {Object} settings Object with parameters.
 * @construCtor
**/
func (this *Limiter) Initialize(setting *BTNodeCfg) {
	this.Decorator.Initialize(setting)
	this.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if this.maxLoop < 1 {
		panic("maxLoop parameter in MaxTime decorator is an obligatory parameter")
	}
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (this *Limiter) OnOpen(tick *Tick) {
	tick.Blackboard.Set("i", 0, tick.GetTree().GetID(), this.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Limiter) OnTick(tick *Tick) b3.Status {
	if this.GetChild() == nil {
		return b3.ERROR
	}
	var i = tick.Blackboard.GetInt("i", tick.GetTree().GetID(), this.GetID())
	if i < this.maxLoop {
		var status = this.GetChild().Execute(tick)
		if status == b3.SUCCESS || status == b3.FAILURE {
			tick.Blackboard.Set("i", i+1, tick.GetTree().GetID(), this.GetID())
		}
		return status
	}

	return b3.FAILURE
}
