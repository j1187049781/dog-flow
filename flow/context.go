package flow


// FlowContext is the context of a node.
type FlowContext struct {
	instance *FlowInstance
	node Node
}

// NewFlowContext creates a new FlowContext.
func NewFlowContext(instance *FlowInstance, node Node) *FlowContext {
	return &FlowContext{
		instance: instance,
		node: node,
	}
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
	return nil
}	

// DoneNode marks a node as done.
func (c *FlowContext) DoneNode(nodeId string) error {
	return nil
}