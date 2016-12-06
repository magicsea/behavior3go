package core

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
)

type IDecorator interface {
	IBaseNode
	SetChild(child IBaseNode)
	GetChild() IBaseNode
}

type Decorator struct {
	BaseNode
	child IBaseNode
}

func (this *Decorator) ctor() {
	this.category = b3.DECORATOR
}

/**
 * Initialization method.
 *
 * @method initialize
 * @constructor
**/
func (this *Decorator) initialize(params *BTNodeCfg) {
	this.BaseNode.initialize(params)

}

//GetChild
func (this *Decorator) GetChild() IBaseNode {
	return this.child
}

func (this *Decorator) SetChild(child IBaseNode) {
	this.child = child
}
