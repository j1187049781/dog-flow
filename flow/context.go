package flow


// FlowContext is the context of a node.
type FlowContext struct {
	runtime *Runtime
	instance *FlowInstance
	node Node
}

// NewFlowContext creates a new FlowContext.
func NewFlowContext(runtime *Runtime, instance *FlowInstance, node Node) *FlowContext {
	return &FlowContext{
		runtime: runtime,
		instance: instance,
		node: node,
	}
}

// Runtime returns the runtime of the context.
func (c *FlowContext) Runtime() *Runtime {
	return c.runtime
}

// Instance returns the instance of the context.
func (c *FlowContext) Instance() *FlowInstance {
	return c.instance
}

// Node returns the node of the context.
func (c *FlowContext) Node() Node {
	return c.node
}

// RunNode runs a node.
func (c *FlowContext) RunNode(nodeId string) error {
	return c.instance.runANode(nodeId)
}	

// DoneNode marks a node as done.
func (c *FlowContext) DoneNode(nodeId string) error {
	return c.instance.doneANode(nodeId)
}