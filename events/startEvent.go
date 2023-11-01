package events

// StartEvent is an event that is fired when a node is started.
type StartEvent struct {
	Id string
	FlowInstanceId string
	NodeId string
}

func NewStartEvent(id string, flowInstanceId string, nodeId string) *StartEvent {
	return &StartEvent{
		Id: id,
		FlowInstanceId: flowInstanceId,
		NodeId: nodeId,
	}
}

func (e *StartEvent) ID() string {
	return e.Id
}

func (e *StartEvent) ToFlowInstanceId() string {
	return e.FlowInstanceId
}

func (e *StartEvent) ToNodeId() string {
	return e.NodeId
}


