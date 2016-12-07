package decorators

import (
	"time"

	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

/**
 * The MaxTime decorator limits the maximum time the node child can execute.
 * Notice that it does not interrupt the execution itself (i.e., the child
 * must be non-preemptive), it only interrupts the node after a `RUNNING`
 * status.
 *
 * @module b3
 * @class MaxTime
 * @extends Decorator
**/
type MaxTime struct {
	Decorator
	maxTime int64
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
func (this *MaxTime) Initialize(setting *BTNodeCfg) {
	this.Decorator.Initialize(setting)
	this.maxTime = setting.GetPropertyAsInt64("maxTime")
	if this.maxTime < 1 {
		panic("maxTime parameter in Limiter decorator is an obligatory parameter")
	}
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (this *MaxTime) OnOpen(tick *Tick) {
	var startTime int64 = time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), this.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *MaxTime) OnTick(tick *Tick) b3.Status {
	if this.GetChild() == nil {
		return b3.ERROR
	}
	var currTime int64 = time.Now().UnixNano() / 1000000
	var startTime int64 = tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), this.GetID())
	var status = this.GetChild().Execute(tick)
	if currTime-startTime > this.maxTime {
		return b3.FAILURE
	}

	return status
}
