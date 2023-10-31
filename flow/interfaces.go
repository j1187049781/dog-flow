package flow

type Node interface {
	// ID returns the ID of the node.
	ID() string
}


type Edge interface {
	// ID returns the ID of the edge.
	ID() string
	From() Node
	To() Node
	isPassable() bool
}

type Event interface {
	// ID returns the ID of the event.
	ID() string

	ToFlowInstanceId() string

	ToNodeId() string

	Handle(context *FlowContext) error
}

