package runtime


type Event interface {
	// ID returns the ID of the event.
	ID() string

	ToFlowInstanceId() string

	ToNodeId() string

	Handle(context *FlowContext) error
}