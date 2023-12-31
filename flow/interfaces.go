package flow

type Handler interface {
	// OnEvent handles the event.
	OnEvent(context FlowContext, event Event) error
	// Enter is called when the node is entered.
	Enter(context FlowContext) error
	// Leave is called when the node is leaved.
	Leave(context FlowContext) error
}

// Node is a node in a DAG. It should be done by FlowContext's DoneNode method in the end of Run method.
type Node interface {
	// ID returns the ID of the node.
	ID() string
	// Init initializes the node when flow instance start.
	Init(context FlowContext) error
	// Run runs the node.
	Run(context FlowContext) error
	// IsDone returns whether the node is done.
	IsDone(context FlowContext) bool
	// IsRunning returns whether the node is running.
	IsRunning() bool
	// implement Handler
	Handler
}

type Edge interface {
	// ID returns the ID of the edge.
	ID() string

	From() Node

	To() Node

	IsPassable(context FlowContext) bool
}

type Event interface {
	// ID returns the ID of the event.
	ID() string

	ToFlowInstanceId() string

	ToNodeId() string

}

type FlowContext interface {
	GetFlowId() string

	DoneNode(nodeId string) error
}
