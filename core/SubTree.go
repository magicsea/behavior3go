package core

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
)

//子树，通过Name关联树ID查找
type SubTree struct {
	Action
	//tree *BehaviorTree
}

func (this *SubTree) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}
/**
 *执行子树
 *使用sTree.Tick(tar, tick.Blackboard)的方法会导致每个树有自己的tick。
 *如果子树包含running状态，同时复用了子树会导致歧义。
 *改为只使用一个树，一个tick上下文。
**/
func (this *SubTree) OnTick(tick *Tick) b3.Status {

	//使用子树，必须先SetSubTreeLoadFunc
	//子树可能没有加载上来，所以要延迟加载执行
	sTree := subTreeLoadFunc(this.GetName())
	if nil == sTree {
		return b3.ERROR
	}

	if tick.GetTarget() == nil {
		panic("SubTree tick.GetTarget() nil !")
	}

	//tar := tick.GetTarget()
	//return sTree.Tick(tar, tick.Blackboard)

	tick.pushSubtreeNode(this)
	ret := sTree.GetRoot().Execute(tick)
	tick.popSubtreeNode()
	return ret
}

func (this *SubTree) String() string  {
	return "SBT_"+this.GetTitle()
}


var subTreeLoadFunc func(string) *BehaviorTree

//获取子树的方法
func SetSubTreeLoadFunc(f func(string) *BehaviorTree) {
	subTreeLoadFunc = f
}