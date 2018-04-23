package core

import (
	_ "fmt"
)

/**
 * A new Tick object is instantiated every tick by BehaviorTree. It is passed
 * as parameter to the nodes through the tree during the traversal.
 *
 * The role of the Tick class is to store the instances of tree, debug,
 * target and blackboard. So, all nodes can access these informations.
 *
 * For internal uses, the Tick also is useful to store the open node after
 * the tick signal, in order to let `BehaviorTree` to keep track and close
 * them when necessary.
 *
 * This class also makes a bridge between nodes and the debug, passing the
 * node state to the debug if the last is provided.
 *
 * @module b3
 * @class Tick
**/
type Tick struct {
	/**
	 * The tree reference.
	 * @property {b3.BehaviorTree} tree
	 * @readOnly
	**/
	tree *BehaviorTree
	/**
	 * The debug reference.
	 * @property {Object} debug
	 * @readOnly
	 */
	debug interface{}
	/**
	 * The target object reference.
	 * @property {Object} target
	 * @readOnly
	**/
	target interface{}
	/**
	 * The blackboard reference.
	 * @property {b3.Blackboard} blackboard
	 * @readOnly
	**/
	Blackboard *Blackboard
	/**
	 * The list of open nodes. Update during the tree traversal.
	 * @property {Array} _openNodes
	 * @protected
	 * @readOnly
	**/
	_openNodes []IBaseNode

	/**
	 * The number of nodes entered during the tick. Update during the tree
	 * traversal.
	 *
	 * @property {Integer} _nodeCount
	 * @protected
	 * @readOnly
	**/
	_nodeCount int
}

func NewTick() *Tick {
	tick := &Tick{}
	tick.Initialize()
	return tick
}

/**
 * Initialization method.
 * @method Initialize
 * @construCtor
**/
func (this *Tick) Initialize() {
	// set by BehaviorTree
	this.tree = nil
	this.debug = nil
	this.target = nil
	this.Blackboard = nil

	// updated during the tick signal
	this._openNodes = make([]IBaseNode, 0)
	this._nodeCount = 0
}

func (this *Tick) GetTree() *BehaviorTree {
	return this.tree
}

/**
 * Called when entering a node (called by BaseNode).
 * @method _enterNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _enterNode(node IBaseNode) {
	this._nodeCount++
	this._openNodes = append(this._openNodes, node)

	// TODO: call debug here
}

/**
 * Callback when opening a node (called by BaseNode).
 * @method _openNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _openNode(node *BaseNode) {
	// TODO: call debug here
}

/**
 * Callback when ticking a node (called by BaseNode).
 * @method _tickNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _tickNode(node *BaseNode) {
	// TODO: call debug here
	//fmt.Println("Tick _tickNode :", this.debug, " id:", node.GetID(), node.GetTitle())
}

/**
 * Callback when closing a node (called by BaseNode).
 * @method _closeNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _closeNode(node *BaseNode) {
	// TODO: call debug here

	ulen := len(this._openNodes)
	if len(this._openNodes) > 0 {
		this._openNodes = this._openNodes[:ulen-1]
	}

}

/**
 * Callback when exiting a node (called by BaseNode).
 * @method _exitNode
 * @param {Object} node The node that called this method.
 * @protected
**/
func (this *Tick) _exitNode(node *BaseNode) {
	// TODO: call debug here
}

func (this *Tick) GetTarget() interface{} {
	return this.target
}
