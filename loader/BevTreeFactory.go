package loader

import (
	_ "fmt"
	_ "reflect"

	. "github.com/magicsea/behavior3go/actions"
	. "github.com/magicsea/behavior3go/composites"
	//. "github.com/magicsea/behavior3go/config"
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
	. "github.com/magicsea/behavior3go/decorators"
)

/*
//工厂扩展类型接口
type IBevFactory interface {
	CreateCond(cfg *BTNodeCfg) b3core.ICondition
	CreateSelector(cfg *BTNodeCfg) btnode.IBevSelector
	CreateTerNode(cfg *BTNodeCfg) btnode.IBevTerminal
}

//工厂扩展类型
var extBevFactory IBevFactory

func SetExtBevFactory(factroy IBevFactory) {
	extBevFactory = factroy
}
func CreateCond(cfg *BTNodeCfg) btcond.IPrecondition {
	switch cfg.Name {
	case "IfTrue":
		cond := btcond.NewPreconditionTRUE()
		return cond
	case "IfFalse":
		cond := btcond.NewPreconditionFALSE()
		return cond
	case "CondLess":
		cond := NewPreconditionLess(cfg.GetPropertyAsInt("index1"), cfg.GetPropertyAsInt("index2"))
		return cond
	case "CondMidValue":
		cond := NewCondMidValue(cfg.GetPropertyAsInt("index"), cfg.GetPropertyAsInt("min"), cfg.GetPropertyAsInt("max"))
		return cond
	case "CondEqualValue":
		cond := NewCondEqualValuee(cfg.GetPropertyAsInt("index"), cfg.GetPropertyAsInt("value"))
		return cond
	}
	if extBevFactory != nil {
		return extBevFactory.CreateCond(cfg)
	}
	return nil
}

func CreateSelector(cfg *BTNodeCfg) btnode.IBevSelector {
	switch cfg.Name {
	case "Priority":
		node := btnode.NewPrioritySelector(nil, nil)
		node.SetDebugName(cfg.Title)
		return node
	case "Sequence":
		node := btnode.NewSequenceSelector(nil, nil)
		node.SetDebugName(cfg.Title)
		return node
	case "Parallel":
		node := btnode.NewParallelSelector(nil, nil)
		node.SetDebugName(cfg.Title)
		return node
	case "Loop":
		//todo:测试不通过
		node := btnode.NewLoopSelector(nil, nil, cfg.GetPropertyAsInt("maxLoop"))
		node.SetDebugName(cfg.Title)
		return node
	case "Random":
		node := btnode.NewRandomSelector(nil, nil)
		node.SetDebugName(cfg.Title)
		return node
	}
	if extBevFactory != nil {
		return extBevFactory.CreateSelector(cfg)
	}
	return nil
}

func CreateTerNode(cfg *BTNodeCfg) btnode.IBevTerminal {
	switch cfg.Name {
	case "Log":
		node := &PrintNode{btnode.NewTerminalNode(nil, nil), cfg.Properties["info"].(string)}
		return node
	case "Wait":
		time := cfg.GetPropertyAsInt("milliseconds")
		node1 := NewWaitActNode(time, nil)
		return node1
	case "SetValue":
		node := &SetValueNode{btnode.NewTerminalNode(nil, nil), cfg.GetPropertyAsInt("index"), cfg.GetPropertyAsInt("value")}
		return node
	default:
		node1 := &UnknowNode{btnode.NewTerminalNode(nil, nil), "???" + cfg.Title}
		return node1
	}
	if extBevFactory != nil {
		return extBevFactory.CreateTerNode(cfg)
	}
	return nil
}

func buildBevTree(root btnode.IBevNode, nodes map[string]btnode.IBevNode, cfg *BTNodeCfg, config *BTTreeCfg, conds map[string]btcond.IPrecondition) {
	for _, c := range cfg.Children {
		if node, ok := nodes[c]; ok {
			root.AddChildNode(node)
			newCfg := config.Nodes[c]
			buildBevTree(node, nodes, &newCfg, config, conds)
		} else {
			//是否cond
			if cond, ok := conds[c]; ok {
				condCfg := config.Nodes[c]
				node := nodes[condCfg.Child]
				newCfg := config.Nodes[condCfg.Child]
				if selector, ok := node.(btnode.IBevSelector); ok {
					node = btnode.NewSelector(selector)
				}
				node.SetNodePrecondition(cond)
				root.AddChildNode(node)
				buildBevTree(node, nodes, &newCfg, config, conds)

			} else {
				fmt.Println("Err: no node2=>", c)
			}
		}
	}
}

//创建行为树
func NewBevTree(config *BTTreeCfg) *GameBevTree {

	nodes2 := make(map[string]btnode.IBevNode)
	conds := make(map[string]btcond.IPrecondition)

	for _, v := range config.Nodes {
		var node2 btnode.IBevNode
		key := v.Id
		if len(v.Child) < 1 {
			node2 = CreateSelector(&v)
			if node2 == nil {
				terNode := CreateTerNode(&v)
				if terNode == nil {
					continue
				}
				node2 = btnode.NewTerminal(terNode)
				node2.SetDebugName(v.Title)
			}

			if node2 == nil {
				node2 = btnode.NewTerminalNode(nil, nil)
				node2.SetDebugName("???" + v.Title)
			}
			nodes2[key] = node2
		} else {
			//cond
			cond := CreateCond(&v)
			if cond == nil {
				fmt.Println("err no cond:", v.Name)
				continue
			}
			conds[key] = cond

		}

	}

	root := nodes2[config.Root]
	rootCfg := config.Nodes[config.Root]
	buildBevTree(root, nodes2, &rootCfg, config, conds)
	//TestRenderTree(root, 4, inboard, outboard, 1)

	inboard := btboard.NewBlackboard()
	outboard := btboard.NewBlackboard()
	tree := &GameBevTree{root, inboard, outboard}

	return tree
}
*/

func CreateStructMaps() *b3.RegisterStructMaps {
	st := b3.NewRegisterStructMaps()
	//actions
	st.Register("Error", &Error{})
	st.Register("Failer", &Failer{})
	st.Register("Runner", &Runner{})
	st.Register("Succeeder", &Succeeder{})
	st.Register("Wait", &Wait{})
	//composites
	st.Register("MemPriority", &MemPriority{})
	st.Register("MemSequence", &MemSequence{})
	st.Register("Priority", &Priority{})
	st.Register("Sequence", &Sequence{})

	//decorators
	st.Register("Inverter", &Inverter{})
	st.Register("Limiter", &Limiter{})
	st.Register("MaxTime", &MaxTime{})
	st.Register("Repeater", &Repeater{})
	st.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})
	return st
	/*
		var regStruct map[string]interface{}
		regStruct = make(map[string]interface{})
		//actions
		regStruct["Error"] = Error{}
		regStruct["Failer"] = Failer{}
		regStruct["Runner"] = Runner{}
		regStruct["Succeeder"] = Succeeder{}
		regStruct["Wait"] = Wait{}

		//composites
		regStruct["MemPriority"] = MemPriority{}
		regStruct["MemSequence"] = MemSequence{}
		regStruct["Priority"] = Priority{}
		regStruct["Sequence"] = Sequence{}

		//decorators
		regStruct["Inverter"] = Inverter{}
		regStruct["Limiter"] = Limiter{}
		regStruct["MaxTime"] = MaxTime{}
		regStruct["Repeater"] = Repeater{}
		regStruct["RepeatUntilFailure"] = RepeatUntilFailure{}
		regStruct["RepeatUntilSuccess"] = RepeatUntilSuccess{}

		if val, ok := regStruct[name]; ok {
			t := reflect.ValueOf(val).Type()
			v := reflect.New(t).Interface()
			fmt.Println(v)
			data, succ := v.(IBaseNode)
			return data, succ
		}

		return nil, false
	*/
}
