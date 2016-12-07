package core

import (
	_ "fmt"

	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/config"
)

type IBaseWrapper interface {
	_execute(tick *Tick) b3.Status
	_enter(tick *Tick)
	_open(tick *Tick)
	_tick(tick *Tick) b3.Status
	_close(tick *Tick)
	_exit(tick *Tick)
}
type IBaseNode interface {
	IBaseWrapper

	Ctor()
	Initialize(params *BTNodeCfg)
	GetCategory() string
	Execute(tick *Tick) b3.Status
	GetName() string
	GetTitle() string
	SetBaseNodeWorker(worker IBaseWorker)
}

/**
 * The BaseNode class is used as super class to all nodes in BehaviorJS. It
 * comprises all common variables and methods that a node must have to
 * execute.
 *
 * **IMPORTANT:** Do not inherit from this class, use `b3.Composite`,
 * `b3.Decorator`, `b3.Action` or `b3.Condition`, instead.
 *
 * The attributes are specially designed to serialization of the node in a
 * JSON format. In special, the `parameters` attribute can be set into the
 * visual editor (thus, in the JSON file), and it will be used as parameter
 * on the node initialization at `BehaviorTree.load`.
 *
 * BaseNode also provide 5 callback methods, which the node implementations
 * can override. They are `enter`, `open`, `tick`, `close` and `exit`. See
 * their documentation to know more. These callbacks are called inside the
 * `_execute` method, which is called in the tree traversal.
 *
 * @module b3
 * @class BaseNode
**/
type BaseNode struct {
	IBaseWorker
	/**
	 * Node ID.
	 * @property {string} id
	 * @readonly
	**/
	id string

	/**
	 * Node name. Must be a unique identifier, preferable the same name of the
	 * class. You have to set the node name in the prototype.
	 *
	 * @property {String} name
	 * @readonly
	**/
	name string

	/**
	 * Node category. Must be `b3.COMPOSITE`, `b3.DECORATOR`, `b3.ACTION` or
	 * `b3.CONDITION`. This is defined automatically be inheriting the
	 * correspondent class.
	 *
	 * @property {CONSTANT} category
	 * @readonly
	**/
	category string

	/**
	 * Node title.
	 * @property {String} title
	 * @optional
	 * @readonly
	**/
	title string

	/**
	 * Node description.
	 * @property {String} description
	 * @optional
	 * @readonly
	**/
	description string

	/**
	 * A dictionary (key, value) describing the node parameters. Useful for
	 * defining parameter values in the visual editor. Note: this is only
	 * useful for nodes when loading trees from JSON files.
	 *
	 * **Deprecated since 0.2.0. This is too similar to the properties
	 * attribute, thus, this attribute is deprecated in favor to
	 * `properties`.**
	 *
	 * @property {Object} parameters
	 * @deprecated since 0.2.0.
	 * @readonly
	**/
	parameters map[string]interface{}

	/**
	 * A dictionary (key, value) describing the node properties. Useful for
	 * defining custom variables inside the visual editor.
	 *
	 * @property properties
	 * @type {Object}
	 * @readonly
	**/
	properties map[string]interface{}
}

func (this *BaseNode) Ctor() {

}

func (this *BaseNode) SetName(name string) {
	this.name = name
}
func (this *BaseNode) SetTitle(name string) {
	this.name = name
}

func (this *BaseNode) SetBaseNodeWorker(worker IBaseWorker) {
	this.IBaseWorker = worker
}

/**
 * Initialization method.
 * @method Initialize
 * @construCtor
**/
func (this *BaseNode) Initialize(params *BTNodeCfg) {
	//this.id = b3.CreateUUID()
	//this.title       = this.title || this.name
	this.description = ""
	this.parameters = make(map[string]interface{})
	this.properties = make(map[string]interface{})

	this.id = params.Id //|| node.id;
	this.name = params.Name
	this.title = params.Title             //|| node.title;
	this.description = params.Description // || node.description;
	this.properties = params.Properties   //|| node.properties;

}

func (this *BaseNode) GetCategory() string {
	return this.category
}

func (this *BaseNode) GetID() string {
	return this.id
}

func (this *BaseNode) GetName() string {
	return this.name
}
func (this *BaseNode) GetTitle() string {
	//fmt.Println("GetTitle ", this.title)
	return this.title
}

/**
 * This is the main method to propagate the tick signal to this node. This
 * method calls all callbacks: `enter`, `open`, `tick`, `close`, and
 * `exit`. It only opens a node if it is not already open. In the same
 * way, this method only close a node if the node  returned a status
 * different of `b3.RUNNING`.
 *
 * @method _execute
 * @param {Tick} tick A tick instance.
 * @return {Constant} The tick state.
 * @protected
**/
func (this *BaseNode) _execute(tick *Tick) b3.Status {
	//fmt.Println("_execute :", this.title)
	// ENTER
	this._enter(tick)

	// OPEN
	if !tick.Blackboard.GetBool("isOpen", tick.tree.id, this.id) {
		this._open(tick)
	}

	// TICK
	var status = this._tick(tick)

	// CLOSE
	if status != b3.RUNNING {
		this._close(tick)
	}

	// EXIT
	this._exit(tick)

	return status
}
func (this *BaseNode) Execute(tick *Tick) b3.Status {
	return this._execute(tick)
}

/**
 * Wrapper for enter method.
 * @method _enter
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (this *BaseNode) _enter(tick *Tick) {
	tick._enterNode(this)
	this.OnEnter(tick)
}

/**
 * Wrapper for open method.
 * @method _open
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (this *BaseNode) _open(tick *Tick) {
	//fmt.Println("_open :", this.title)
	tick._openNode(this)
	tick.Blackboard.Set("isOpen", true, tick.tree.id, this.id)
	this.OnOpen(tick)
}

/**
 * Wrapper for tick method.
 * @method _tick
 * @param {Tick} tick A tick instance.
 * @return {Constant} A state constant.
 * @protected
**/
func (this *BaseNode) _tick(tick *Tick) b3.Status {
	//fmt.Println("_tick :", this.title)
	tick._tickNode(this)
	return this.OnTick(tick)
}

/**
 * Wrapper for close method.
 * @method _close
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (this *BaseNode) _close(tick *Tick) {
	tick._closeNode(this)
	tick.Blackboard.Set("isOpen", false, tick.tree.id, this.id)
	this.OnClose(tick)
}

/**
 * Wrapper for exit method.
 * @method _exit
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (this *BaseNode) _exit(tick *Tick) {
	tick._exitNode(this)
	this.OnExit(tick)
}
