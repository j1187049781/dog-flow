package flow

type Node interface {
	// ID returns the ID of the node.
	ID() string
	// Run runs the node.
	Run(context *FlowContext) error
	// Done marks the node as done.
	Done(context *FlowContext) error
	// IsDone returns whether the node is done.
	IsDone(context *FlowContext) bool
	// IsStart returns whether the node is a start node.
	IsStart() bool
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

