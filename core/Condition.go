package core

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
)

type ICondition interface {
	IBaseNode
}

type Condition struct {
	BaseNode
}

func (this *Condition) ctor() {
	this.category = b3.CONDITION
}

/**
 * Initialization method.
 *
 * @method initialize
 * @constructor
**/
func (this *Condition) initialize(params *BTNodeCfg) {
	this.BaseNode.initialize(params)
}
