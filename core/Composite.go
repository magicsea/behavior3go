package core

import (
	"fmt"

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
	BaseWorker

	children []IBaseNode
}

func (this *Composite) Ctor() {

	this.category = b3.COMPOSITE
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (this *Composite) Initialize(params *BTNodeCfg) {
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
	this.children = make([]IBaseNode, 0)
	fmt.Println("Composite Initialize")
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
func (this *Composite) tick(tick *Tick) b3.Status {
	fmt.Println("tick Composite1")
	return b3.ERROR
}
