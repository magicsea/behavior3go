package core

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
)

//子树，通过Name关联树ID查找
type SubTree struct {
	Action
}

func (this *SubTree) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
}

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
	tar := tick.GetTarget()
	//	glog.Info("subtree: ", this.treeName, " id ", player.id)
	return sTree.Tick(tar, tick.Blackboard)
}

var subTreeLoadFunc func(string) *BehaviorTree

//获取子树的方法
func SetSubTreeLoadFunc(f func(string) *BehaviorTree) {
	subTreeLoadFunc = f
}
