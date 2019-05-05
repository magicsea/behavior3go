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
)

//所有的树管理
var mapTreesByID = sync.Map{}

func init() {
	//获取子树的方法
	SetSubTreeLoadFunc(func(id string) *BehaviorTree {
		t, ok := mapTreesByID.Load(id)
		if ok {
			return t.(*BehaviorTree)
		}
		return nil
	})
}

func main() {
	projectConfig, ok := LoadRawProjectCfg("example.b3")
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}

	//自定义节点注册
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	var firstTree *BehaviorTree
	//载入
	for _, v := range projectConfig.Data.Trees {
		tree := CreateBevTreeFromConfig(&v, maps)
		tree.Print()
		//保存到树管理
		mapTreesByID.Store(v.ID, tree)
		if firstTree == nil {
			firstTree = tree
		}
	}

	//输入板
	board := NewBlackboard()
	//循环每一帧
	for i := 0; i < 5; i++ {
		firstTree.Tick(i, board)
	}
}
