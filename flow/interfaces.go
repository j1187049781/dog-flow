package flow

// Node is a node in a DAG. It should be done by FlowContext's DoneNode method in the end of Run method.
type Node interface {
	// ID returns the ID of the node.
	ID() string
	// Init initializes the node.
	Init(context FlowContext) error
	// Run runs the node.
	Run(context FlowContext) error
	// IsDone returns whether the node is done.
	IsDone(context FlowContext) bool
	// IsStart returns whether the node is a start node.
	IsStart() bool
}

type Edge interface {
	// ID returns the ID of the edge.
	ID() string

	From() Node

	To() Node

	isPassable(context FlowContext) bool
}

type Event interface {
	// ID returns the ID of the event.
	ID() string

	ToFlowInstanceId() string

	ToNodeId() string

	Handle(context FlowContext) error
}

type FlowContext interface {
	GetFlowId() string

	DoneNode(nodeId string) error
}
