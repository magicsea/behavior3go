package core

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
)

type IComposite interface {
	IBaseNode
	GetChildCount() int
	GetChild(index int) IBaseNode
	AddChild(child IBaseNode)
}

type Composite struct {
	BaseNode
	children []IBaseNode
}

func (this *Composite) ctor() {
	this.category = b3.COMPOSITE
}

/**
 * Initialization method.
 *
 * @method initialize
 * @constructor
**/
func (this *Composite) initialize(params *BTNodeCfg) {
	this.BaseNode.initialize(params)
	this.children = make([]IBaseNode, 0)
}

/**
 *
 * @method GetChildCount
 * @getChildCount
**/
func (this *Composite) GetChildCount() int {
	return len(this.children)
}

//GetChild
func (this *Composite) GetChild(index int) IBaseNode {
	return this.children[index]
}

//AddChild
func (this *Composite) AddChild(child IBaseNode) {
	this.children = append(this.children, child)
}
