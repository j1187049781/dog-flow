package event

type Event interface {
	// ID returns the ID of the event.
	ID() string

	ToFlowInstanceId() string

	ToNodeId() string
}