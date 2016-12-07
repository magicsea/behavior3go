package actions

import (
	"time"

	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

/**
 * Wait a few seconds.
 *
 * @module b3
 * @class Wait
 * @extends Action
**/
type Wait struct {
	Action
	endTime int64
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
func (this *Wait) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.endTime = setting.GetPropertyAsInt64("milliseconds")
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (this *Wait) OnOpen(tick *Tick) {
	var startTime int64 = time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), this.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Wait) OnTick(tick *Tick) b3.Status {
	var currTime int64 = time.Now().UnixNano() / 1000000
	var startTime = tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), this.GetID())

	if currTime-startTime > this.endTime {
		return b3.SUCCESS
	}

	return b3.RUNNING
}
