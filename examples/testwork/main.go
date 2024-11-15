/*
从原生工程文件加载
*/
package main

import (
	"fmt"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
	. "github.com/magicsea/behavior3go/examples/share"
	. "github.com/magicsea/behavior3go/loader"
	"sync"
	"time"
)

// 所有的树管理
var mapTreesByID = sync.Map{}
var maps = b3.NewRegisterStructMaps()

func init() {
	//自定义节点注册
	maps.Register("Log", new(LogTest))
	maps.Register("SetValue", new(SetValue))
	maps.Register("IsValue", new(IsValue))

	//获取子树的方法
	SetSubTreeLoadFunc(func(id string) *BehaviorTree {
		//println("==>load subtree:",id)
		t, ok := mapTreesByID.Load(id)
		if ok {
			return t.(*BehaviorTree)
		}
		return nil
	})
}

func main() {
	projectConfig, ok := LoadRawProjectCfg("testwork.b3")
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}

	var firstTree *BehaviorTree
	//载入
	for _, v := range projectConfig.Data.Trees {
		tree := CreateBevTreeFromConfig(&v, maps)
		tree.Print()
		mapTreesByID.Store(v.ID, tree)
		if firstTree == nil {
			firstTree = tree
		}
	}
	time.Sleep(time.Second)
	//输入板
	board := NewBlackboard()
	//循环每一帧
	for i := 0; i < 40; i++ {
		firstTree.Tick(i, board)
		time.Sleep(time.Millisecond * 100)
	}
}
